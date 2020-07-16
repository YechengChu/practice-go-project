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

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	sqlBytes, err := ioutil.ReadFile("createTable.sql")
	if err != nil {
		panic(err)
	}
	sqlCommand := string(sqlBytes)

	var errOpenDB error
	db, errOpenDB = sql.Open("postgres", psqlInfo)
	fmt.Printf("db has type of %T\n", db)
	checkErr(errOpenDB)
	defer db.Close()

	_, err = db.Exec(sqlCommand)
	checkErr(err)
	insert("szcyc001@123.com", "Asdfghjkl")
	insert("szcyc003@111.com", "123aaa")
	hasAcc, password := query("szcyc001@123.com")
	hasAcc2, password2 := query("szcyc001@163.com")
	fmt.Printf("%v, %s\n", hasAcc, password)
	fmt.Printf("%v, %s\n", hasAcc2, password2)
	fmt.Println("Everything done!")
}
