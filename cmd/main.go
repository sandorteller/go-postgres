package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	var balance = Balance{1, 123}
	insertBalance(db, &balance)
	updateBalance(db, balance.accountId)
	getAllBalances(db)
	deleteBalance(db, balance.accountId)
}
