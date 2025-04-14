package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

func main() {
	fmt.Println("teesss")
	tbl := "C##BOOK_STORE"
	pws := "Qwerty123."
	addr := "localhost:1521/orcl"
	libdir := "C:\\oracle\\instantclient_21_6\\windows"
	conn_str := fmt.Sprintf("user=%s password=%s connectString=%s libDir=%s", tbl, pws, addr, libdir)
	db, err := sql.Open("godror", conn_str)

	if err != nil {
		log.Fatalf("sql.Open failed: %v", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("DB Ping failed: %v", err)
		return
	}

	fmt.Println("✅ Successfully connected to Oracle!")
}
