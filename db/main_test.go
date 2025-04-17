package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/godror/godror"
	dotenv "github.com/lpernett/godotenv"
	"github.com/luke_design_pattern/config"
)

var testStore Store

func TestMain(m *testing.M) {

	err := dotenv.Load("../.env")
	if err != nil {

		log.Fatalf("testing err: %v", err)
		return
	}
	tbl := os.Getenv("ORA_DB_NAME")
	log.Printf("testing check ==> %s", tbl)
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
		log.Fatalf("===> 1:  %v", err)
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

	log.Printf("load testing.M: %s", "success connect!")

	testStore = NewStore(db)
	os.Exit(m.Run())
}
