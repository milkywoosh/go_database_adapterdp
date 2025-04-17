package db

import "fmt"

type DBTX interface {
	Exec()
	Query()
	QueryRow()
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
