package util

import (
	"encoding/json"

	db "github.com/code/Banking-System/Db"
)

//TransactionPerform ..
func TransactionPerform(data db.Transaction) (string, bool) {

	dbref := db.Connect()
	if dbref == nil {
		return "Cannot connect to database during Transaction Perform", false
	}
	out, b := data.DoTransaction(dbref)
	if b == false {
		return out, false
	}
	return out, true

}

//DetailsPerform ..
func DetailsPerform(accNo int) string {

	dbref := db.Connect()
	if dbref == nil {
		return "Oops, cannot connect database."
	}
	det := &db.Transaction{
		AccFrom: accNo,
	}
	out, b, s := det.GetByAccFrom(dbref)
	if b == false {
		return s
	}
	result, _ := json.Marshal(out)
	return string(result)
}
