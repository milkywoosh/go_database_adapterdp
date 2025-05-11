package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/godror/godror"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"

	dotenv "github.com/lpernett/godotenv"
	"github.com/luke_design_pattern/api"
	"github.com/luke_design_pattern/config"
	"github.com/luke_design_pattern/db"
)

func main() {
	err := dotenv.Load()
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	// tbl_ora := os.Getenv("ORA_DB_NAME")
	// pws_ora := os.Getenv("ORA_DB_PASSWORD")
	// addr_ora := os.Getenv("ORA_DB_CONN_STRING")
	// libdir_ora := os.Getenv("ORA_DB_LIB_DIR")

	tbl_pg := os.Getenv("PG_DB_NAME")
	host_pg := os.Getenv("PG_DB_HOST")
	port_pg := os.Getenv("PG_DB_PORT")
	username_pg := os.Getenv("PG_DB_USERNAME")
	pws_pg := os.Getenv("PG_DB_PASSWORD")
	addr_pg := os.Getenv("PG_DB_CONN_STRING")
	ssl_pg := os.Getenv("PG_DB_SSL")
	libdir_pg := os.Getenv("PG_DB_LIB_DIR")

	// new_ora_conn, err := config.NewCredential("ORACLE", "", "", tbl_ora, "", pws_ora, addr_ora, libdir_ora, "")
	// if err != nil {
	// 	log.Fatalf("ORA: %v ", err)
	// 	return
	// }
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

	// ora_credential, err := new_ora_conn.GetConnectionString()
	// if err != nil {
	// 	log.Fatalf("===> ora get conn string:  %v", err)
	// 	return
	// }
	pg_credential, err := new_pg_conn.GetConnectionString()
	if err != nil {
		log.Fatalf("===> pg get conn string:  %v", err)
		return
	}

	// db_ora, err := sql.Open("godror", ora_credential)
	// if err != nil {
	// 	log.Fatalf("sql.Open godror failed: %v", err)
	// 	return
	// }

	// note => use stdlib
	// db_pg, err := sql.Open("postgres", pg_credential)
	configpg, err := pgx.ParseConfig(pg_credential)
	if err != nil {
		log.Fatalf("sql.Open pg failed: %v", err)
		return
	}
	db_pg := stdlib.OpenDB(*configpg)

	// defer db_ora.Close()
	defer db_pg.Close()

	// if err := db_ora.Ping(); err != nil {
	// 	log.Fatalf("DB Ping ora failed: %v", err)
	// 	return
	// }
	if err := db_pg.Ping(); err != nil {
		log.Fatalf("DB Ping pg failed: %v", err)
		return
	}

	pg_sql := db.NewStore(db_pg, "POSTGRES")
	server_api, err := api.NewServer(new_pg_conn, pg_sql)
	if err != nil {
		log.Fatalf("api.NewServer failed: %v", err)
		os.Exit(1)
	}

	// note: port apps and PORT to DB is different
	server_api.Start(fmt.Sprintf(":%s", "8000"))

}
