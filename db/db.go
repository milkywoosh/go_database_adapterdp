package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Queries struct {
	db DBTX
}

// constructor
func New(db_arg DBTX) *Queries {
	return &Queries{
		db: db_arg,
	}
}

func (q *Queries) WithTx() error {
	return fmt.Errorf("error unknown")
}

const createUser = `
INSERT INTO USERS (
  username,
  email
) VALUES (
  $1, $2
) RETURNING username,  email
`

// for now disini dulu!
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {
	// row := q.db.QueryRow(ctx, createUser,
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Email,
	)
	var i Users
	err := row.Scan(
		&i.Username,
		&i.Email,
	)
	return i, err
}
