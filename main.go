package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	dotenv "github.com/lpernett/godotenv"
	"github.com/luke_design_pattern/config"

	_ "github.com/godror/godror"
)

func main() {
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	tbl := os.Getenv("ORA_DB_NAME")
	pws := os.Getenv("ORA_DB_PASSWORD")
	addr := os.Getenv("ORA_DB_CONN_STRING")
	libdir := os.Getenv("ORA_DB_LIB_DIR")

	new_ora_conn, err := config.NewCredential("ORACLE", tbl, pws, addr, libdir)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	ora_credential, err := new_ora_conn.GetConnectionString()
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	db, err := sql.Open("godror", ora_credential)

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
