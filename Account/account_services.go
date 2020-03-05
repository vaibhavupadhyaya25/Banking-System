package account

import (
	"encoding/json"

	model "github.com/code/Banking-System/Account/Model"
	util "github.com/code/Banking-System/Account/Util"
	db "github.com/code/Banking-System/Db"
	"github.com/rightjoin/fuel"
	"gopkg.in/go-playground/validator.v9"
)

//AccountService ...
type AccountService struct {
	fuel.Service
	createaccount fuel.POST   `route:"create"`
	readaccount   fuel.GET    `route:"read/{accnum}"`
	updateaccount fuel.PUT    `route:"update"`
	deleteaccount fuel.DELETE `route:"delete/{accnum}"`
}

//Readaccount ...
func (s *AccountService) Readaccount(accnum int) interface{} {

	pgdb := db.Connect()
	var account string
	account = util.ReadAccount(pgdb, accnum)
	return string(account)
}

//Createaccount ...
func (s *AccountService) Createaccount(a fuel.Aide) interface{} {

	req := a.Post()
	d, _ := json.Marshal(req)
	var res model.AccountRequest
	json.Unmarshal(d, &res)
	validate := validator.New()
	err1 := validate.Struct(res)
	if err1 != nil {
		return err1.Error()
	}
	data, err := util.PostValidate(res)
	if err != nil {
		return err
	}
	pgdb := db.Connect()
	util.Create(pgdb, data)
	var response model.AccountResponse
	response.Cid = data.Cid
	response.AccountNumber = data.AccountNumber
	response.IsActive = data.IsActive
	return response
}

//Updateaccount ...
func (s *AccountService) Updateaccount(a fuel.Aide) interface{} {
	req := a.Post()
	d, _ := json.Marshal(req)
	var res model.UpdateRequest
	json.Unmarshal(d, &res)
	validate := validator.New()
	err1 := validate.Struct(res)
	if err1 != nil {
		return err1.Error()
	}
	data, err := util.UpdateValidate(res)
	if err != nil {
		return err
	}
	pgdb := db.Connect()
	util.Update(pgdb, data)
	var response model.UpdateResponse
	response.Balance = data.Balance
	response.AccountNumber = data.AccountNumber
	response.IsActive = data.IsActive
	return response

}

//Deleteaccount ...
func (s *AccountService) Deleteaccount(accnum int) interface{} {

	pgdb := db.Connect()
	result := util.Delete(pgdb, accnum)
	return result
}
