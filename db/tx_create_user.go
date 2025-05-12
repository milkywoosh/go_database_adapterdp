package db

import (
	"context"
	"database/sql"
	"fmt"
)

// NOTE: khusus query ORACLE => placeholder :1, :2, :3
// NOTE: performance consideration param using slice string{"field1", "field2", "field3"}
const createUserOra string = `
INSERT INTO USERS (
  username,
  email,
  firstname,
  lastname,
  password
) VALUES (
  :1, :2, :3, :4, :5
) RETURNING username, email, firstname, lastname INTO  :6, :7, :8, :9
`
const createUserPG string = `
INSERT INTO USERS (
  username,
  email,
  firstname,
  lastname,
  password
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING username, email, firstname, lastname
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

		// return arg.AfterCreate(result.Users)
		return nil

	})

	return result, err
}
