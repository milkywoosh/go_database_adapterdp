package internal

import (
	"context"
	"fmt"
	"log"
)

func (userstore *UserStore) execTx(ctx context.Context, fn func(*UserQueries) error) error {

	tx, err := userstore.connPool.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("log err user store 1: %v", err)
		return err
	}

	q := NewUserQuery(tx, userstore.dbtype)
	err = fn(q) // kinda callback
	if err != nil {
		log.Printf("log err user store 2: %v", err)
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %w rb err: %w", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (purchasestore *PurchaseStore) execTx(ctx context.Context, fn func(*PurchaseQueries) error) error {

	tx, err := purchasestore.connPool.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("log err purchase store 1: %v", err)
		return err
	}

	q := NewPurchaseQueries(tx, purchasestore.dbtype)
	err = fn(q) // kinda callback
	if err != nil {
		log.Printf("log err purchase store 2: %v", err)
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %w rb err: %w", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
