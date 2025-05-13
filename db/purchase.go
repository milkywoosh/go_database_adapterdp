package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/luke_design_pattern/util"
)

// {
// 	"PURCHASE_HISTORIES": [
// 		{
// 			"ID" : 3,
// 			"TRANSACTION_NUMBER" : "TRXPRC_20250421_RANDOMCHAR1"
// 			"DATE_OF_SALE" : "2024-12-17T16:20:18.000Z",
// 			"CUSTOMER_ID" : 4,
// 			"TOTAL_PRICE_PAYMENT" : 110000,
// 			"STATUS" : "completed",
// 		}
// 	]}

// {
// 	"PURCHASE_ITEMS": [
// 		{
// 			"ID" : 6,
// 			"BOOK_ID" : 93,
// 			"PURCHASE_HISTORY_ID" : 2,
// 			"QTY" : 2,
// 			"TOTAL_PRICE" : 0
// 		}
// 	]}

// create temporary purchase history

type CreatePurchaseBookParams struct {
	// Date               *time.Time ==> auto generate from golang time.Date
	BookID            int
	PurchaseHistoryID int
	Qty               int
	TotalPrice        int
}

func GenerateRandomTrxNumber(CustomerID int) string {
	year, month, date := time.Now().Date()

	month_int := int(month)
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	scd := time.Now().Second() * int(util.RandomInt(0, 5000))
	custid := CustomerID
	generateTrxNumber := fmt.Sprintf("PRCBOOK%d%d%d%d%d%d%d", year, month_int, date, hour, minute, scd, custid)

	return generateTrxNumber
}

type CreatePurchaseHistoryResult struct {
}

type BookToPurchase struct {
	BookID            int
	Qty               int
	PurchaseHistoryID int
	CurrentStockQty   int
}

type ListBooksToPurchase []BookToPurchase

type CreateBookToPurchaseParams struct {
	BookID            int
	PurchaseHistoryID int
	Qty               int
	TotalPrice        float64
	PurchaseNumber    string
}

type EditBookToPurchaseParams struct {
	Qty               int
	TotalPrice        float64
	BookID            int
	PurchaseHistoryID int
	PurchaseNumber    string
}

// CreatePurchase() PurchaseHistory => purchase history
// AddListBook()
// DeletePurchase() ["delete list to buy", "delete purchase history"]
// FinalizePurchase() => update purchase history "completed", set "total_price_payment"

// AddStock(bookID, )
// DecreaseStock()

const createPurchaseHistoryOra = `
INSERT INTO PURCHASE_HISTORIES (
	DATE_OF_SALE, 
	CUSTOMER_ID, 
	TOTAL_PRICE_PAYMENT, 
	"STATUS", 
	PURCHASE_NUMBER
) VALUES (CURRENT_TIMESTAMP, :1, :2, :3, :4)
	RETURNING ID, PURCHASE_NUMBER INTO
`
const createPurchaseHistoryPG = `
INSERT INTO PURCHASE_HISTORIES (
	DATE_OF_SALE, 
	CUSTOMER_ID, 
	TOTAL_PRICE_PAYMENT, 
	STATUS, 
	PURCHASE_NUMBER
) VALUES (CURRENT_TIMESTAMP, $1, $2, $3, $4)
	RETURNING ID, DATE_OF_SALE, CUSTOMER_ID, TOTAL_PRICE_PAYMENT,STATUS, PURCHASE_NUMBER
`

type CreatePurchaseHistoryParams struct {
	// Date              *time.Time ==> auto generate from golang time.Date
	CustomerID        int
	TotalPricePayment float64
	Status            string // pending or completed
	PurchaseNumber    string // PRCBOOK_20250421_RANDOMCHAR
}

func (q *Queries) CreatePurchaseHistory(ctx context.Context, arg CreatePurchaseHistoryParams) (PurchaseHistory, error) {

	if q.dbtype == "ORACLE" {
		var i PurchaseHistory

		_, err := q.db.ExecContext(ctx, createPurchaseHistoryOra,
			arg.CustomerID,
			arg.TotalPricePayment,
			arg.Status,
			arg.PurchaseNumber,
			sql.Out{Dest: &i.PurchaseID},
			sql.Out{Dest: &i.PurchaseNumber},
		)

		return i, err
	} else if q.dbtype == "POSTGRES" {
		var i PurchaseHistory
		// ID, DATE_OF_SALE, CUSTOMER_ID, TOTAL_PRICE_PAYMENT,"STATUS", PURCHASE_NUMBER
		err := q.db.QueryRowContext(ctx, createPurchaseHistoryPG,
			arg.CustomerID,
			arg.TotalPricePayment,
			arg.Status,
			arg.PurchaseNumber,
		).Scan(
			&i.PurchaseID,
			&i.Date,
			&i.CustomerID,
			&i.TotalPricePayment,
			&i.Status,
			&i.PurchaseNumber,
		)
		return i, err
	} else {
		return PurchaseHistory{}, fmt.Errorf("tipe database berikut tidak diketahui ==> %s", q.dbtype)
	}

}

const fetchBookPG = `
	SELECT 
		b.ID, 
		b.STOCK_QTY, 
		b.TITLE,
		b.PRICE 
	FROM BOOKS b WHERE b.ID = $1
`
const checkExistedBookListPG string = `
	SELECT pi.ID FROM PURCHASE_ITEMS pi
	WHERE pi.PURCHASE_NUMBER = $1
	AND pi.PURCHASE_HISTORY_ID = $2
	AND pi.BOOK_ID = $3
`

const createNewPurchaseItemsBookOra = `
 INSERT INTO PURCHASE_ITEMS 
	(BOOK_ID, PURCHASE_HISTORY_ID, QTY, TOTAL_PRICE, PURCHASE_NUMBER) 
 VALUES(:1, :2, :3, :4, :5)
`
const createNewPurchaseItemsBookPG = `
 INSERT INTO PURCHASE_ITEMS 
	(BOOK_ID, PURCHASE_HISTORY_ID, QTY, TOTAL_PRICE, PURCHASE_NUMBER) 
 VALUES($1, $2, $3, $4, $5)
`

// add book one by one
func (q *Queries) AddListBook(ctx context.Context, arg CreateBookToPurchaseParams) (BookToPurchase, error) {

	// note ==> harusnya dalam transaksi
	var i BookToPurchase
	var bookModel Book
	var existedOnBookList bool = false
	var purchaseHistID int

	// check jika list buku dengan history id dan purchase number yg sama sudah ada error
	err := q.db.QueryRowContext(ctx, checkExistedBookListPG, arg.PurchaseNumber, arg.PurchaseHistoryID, arg.BookID).Scan(
		&purchaseHistID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			existedOnBookList = false
		} else {
			return i, err
		}
	} else {
		existedOnBookList = true
	}

	if existedOnBookList {
		var err ErrStokBukuHabis
		err.Msg = fmt.Sprintf("list buku %d pada purchase number %s sudah ada, lakukan edit list untuk ubah jumlah buku", arg.BookID, arg.PurchaseNumber)

		return i, err
	}

	err = q.db.QueryRowContext(ctx, fetchBookPG, arg.BookID).Scan(
		&bookModel.BookID,
		&bookModel.StockQty,
		&bookModel.Title,
		&bookModel.Price,
	)

	i.BookID = bookModel.BookID
	i.Qty = arg.Qty
	i.PurchaseHistoryID = arg.PurchaseHistoryID

	if err != nil {
		if err == sql.ErrNoRows {

			var err ErrIDBukuTidakTerdaftar
			err.Msg = fmt.Sprintf("ID buku tersebut tidak terdaftar ==> %d", arg.BookID)
			return i, err
		}

		return i, err
	}

	i.CurrentStockQty = bookModel.StockQty

	if bookModel.StockQty < 1 {

		var err ErrStokBukuHabis
		err.Msg = fmt.Sprintln("stock buku habis")
		return i, err
	}

	// check kuota < permintaan
	if bookModel.StockQty < arg.Qty {

		var err ErrStokBukuKurang
		err.Msg = fmt.Sprintln("gagal, jumlah pembelian buku melebihi stok persediaan")
		return i, err
	}

	totalPricePerBook := bookModel.Price * float64(arg.Qty)

	_, err = q.db.ExecContext(ctx, createNewPurchaseItemsBookPG,
		arg.BookID,
		arg.PurchaseHistoryID,
		arg.Qty,
		totalPricePerBook,
		arg.PurchaseNumber,
	)
	if err != nil {
		if strings.Contains(err.Error(), "ORA-02291") {
			log.Printf("2: %v", err)
			return i, err
		}
		return i, err
	}

	return i, nil
}

const editListBookPG string = `
	update purchase_items 
		set qty = $4,
		total_price = $5
	where book_id = $1
		and purchase_history_id = $2
		and purchase_number = $3
`
const lockRowEditListBookPG string = `
	SELECT 1
	FROM purchase_items
	WHERE book_id = $1
		AND purchase_history_id = $2
		AND purchase_number = $3
	FOR UPDATE
`

// jika ingin mengubah jumlah list book
func (store *SQLStore) EditListBookTx(ctx context.Context, arg EditBookToPurchaseParams) (int64, error) {
	var err error
	var result sql.Result
	var rowsAffected int64
	// note misalnya tidak dalam transaksi, apakah akan terjadi update sebagian??? setelah tested result: iya
	// note jika akan melakukan beberapa operasi UPDATE, INSERT disertai logic harus dalam *SQLStore execTx() function!!!

	err = store.execTx(ctx, func(q *Queries) error {
		_, err = q.db.ExecContext(ctx, lockRowEditListBookPG, arg.BookID, arg.PurchaseHistoryID, arg.PurchaseNumber)
		if err != nil {
			return err
		}
		result, err = q.db.ExecContext(ctx, editListBookPG, arg.BookID, arg.PurchaseHistoryID, arg.PurchaseNumber, arg.Qty, arg.TotalPrice)
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

type DeletePurchaseItemsTxParams struct {
	PurchaseNumber string
}

const lockRowPurchaseItemsByPurchaseNumberPG string = `
	SELECT 1
		FROM purchase_items
		WHERE purchase_number = $1
	FOR UPDATE
`
const fetchStatusPurchaseNumberPG string = `
	SELECT ph.status 
		FROM purchase_histories ph
	WHERE ph.purchase_number = $1
`

const deletePurchaseItemsPG string = `
	DELETE FROM purchase_items
	WHERE purchase_number = $1
`

const deletePurchaseHistoryPG string = `
	DELETE FROM purchase_histories
	WHERE purchase_number = $1
`

func (store *SQLStore) DeletePurchaseTx(ctx context.Context, args DeletePurchaseItemsTxParams) error {
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
		err = store.execTx(ctx, func(q *Queries) error {

			_, err = q.db.ExecContext(ctx, lockRowPurchaseItemsByPurchaseNumberPG, args.PurchaseNumber)

			if err != nil {
				return err
			}

			// check status
			err = q.db.QueryRowContext(ctx, fetchStatusPurchaseNumberPG, args.PurchaseNumber).Scan(
				&PurchaseHistories.Status,
			)
			if err != nil {
				return err
			}

			if PurchaseHistories.Status != "pending" {
				var err ErrStatusNotAcceptable
				err.Msg = fmt.Sprintf("status saat ini ==> '%s' sehingga tidak dapat proses penghapusan", PurchaseHistories.Status)
				return err
			}

			_, err = q.db.ExecContext(ctx, deletePurchaseItemsPG, args.PurchaseNumber)
			if err != nil {
				return err
			}

			_, err = q.db.ExecContext(ctx, deletePurchaseHistoryPG, args.PurchaseNumber)
			if err != nil {
				return err
			}

			return err
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

func (q *Queries) FinalizePurchase() error {
	return fmt.Errorf("not implemented yet %s", "not ready")
}

const lockRowBook string = `
	SELECT 1
	FROM books
	WHERE id = $1
	FOR UPDATE
`
const adjustStockBook string = `
 UPDATE books
 	SET stock_qty = $1
	WHERE id = $2
`

// adjust by increase or decrease, ketika dibeli decrease (-) ketika ditambah increase (+)
func (q *Queries) AdjustStockBook(ctx context.Context, bookID int, corrector int) error {
	var currentQty int
	var err error
	var rowAffected int64

	if corrector < 1 {
		// this struct is error and implement error interface
		return ErrNegativeNumber{
			Msg: fmt.Sprintf("nilai corrector negatif ==> %d silahkan sesuaikan", corrector),
		}
	}

	if bookID < 1 {
		var err ErrIDBukuTidakTerdaftar
		err.Msg = fmt.Sprintf("ID buku berikut tidak terdaftar ==> %d", bookID)
		return err
	}
	// lock row book id
	_, err = q.db.ExecContext(ctx, lockRowBook, bookID)
	if err != nil {
		return err
	}

	err = q.db.QueryRowContext(ctx, `select stock_qty from books where id = $1`, bookID).Scan(
		&currentQty,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			var err ErrIDBukuTidakTerdaftar
			err.Msg = fmt.Sprintf("ID buku berikut tidak terdaftar ==> %d", bookID)
			return err
		}

		return err
	}

	if currentQty < 1 {
		var err ErrStokBukuHabis
		err.Msg = fmt.Sprintf("stock saat ini tidak cukup ==> %d tidak dapat dikurang lagi", currentQty)
		return err
	}
	var resultAdjust = currentQty + corrector

	if resultAdjust < 0 {
		var err ErrNegativeNumber
		err.Msg = fmt.Sprintf("nilai corrector menyebabkan nilai stock negatif ==> %d silahkan sesuaikan", resultAdjust)
		return err
	}

	rowResult, err := q.db.ExecContext(ctx, adjustStockBook, resultAdjust, bookID)
	if err != nil {
		return err
	}

	rowAffected, err = rowResult.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected < 1 {
		var err ErrUpdateNolData
		err.Msg = fmt.Sprintf("error, data terupdate %d/!/", rowAffected)
		return err
	}

	return nil
}
