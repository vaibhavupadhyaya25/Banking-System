package db

import (
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
