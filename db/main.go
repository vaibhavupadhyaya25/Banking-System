package db

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
)

func Connect() *pg.DB {

	opts := &pg.Options{
		User:     "postgres",
		Password: "123456",
		Addr:     "10.1.7.110:5432",
		Database: "bank",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect \n")
		os.Exit(100)
	}
	log.Printf("Connected Succesfully \n")
	CreateAccountDetailsTable(db)
	CreateBankDetailsTable(db)
	CreateTableTransaction(db)
	CreateCustomerTable(db)
	return db

}
