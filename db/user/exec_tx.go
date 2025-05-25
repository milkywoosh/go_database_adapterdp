package db

import (
	"context"
	"fmt"
)

func (userstore *UserStore) execTx(ctx context.Context, fn func(*UserQueries) error) error {

	tx, err := userstore.connPool.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := NewUserQuery(tx, userstore.dbtype)
	err = fn(q) // kinda callback
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
