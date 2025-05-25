package db

import (
	"database/sql"
)

type UserStore struct {
	connPool     *sql.DB
	*UserQueries // the sqlc-generated querier
}

func NewUserStore(dbtype string, db *sql.DB) UserTransaction {
	return &UserStore{
		connPool:    db,
		UserQueries: NewUserQuery(db, dbtype),
	}
}
