package main

import (
	"log"

	"github.com/rightjoin/fuel"
)

func main() {
	log.Println("Welcome to Pooraa Banking")
	//router := mux.NewRouter()
	//log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
	//handlers.AllowedOrigins([]string{"*"})
	server := fuel.NewServer()
	server.AddService(&AccountService{})
	server.AddService(&TransactionService{})
	server.AddService(&myservices{})

	server.Run()
}
