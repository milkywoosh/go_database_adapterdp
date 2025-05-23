package db

import (
	"context"
	"database/sql"
)

// note misalnya tidak dalam transaksi, apakah akan terjadi update sebagian??? iyaa
// note jika akan melakukan UPDATE, INSERT disertai logic harus dalam *SQLStore execTx() function!!!

// need interface segregation! => untuk memisahkan transaction tiap REPO [UserRepo, PurchaseRepo ...]
type Store interface {
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	PurchaseBookTx(ctx context.Context, arg CreatePurchaseBookTxParams) (CreatePurchaseBookTxResult, error)
	EditListBookTx(ctx context.Context, arg EditBookToPurchaseParams) (int64, error)
	DeletePurchaseTx(ctx context.Context, arg DeletePurchaseItemsTxParams) error
}

// trial can be deleted anytime
type UserStore struct {
	Store
}

func NewUserStore(store Store) *UserStore {
	return &UserStore{
		store,
	}
}

type PurchaseStore struct {
	Store
}

func NewPurchaseStore(store Store) *PurchaseStore {
	return &PurchaseStore{
		store,
	}
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
