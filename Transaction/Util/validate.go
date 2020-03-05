package util

import (
	"fmt"
	"strconv"

	db "github.com/code/Banking-System/Db"

	reqModel "github.com/code/Banking-System/Transaction/Model"
)

//ExtractData ..
func ExtractData(res reqModel.TransactionRequest) (db.Transaction, bool, string) {

	var data db.Transaction
	val1, err1 := strconv.Atoi(res.AccFrom)
	if err1 != nil {
		return data, false, "invalid account number from "
	}
	data.AccFrom = val1
	val2, err2 := strconv.Atoi(res.AccTo)
	if err2 != nil {
		return data, false, "invalid account number from"
	}
	data.AccTo = val2
	fmt.Println(res.Amount)
	val3, err3 := strconv.ParseFloat(res.Amount, 64)
	if err3 != nil {
		fmt.Println(err3)
		return data, false, "invalid amount"
	}
	if val3 <= 0 {
		return data, false, "amount should be more than zero"
	}
	data.Amount = val3
	data.Type = res.Type
	return data, true, "Everything is fine."

}

//TransactionValidate ..
func TransactionValidate(data db.Transaction) (string, bool) {

	dbref := db.Connect()
	if dbref == nil {
		return "Can not connect to database.", false
	}
	var output []db.AccountDetail
	output = data.GetByAcc(dbref)
	if len(output) < 2 {
		return "account number not found", false
	}
	if data.CheckActive(dbref, data.AccFrom) == false {
		return "AccFrom is not active", false
	}
	if data.CheckActive(dbref, data.AccTo) == false {
		return "AccTo is not active", false
	}
	if data.Type == "credit" {
		st, b := data.CheckBalance(dbref, data.AccFrom)
		if b == false {
			return st, false
		}
	} else if data.Type == "debit" {
		st, b := data.CheckBalance(dbref, data.AccTo)
		if b == false {
			return st, false
		}
	}
	return "Validation Successfull", true

}

//DetailsValidate ..
func DetailsValidate(accNo int) (string, bool) {

	dbref := db.Connect()
	if dbref == nil {
		return "Can not connect to database.", false
	}
	if accNo <= 0 {
		return "Invalind account number", false
	}
	det := &db.Transaction{
		AccFrom: accNo,
	}
	s, b := det.CheckIfExist(dbref)
	if b == false {
		return s, false
	}
	return s, true
}
