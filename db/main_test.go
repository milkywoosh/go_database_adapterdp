package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/godror/godror"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	dotenv "github.com/lpernett/godotenv"
	"github.com/luke_design_pattern/config"
)

var testStoreOra Store
var testStorePG Store

func TestMain(m *testing.M) {

	err := dotenv.Load("../.env")
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	tbl_ora := os.Getenv("ORA_DB_NAME")
	pws_ora := os.Getenv("ORA_DB_PASSWORD")
	addr_ora := os.Getenv("ORA_DB_CONN_STRING")
	libdir_ora := os.Getenv("ORA_DB_LIB_DIR")

	tbl_pg := os.Getenv("PG_DB_NAME")
	host_pg := os.Getenv("PG_DB_HOST")
	port_pg := os.Getenv("PG_DB_PORT")
	username_pg := os.Getenv("PG_DB_USERNAME")
	pws_pg := os.Getenv("PG_DB_PASSWORD")
	addr_pg := os.Getenv("PG_DB_CONN_STRING")
	ssl_pg := os.Getenv("PG_DB_SSL")
	libdir_pg := os.Getenv("PG_DB_LIB_DIR")

	new_ora_conn, err := config.NewCredential("ORACLE", "", "", tbl_ora, "", pws_ora, addr_ora, libdir_ora, "")
	if err != nil {
		log.Fatalf("ORA: %v ", err)
		return
	}
	// POSTGRES_HOST=127.0.0.1
	// POSTGRES_USERNAME=postgres
	// POSTGRES_PASSWORD=postgres
	// POSTGRES_DBNAME=toko_buku_online_nextjs
	// POSTGRES_PORT=5434
	// #5434 pg docker

	new_pg_conn, err := config.NewCredential("POSTGRES", host_pg, port_pg, tbl_pg, username_pg, pws_pg, addr_pg, libdir_pg, ssl_pg)
	if err != nil {
		log.Fatalf("PG: %v ", err)
		return
	}

	ora_credential, err := new_ora_conn.GetConnectionString()
	if err != nil {
		log.Fatalf("===> ora get conn string:  %v", err)
		return
	}
	pg_credential, err := new_pg_conn.GetConnectionString()
	if err != nil {
		log.Fatalf("===> pg get conn string:  %v", err)
		return
	}
	db_ora, err := sql.Open("godror", ora_credential)
	if err != nil {
		log.Fatalf("sql.Open godror failed: %v", err)
		return
	}

	// note => use stdlib
	// db_pg, err := sql.Open("postgres", pg_credential)
	configpg, err := pgx.ParseConfig(pg_credential)
	if err != nil {
		log.Fatalf("sql.Open pg failed: %v", err)
		return
	}
	db_pg := stdlib.OpenDB(*configpg)
	defer db_ora.Close()
	defer db_pg.Close()

	if err := db_ora.Ping(); err != nil {
		log.Fatalf("DB Ping ora failed: %v", err)
		return
	}
	if err := db_pg.Ping(); err != nil {
		log.Fatalf("DB Ping pg failed: %v", err)
		return
	}

	testStoreOra = NewOra(db_ora).GetConn()
	testStorePG = NewPG(db_pg).GetConn()

	os.Exit(m.Run())
}
