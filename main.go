package main

import (
	"log"

	tr "github.com/code/Banking-System/Transaction"

	ac "github.com/code/Banking-System/Account"
	cr "github.com/code/Banking-System/Customer"
	"github.com/rightjoin/fuel"
)

//main ..
func main() {
	log.Println("Welcome to Pooraa Banking")
	server := fuel.NewServer()
	server.AddService(&tr.TransactionService{})
	server.AddService(&ac.AccountService{})
	server.AddService(&cr.Myservices{})

	// db.CreateTables()

	server.Run()

}
