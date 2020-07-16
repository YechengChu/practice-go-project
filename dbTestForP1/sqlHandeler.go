package main

import (
	"database/sql"
	"fmt"

	"io/ioutil"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123aaa"
	dbname   = "postgres"
)

var db *sql.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func insert(givenAcc string, givenPass string) {
	psqlInfo := fmt.Sprintf("INSERT INTO users(account, password) VALUES('%s','%s');", givenAcc, givenPass)
	_, err := db.Exec(psqlInfo)
	checkErr(err)
}

func query(givenAcc string) (hasAccount bool, pass string) {
	hasAccount = false
	pass = ""
	psqlInfo := fmt.Sprintf("SELECT password FROM users WHERE account='%s';", givenAcc)
	info, err := db.Query(psqlInfo)
	checkErr(err)
	fmt.Printf("info has type %T\n", info)
	for info.Next() {
		err = info.Scan(&pass)
		checkErr(err)
		hasAccount = true
	}
	return
} // query

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	sqlBytes, err := ioutil.ReadFile("createTable.sql")
	checkErr(err)
	sqlCommand := string(sqlBytes)

	var errOpenDB error
	db, errOpenDB = sql.Open("postgres", psqlInfo)
	fmt.Printf("db has type of %T\n", db)
	checkErr(errOpenDB)

}
