package internal

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// DBTX interface => *sql.DB => karena implement 3 signatures
type Queries struct {
	dbtype string
	db     DBTX
}

// constructor
func New(db_arg DBTX, dbtype_arg string) *Queries {
	return &Queries{
		dbtype: dbtype_arg,
		db:     db_arg,
	}
}

type OraAdapter struct {
	Adaptee *SQLStore // di-implementasi *SQLStore
}

func NewOra(db_arg *sql.DB) *OraAdapter {
	return &OraAdapter{
		Adaptee: NewStore(db_arg, "ORACLE"),
	}
}

func (oa *OraAdapter) GetConn() *SQLStore {
	return oa.Adaptee
}

type PgAdapter struct {
	Adaptee *SQLStore
}

func NewPG(db_arg *sql.DB) *PgAdapter {
	return &PgAdapter{
		Adaptee: NewStore(db_arg, "POSTGRES"),
	}
}

func (oa *PgAdapter) GetConn() *SQLStore {
	return oa.Adaptee
}
