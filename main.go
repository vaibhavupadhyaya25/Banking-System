package main

import (
	"log"

	"github.com/rightjoin/fuel"
)

func main() {
	log.Println("Welcome to Pooraa Banking")
	server := fuel.NewServer()
	server.AddService(&AccountService{})
	server.AddService(&TransactionService{})
	server.AddService(&myservices{})

	server.Run()
}
