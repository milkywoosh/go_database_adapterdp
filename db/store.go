package db

import (
	"context"
	"database/sql"
)

// note misalnya tidak dalam transaksi, apakah akan terjadi update sebagian??? iyaa
// note jika akan melakukan UPDATE, INSERT disertai logic harus dalam *SQLStore execTx() function!!!

type Store interface {
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	PurchaseBookTx(ctx context.Context, arg CreatePurchaseBookTxParams) (CreatePurchaseBookTxResult, error)
	EditListBookTx(ctx context.Context, arg EditBookToPurchaseParams) (int64, error)
	DeletePurchaseTx(ctx context.Context, arg DeletePurchaseItemsTxParams) error
}

type SQLStore struct {
	connPool *sql.DB // *godror.Conn() atau pgx.Conn() depend definisi dbtype
	*Queries
}

func NewStore(connPool *sql.DB, dbtype_arg string) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool, dbtype_arg),
	}
}
