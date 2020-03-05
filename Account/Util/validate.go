package util

import (
	"strconv"
	"time"

	model "github.com/code/Banking-System/Account/Model"
	db "github.com/code/Banking-System/Db"
)

//NegativeValidation ...
func NegativeValidation(accnum int) string {
	if accnum < 0 {
		return "Error account number cant be negative"
	}
	return ""
}

//PostValidate ...
func PostValidate(res model.AccountRequest) (db.AccountDetail, error) {

	var data db.AccountDetail
	val, err := strconv.Atoi(res.Cid)
	if err != nil {
		return data, err
	}
	data.Cid = val
	val1, err1 := strconv.Atoi(res.AccountNumber)
	if err1 != nil {
		return data, err1
	}
	data.AccountNumber = val1
	val2, err2 := strconv.ParseFloat(res.Balance, 64)
	if err2 != nil {
		return data, err2
	}
	data.Balance = val2
	val3, err3 := strconv.ParseFloat(res.Interest, 64)
	if err3 != nil {
		return data, err3
	}
	data.Interest = val3
	val4, err4 := strconv.ParseBool(res.IsActive)
	if err4 != nil {
		return data, err4
	}
	data.IsActive = val4
	data.AccountType = res.AccountType
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	return data, nil
}

//UpdateValidate ...
func UpdateValidate(res model.UpdateRequest) (db.AccountDetail, error) {

	var data db.AccountDetail
	val1, err1 := strconv.Atoi(res.AccountNumber)
	if err1 != nil {
		return data, err1
	}
	data.AccountNumber = val1
	val2, err2 := strconv.ParseFloat(res.Balance, 64)
	if err2 != nil {
		return data, err2
	}
	data.Balance = val2
	val4, err4 := strconv.ParseBool(res.IsActive)
	if err4 != nil {
		return data, err4
	}
	data.IsActive = val4
	data.UpdatedAt = time.Now()
	return data, nil
}
