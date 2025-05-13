package db

import (
	"context"
	"fmt"
	"log"
)

// ExecTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.connPool.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("check log error 1 ==> %v", r)
			tx.Rollback()
			err = fmt.Errorf("tx err: %v", r)
		} else if err != nil {
			log.Printf("check log error 2 ==> %v", err)
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("tx err: %v rb err: %v", err, rbErr)
			}
		} else {
			err = tx.Commit()
			log.Printf("check log error 3 ==> %v", err)
		}
	}()

	// initiate q to construct *Queries
	q := New(tx, store.Queries.dbtype)

	// pass q into fn to satisfy this [fn func(*Queries) error]
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v rb err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}
