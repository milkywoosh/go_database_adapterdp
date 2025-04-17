package db

import (
	"context"
	"database/sql"
)

type Store interface {
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
}

type SQLStore struct {
	connPool *sql.DB // *godror.Conn
	*Queries
}

func NewStore(connPool *sql.DB) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
