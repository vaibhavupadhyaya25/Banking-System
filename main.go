package main

import (
	"log"

	cr "github.com/Banking-System/Customer"
	//db "github.com/Banking-System/Db"
	"github.com/rightjoin/fuel"
)

func main() {
	log.Println("Welcome to Pooraa Banking")
	server := fuel.NewServer()
	//server.AddService(&AccountService{})
	//server.AddService(&TransactionService{})
	server.AddService(&cr.Myservices{})
	//db.CreateTables()

	server.Run()

}
