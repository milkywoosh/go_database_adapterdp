package internal

import (
	"database/sql"
)

// note misalnya tidak dalam transaksi, apakah akan terjadi update sebagian??? iyaa
// note jika akan melakukan UPDATE, INSERT disertai logic harus dalam *SQLStore execTx() function!!!

// need interface segregation! => untuk memisahkan transaction tiap REPO [UserRepo, PurchaseRepo ...]

type SQLStore struct {
	// connPool *sql.DB // *godror.Conn() atau pgx.Conn() depend definisi dbtype
	*UserStore
	*PurchaseStore
}

func NewStore(connPool *sql.DB, dbtype_arg string) *SQLStore {
	return &SQLStore{
		UserStore:     NewUserStore(connPool, dbtype_arg),
		PurchaseStore: NewPurchaseStore(connPool, dbtype_arg),
	}
}
