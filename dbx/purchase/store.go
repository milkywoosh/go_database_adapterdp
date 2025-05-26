package purchase

import (
	"context"
	"database/sql"
)

type PurchaseStore struct {
	connPool         *sql.DB
	*PurchaseQueries // the sqlc-generated querier
}

func NewPurchaseStore(dbtype string, db *sql.DB) PurchaseTransaction {
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
