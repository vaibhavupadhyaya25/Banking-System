package db

import (
	"fmt"

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
	err := CreateTableTransaction(db)
	if err != nil {
		fmt.Println("Opps something went wrong. Reason: ", err)
	}
	return db
}
