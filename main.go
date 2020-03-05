package main

import (
	"log"

	tr "github.com/code/Banking-System/Transaction"

	db "github.com/code/Banking-System/Db"

	"github.com/rightjoin/fuel"
)

//main ..
func main() {
	log.Println("Welcome to Pooraa Banking")
	server := fuel.NewServer()
	server.AddService(&tr.TransactionService{})

	db.CreateTables()

	server.Run()
}
