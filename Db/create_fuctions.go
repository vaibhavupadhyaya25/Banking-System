package db

import (
	"log"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

//CreateTableTransaction ..
func CreateTableTransaction(db *pg.DB) error {

	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&Transaction{}, opts)
	return err
}

//CreateAccountDetailsTable ...
func CreateAccountDetailsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&AccountDetail{}, opts)
	if createErr != nil {
		log.Printf("Error while creating table AccountDetails, Reason: %v\n", createErr)
		return createErr
	}
	log.Printf("Table AccountDetails created successfully.\n")
	return nil
}
