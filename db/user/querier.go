package db

import (
	"context"

	"github.com/luke_design_pattern/db"
)

type UserQueries struct {
	dbtype string
	db     db.DBTX
}

func NewUserQuery(db_arg db.DBTX, dbtype string) *UserQueries {
	return &UserQueries{
		db:     db_arg,
		dbtype: dbtype,
	}
}

type UserQuerier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
}

var _ UserQuerier = (*UserQueries)(nil)
