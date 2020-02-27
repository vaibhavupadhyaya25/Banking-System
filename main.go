package main

import (
	db "./db"
	"github.com/rightjoin/fuel"
)

func main() {
	db.Connect()
	// SaveTransaction(dbref)

	server := fuel.NewServer()
	server.AddService(&TransactionService{})
	server.Run()

}

// func SaveTransaction(dbref *pg.DB) {

// 	newData1 := &db.Transaction{
// 		AccFrom: 12345678912,
// 		AccTo:   12345678913,
// 		Time:    time.Now(),
// 		Amount:  500,
// 		Type:    "credit",
// 	}
// 	newData2 := &db.Transaction{
// 		AccFrom: 12345678913,
// 		AccTo:   12345678912,
// 		Time:    time.Now(),
// 		Amount:  500,
// 		Type:    "debit",
// 	}
// 	totalItems := []*db.Transaction{newData1, newData2}
// 	newData1.SaveMultiple(dbref, totalItems)
// }
