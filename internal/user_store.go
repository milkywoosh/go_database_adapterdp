package internal

import (
	"context"
	"database/sql"
)

type UserStoreTx interface {
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	AssignRoleTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
}

type UserStore struct {
	connPool     *sql.DB
	*UserQueries // the sqlc-generated querier
}

func NewUserStore(db *sql.DB, dbtype string) *UserStore {
	return &UserStore{
		connPool:    db,
		UserQueries: NewUserQuery(db, dbtype),
	}
}

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
func (store *UserStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *UserQueries) error {
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

func (store *UserStore) AssignRoleTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	return CreateUserTxResult{}, nil
}
