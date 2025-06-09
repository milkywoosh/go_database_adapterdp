package internal

import (
	"context"
	"database/sql"
	"fmt"
)

type PurchaseStoreTx interface {
	PurchaseBookTx(ctx context.Context, arg CreatePurchaseBookTxParams) (CreatePurchaseBookTxResult, error)
	EditListBookTx(ctx context.Context, arg EditBookToPurchaseParams) (int64, error)
}

type PurchaseStore struct {
	connPool         *sql.DB
	*PurchaseQueries // the sqlc-generated querier
}

func NewPurchaseStore(db *sql.DB, dbtype string) *PurchaseStore {
	return &PurchaseStore{
		connPool:        db,
		PurchaseQueries: NewPurchaseQueries(db, dbtype),
	}
}

func (store *PurchaseStore) PurchaseBookTx(ctx context.Context, arg CreatePurchaseBookTxParams) (CreatePurchaseBookTxResult, error) {
	// var result_purchase_history PurchaseHistory
	// var err error
	// err = store.execTx(ctx, func(q *Queries) error {
	// 	result_purchase_history, err = q.CreatePurchaseHistory(ctx, CreatePurchaseHistoryParams{})

	// 	return err
	// })

	return CreatePurchaseBookTxResult{}, nil
}

// jika ingin mengubah jumlah list book
func (store *PurchaseStore) EditListBookTx(ctx context.Context, arg EditBookToPurchaseParams) (int64, error) {
	var err error
	var result sql.Result
	var rowsAffected int64
	// note misalnya tidak dalam transaksi, apakah akan terjadi update sebagian??? setelah tested result: iya
	// note jika akan melakukan beberapa operasi UPDATE, INSERT disertai logic harus dalam *SQLStore execTx() function!!!

	err = store.execTx(ctx, func(q *PurchaseQueries) error {

		err := q.LockRowEditListBookPG(ctx, arg.BookID, arg.PurchaseHistoryID, arg.PurchaseNumber)
		if err != nil {
			return err
		}

		result, err = q.EditListBookPG(ctx, arg.BookID, arg.PurchaseHistoryID, arg.PurchaseNumber, arg.Qty, arg.TotalPrice)
		if err != nil {
			return err
		}

		rowsAffected, err = result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected < 1 {
			var err ErrUpdateNolData
			err.Msg = fmt.Sprintln("error, tidak ada data terupdate")
			return err
		}
		if rowsAffected > 1 {

			var err ErrUpdateMultipleData
			err.Msg = fmt.Sprintf("error, terupdate ==> %d data", rowsAffected)
			return err
		}

		return nil
	})

	return rowsAffected, err

}

func (store *PurchaseStore) DeletePurchaseTx(ctx context.Context, args DeletePurchaseItemsTxParams) error {
	// param : purchase_number

	// if status != "pending" {
	// purchase_number berikut ==> tidak dapat dihapus. Status sudah 'completed'
	// }

	// << transaction >>
	// lock rows by purchase_number
	// delete operation by purchase_number

	var err error
	var PurchaseHistories PurchaseHistory

	if store.dbtype == "POSTGRES" {
		err = store.execTx(ctx, func(q *PurchaseQueries) error {

			err := q.LockRowPrcItemByPrcNumber(ctx, args.PurchaseNumber)
			if err != nil {
				return fmt.Errorf("err lock row deleting: %w", err)
			}

			PurchaseHistories.Status, err = q.CheckStatusPurchase(ctx, args.PurchaseNumber)
			if err != nil {
				return err
			}
			if PurchaseHistories.Status != "pending" {
				var err ErrStatusNotAcceptable
				err.Msg = fmt.Sprintf("status saat ini ==> '%s' sehingga tidak dapat proses penghapusan", PurchaseHistories.Status)
				return err
			}

			err = q.DeletePrcItemWithHistory(ctx, args.PurchaseNumber)
			if err != nil {
				return err
			}

			return nil
		})

		return err

	} else if store.dbtype == "ORACLE" {
		var err ErrDBTypeNotImplemented
		err.Msg = fmt.Sprintf("DB Type is not currently implemented ==> %s", store.dbtype)
		return err
	} else {
		var err ErrDBTypeNotImplemented
		err.Msg = fmt.Sprintf("DB Type is not currently implemented ==> %s", store.dbtype)
		return err
	}

	// return fmt.Errorf("not implemented yet %s", "not ready")

}

func (store *PurchaseStore) FinalizePurchaseTx(ctx context.Context, purchaseNumber string) error {

	err := store.execTx(ctx, func(q *PurchaseQueries) error {

		err := q.AdjustBookQuantity(ctx, purchaseNumber)
		if err != nil {
			return fmt.Errorf("err adjust book qty %w", err)
		}
		return nil
	})

	return fmt.Errorf("not implemented yet %v", err)
}

func (store *PurchaseStore) CreatePurchaseHistoryTx(ctx context.Context, arg CreatePurchaseHistoryParams) (PurchaseHistory, error) {
	var i PurchaseHistory

	err := store.execTx(ctx, func(pq *PurchaseQueries) error {
		var err error

		i, err = pq.CreatePurchaseHistory(ctx, arg)
		if err != nil {
			return err
		}

		return nil
	})

	return i, err
}
