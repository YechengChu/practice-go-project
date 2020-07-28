package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"io/ioutil"

	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)

var db *sql.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func getDbInfo() {
	var userHost string
	fmt.Print("Server[localhost]: ")
	fmt.Scanln(&userHost)
	if strings.TrimSpace(userHost) != "" {
		host = userHost
	}

	var userDB string
	fmt.Print("Database[postgres]: ")
	fmt.Scanln(&userDB)
	if strings.TrimSpace(userDB) != "" {
		dbname = userDB
	}

	var userPort string
	fmt.Print("Port[5432]: ")
	fmt.Scanln(&userPort)
	if strings.TrimSpace(userPort) != "" {
		intPort, err := strconv.Atoi(userPort)
		checkErr(err)
		port = intPort
	}

	var userName string
	fmt.Print("Username[postgres]: ")
	fmt.Scanln(&userName)
	if strings.TrimSpace(userName) != "" {
		user = userName
	}

	fmt.Print("Password for user postgres: ")
	fmt.Scanln(&password)
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
	// fmt.Printf("info has type %T\n", info)
	for info.Next() {
		err = info.Scan(&pass)
		checkErr(err)
		hasAccount = true
	}
	return
} // query

func initDB() {
	getDbInfo()
	openDBSQL := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	sqlBytes, err := ioutil.ReadFile("createTable.sql")
	checkErr(err)
	sqlCommand := string(sqlBytes)

	var errOpenDB error
	db, errOpenDB = sql.Open("postgres", openDBSQL)
	// fmt.Printf("db has type of %T\n", db)
	checkErr(errOpenDB)
	_, err = db.Exec(sqlCommand)
	checkErr(err)
}

func closeDB() {
	db.Close()
}
