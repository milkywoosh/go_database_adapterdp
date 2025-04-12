package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "log"

	_ "github.com/godror/godror"
)

func main() {

	// note: FIX GOOD CONNECTION TO ORACLE
	db, err := sql.Open("godror", `user="C##BOOK_STORE" password="Qwerty123." connectString="localhost:1521/ORCL" libDir="C:/instantclient_21.6/windows"`)
	db, err := sql.Open("godror", `user="C##BOOK_STORE" password="Qwerty123." connectString="localhost:1521/ORCL" libDir="C:/instantclient_21.6/windows"`)

	//budimanokky93.medium.com/golang-easy-fetch-millions-of-data-using-concurrent-80716595e674
	if err != nil {
		cek := fmt.Sprintf("%s: ==> %v", "tess: ", err)
		log.Fatal(cek)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	if err != nil {
		err := fmt.Sprintf("%s %s", "err", err.Error())
		panic(err)
	}

	row, err := db.Query("SELECT 'Hello World' FROM dual")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var message string

		row.Scan(&message)

		fmt.Println(message)
	}

	// https://budimanokky93.medium.com/golang-easy-fetch-millions-of-data-using-concurrent-80716595e674

	// rows, err := db.Query("SELECT ID FROM BU_ROLES WHERE ROWNUM <= 5")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer rows.Close()

	// doDBThingsThroughInstantClient(localDB)
}
