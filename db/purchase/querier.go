package db

import (
	"context"
	"database/sql"

	"github.com/luke_design_pattern/db"
)

type PurchaseQueries struct {
	dbtype string
	db     db.DBTX
}

func NewPurchaseQueries(dbtype string, db_arg *sql.DB) *PurchaseQueries {
	return &PurchaseQueries{
		dbtype: dbtype,
		db:     db_arg,
	}
}

type PurchaseQuerier interface {
	CreatePurchaseHistory(ctx context.Context, arg CreatePurchaseHistoryParams) (PurchaseHistory, error)
	AddListBook(ctx context.Context, arg CreateBookToPurchaseParams) (BookToPurchase, error)
	FinalizePurchase() error
	AdjustStockBook(ctx context.Context, bookID int, corrector int) error
}
