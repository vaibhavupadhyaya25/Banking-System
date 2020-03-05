package Db

import (
	"log"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

// func CreateAccountDetailsTable(db *pg.DB) error {
// 	opts := &orm.CreateTableOptions{
// 		IfNotExists: true,
// 	}
// 	createErr := db.CreateTable(&AccountDetail{}, opts)
// 	if createErr != nil {
// 		log.Printf("Error while creating table AccountDetails, Reason: %v\n", createErr)
// 		return createErr
// 	}
// 	log.Printf("Table AccountDetails created successfully.\n")
// 	return nil
// }

// func CreateTableTransaction(db *pg.DB) error {

// 	opts := &orm.CreateTableOptions{
// 		IfNotExists: true,
// 	}
// 	err := db.CreateTable(&Transaction{}, opts)
// 	return err

// }
func CreateCustomerTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Customerinfo{}, opts)
	if createErr != nil {
		log.Printf("Error while creating tableItems, Reason: %v\n, createErr")
		return createErr
	}
	log.Printf("Customer Table is created\n")
	return nil
}
