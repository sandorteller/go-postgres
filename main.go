package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "postgres"
)

type Balance struct {
	accountId int
	balance   int
}

func main() {
	// connection string
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlConn)
	CheckError(err)

	// close database
	//defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	var balance = Balance{1, 123}
	insertBalance(db, &balance)
}

func insertBalance(db *sql.DB, balance *Balance) {
	// dynamic insert
	insertDynStmt := `insert into "balance"("account_id", "balance") values($1, $2)`
	_, e := db.Exec(insertDynStmt, balance.accountId, balance.balance)
	CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
