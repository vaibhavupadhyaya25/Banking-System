package main

import (
	"encoding/json"

	db "./db"
	"github.com/rightjoin/fuel"
)

type AccountService struct {
	fuel.Service
	createaccount fuel.POST   `route:"create"`
	readaccount   fuel.GET    `route:"read/{accnum}"`
	updateaccount fuel.PUT    `route:"update"`
	deleteaccount fuel.DELETE `route:"delete/{accnum}"`
}

func (s *AccountService) Readaccount(accnum int) interface{} {

	//add validation
	pgdb := db.Connect()
	//Read(pgdb, accnum)
	var account db.AccountDetail
	account = Read(pgdb, accnum)
	bs, err := json.Marshal(account)
	if err != nil {
		return err
	}
	return string(bs)
}

func (s *AccountService) Createaccount(a fuel.Aide) interface{} {

	res := a.Post()
	data := PostValidate(res)

	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	var abc db.AccountDetail
	json.Unmarshal(bs, &abc)

	pgdb := db.Connect()
	Create(pgdb, abc)
	return data
}
func (s *AccountService) Updateaccount(a fuel.Aide) interface{} {

	res := a.Post()
	data := PostValidate(res)
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	var abc db.AccountDetail
	json.Unmarshal(bs, &abc)

	pgdb := db.Connect()
	Update(pgdb, abc)
	return data
}

func (s *AccountService) Deleteaccount(accnum int) interface{} {

	//add validation
	pgdb := db.Connect()
	Delete(pgdb, accnum)
	return "Successfully deleted"
}
