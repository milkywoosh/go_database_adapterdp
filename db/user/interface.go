package db

import "context"

type UserTransaction interface {
	UserQuerier // lower level, abstraction for DBTX query
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	AssignRoleTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
}
