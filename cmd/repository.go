package main

import "database/sql"

func insertBalance(db *sql.DB, balance *Balance) {
	// dynamic insert
	insertDynStmt := `insert into "balance"("account_id", "balance") values($1, $2)`
	_, e := db.Exec(insertDynStmt, balance.accountId, balance.balance)
	CheckError(e)
}

// TODO implement select
func getAllBalances(db *sql.DB) {

}

// TODO implement update
func updateBalance(db *sql.DB, id int) {

}

// TODO implement delete
func deleteBalance(db *sql.DB, id int) {

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
