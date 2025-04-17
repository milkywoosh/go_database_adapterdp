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

// DBTX interface => *sql.DB => karena implement 3 signatures
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

// NOTE: khusus query ORACLE => placeholder :1, :2, :3
// NOTE: performance consideration param using slice string{"field1", "field2", "field3"}
const createUser = `
INSERT INTO USERS (
  username,
  email,
  firstname,
  lastname
) VALUES (
  :1, :2, :3, :4
) RETURNING username, email, firstname, lastname INTO :5, :6, :7, :8
`

/*
	var id int64
	_, err := db.ExecContext(ctx,
  `INSERT INTO users (username) VALUES (:1) RETURNING id INTO :2`,
  "alice",
  sql.Out{Dest: &id},
)

*/
// for now disini dulu!
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {
	var i Users

	_, err := q.db.ExecContext(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
		sql.Out{Dest: &i.Username},
		sql.Out{Dest: &i.Email},
		sql.Out{Dest: &i.Firstname},
		sql.Out{Dest: &i.Lastname},
	)

	return i, err
}
