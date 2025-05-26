package user

import (
	"context"

	"github.com/luke_design_pattern/dbx"
)

type UserQueries struct {
	dbtype string
	db     dbx.DBTX
}

func NewUserQuery(db_arg dbx.DBTX, dbtype string) *UserQueries {
	return &UserQueries{
		db:     db_arg,
		dbtype: dbtype,
	}
}

type UserQuerier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
}

var _ UserQuerier = (*UserQueries)(nil)
