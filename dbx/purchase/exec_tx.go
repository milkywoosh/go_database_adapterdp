package purchase

import (
	"context"
	"fmt"
)

func (userstore *PurchaseStore) execTx(ctx context.Context, fn func(*PurchaseQueries) error) error {

	tx, err := userstore.connPool.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := NewPurchaseQueries(tx, userstore.dbtype)
	err = fn(q) // kinda callback
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
