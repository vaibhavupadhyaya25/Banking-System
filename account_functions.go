package main

import (
	"strconv"
	"time"

	db "./db"
	"github.com/go-pg/pg"
)

func Create(dbRef *pg.DB, res db.AccountDetail) {

	res.Save(dbRef)
}

func ReadAccount(dbRef *pg.DB, accnum int) db.AccountDetail {
	newAD1 := &db.AccountDetail{
		AccountNumber: accnum,
	}
	item := newAD1.Readaccount(dbRef)
	return item
}

func Delete(dbRef *pg.DB, accnum int) {
	newAD1 := &db.AccountDetail{
		AccountNumber: accnum,
	}
	newAD1.Delete(dbRef)
}

func Update(dbRef *pg.DB, res db.AccountDetail) {

	res.Update(dbRef)
}
func PostValidate(res map[string]string) interface{} {

	var abc db.AccountDetail
	val1, err1 := strconv.Atoi(res["AccountNumber"])
	if err1 != nil {
		return "Invalid Account Number"
	}

	abc.Cid, _ = strconv.Atoi(res["Cid"])
	abc.AccountNumber = val1
	val2, err2 := strconv.Atoi(res["Balance"])
	if err2 != nil {
		return "Invalid Balance"
	}
	abc.Balance = val2
	val3, err3 := strconv.Atoi(res["Interest"])
	if err3 != nil {
		return "Invalid Interest"
	}
	abc.Interest = val3
	val4, err4 := strconv.ParseBool(res["IsActive"])
	if err4 != nil {
		return "Invalid Disable"
	}
	abc.IsActive = val4
	abc.AccountType = res["AccountType"]
	abc.CreatedAt = time.Now()
	abc.UpdatedAt = time.Now()
	return abc
}
