package migrations

import (
	"database/sql"
	"log"
	"point-of-sale/config"
)

func Run() {
	// Migrate Tables
	migrate(config.DB, Products)
	migrate(config.DB, Transactions)
	migrate(config.DB, TransactionsDetail)
}

func migrate(dbDriver *sql.DB, query string) {
	statement, err := dbDriver.Prepare(query)
	if err == nil {
		_, creationError := statement.Exec()
		if creationError == nil {
			log.Println("Table Created Succesfully")
		} else {
			log.Println(creationError.Error())
		}
	} else {
		log.Println(err.Error())
	}
}
