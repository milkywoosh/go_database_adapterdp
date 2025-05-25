package purchase

import (
	"context"

	"github.com/luke_design_pattern/dbx"
)

type PurchaseQueries struct {
	dbtype string
	db     dbx.DBTX
}

func NewPurchaseQueries(db_arg dbx.DBTX, dbtype string) *PurchaseQueries {
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
