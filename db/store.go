package db

import (
	"context"
	"database/sql"
)

type Store interface {
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	PurchaseBookTx(ctx context.Context, arg CreatePurchaseBookTxParams) (CreatePurchaseBookTxResult, error)
}

type SQLStore struct {
	connPool *sql.DB // *godror.Conn()
	*Queries
}

func NewStore(connPool *sql.DB, dbtype_arg string) *SQLStore{
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool, dbtype_arg),
	}
}
