package internal

import (
	"context"
	"database/sql"
	"fmt"
)

type UserQueries struct {
	dbtype string
	db     DBTX
}

func NewUserQuery(db_arg DBTX, dbtype string) *UserQueries {
	return &UserQueries{
		db:     db_arg,
		dbtype: dbtype,
	}
}

type UserQuerier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
}

func (q *UserQueries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {

	if q.dbtype == "ORACLE" {
		var i Users
		var err error
		_, err = q.db.ExecContext(ctx, createUserOra,
			arg.Username,
			arg.Email,
			arg.Firstname,
			arg.Lastname,
			arg.Password,
			sql.Out{Dest: &i.Username},
			sql.Out{Dest: &i.Email},
			sql.Out{Dest: &i.Firstname},
			sql.Out{Dest: &i.Lastname},
		)

		return i, err
	} else if q.dbtype == "POSTGRES" {
		var i Users
		var err error = q.db.QueryRowContext(ctx, createUserPG,
			arg.Username,
			arg.Email,
			arg.Firstname,
			arg.Lastname,
			arg.Password,
		).Scan(
			&i.Username,
			&i.Email,
			&i.Firstname,
			&i.Lastname,
		)

		if err != nil {
			return i, err
		}

		return i, err
	} else {
		var i Users
		var err error = fmt.Errorf("dbtype is not recognized")
		return i, err
	}

}

var _ UserQuerier = (*UserQueries)(nil)
