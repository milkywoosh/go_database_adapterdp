package purchase

import "context"

type PurchaseTransaction interface {
	PurchaseBookTx(ctx context.Context, arg CreatePurchaseBookTxParams) (CreatePurchaseBookTxResult, error)
	EditListBookTx(ctx context.Context, arg EditBookToPurchaseParams) (int64, error)
	DeletePurchaseTx(ctx context.Context, arg DeletePurchaseItemsTxParams) error
}
