package db

import (
	"context"

)

type CreateUserTxParams struct {
	CreateUserParams
	
	AfterCreate func(user Users) error // note: diisi function APAPUN yg penting passing argument tipe Users dan return Error!
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// means: AfterCreate is a function field (callback) that takes a Users object and returns an error.
		// It's a callback that's run after the user has been created in the database, inside the transaction.
		result.Users, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}

		return arg.AfterCreate(result.Users)
	})

	return result, err
}
