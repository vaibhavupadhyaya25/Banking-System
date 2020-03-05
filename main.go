package main

import (
	"log"

	ac "github.com/Banking-System/Account"
	"github.com/rightjoin/fuel"
)

func main() {
	log.Println("Welcome to Pooraa Banking")
	server := fuel.NewServer()
	server.AddService(&ac.AccountService{})
	//server.AddService(&TransactionService{})
	//server.AddService(&myservices{})

	server.Run()
}
