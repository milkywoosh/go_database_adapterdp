package db

import "context"

type CreatePurchaseBookTxParams struct{}
type CreatePurchaseBookTxResult struct{}

func (store *SQLStore) PurchaseBookTx(ctx context.Context, arg CreatePurchaseBookTxParams) (CreatePurchaseBookTxResult, error) {
	// var result_purchase_history PurchaseHistory
	// var err error
	// err = store.execTx(ctx, func(q *Queries) error {
	// 	result_purchase_history, err = q.CreatePurchaseHistory(ctx, CreatePurchaseHistoryParams{})

	// 	return err
	// })

	return CreatePurchaseBookTxResult{}, nil
}
