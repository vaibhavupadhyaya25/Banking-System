package transaction

import (
	"encoding/json"

	util "github.com/code/Banking-System/Transaction/Util"

	reqModel "github.com/code/Banking-System/Transaction/Model"

	"github.com/rightjoin/fuel"
	"gopkg.in/go-playground/validator.v9"
)

//TransactionService ..
type TransactionService struct {
	fuel.Service
	transactionDetails fuel.GET `route:"{accNo}"`
	performTransaction fuel.POST
}

//PerformTransaction ..
func (s *TransactionService) PerformTransaction(a fuel.Aide) string {

	res := a.Post()
	d, _ := json.Marshal(res)
	var result reqModel.TransactionRequest
	json.Unmarshal(d, &result)

	validate := validator.New()
	err1 := validate.Struct(result)
	if err1 != nil {
		return string(err1.Error())
	}

	data, b1, out1 := util.ExtractData(result)
	if b1 == false {
		return out1
	}

	out2, b2 := util.TransactionValidate(data)
	if b2 == false {
		return out2
	}
	out3, b3 := util.TransactionPerform(data)
	if b3 == false {
		return out3
	}
	return "Transaction Done Successfully."

}

//TransactionDetails ..
func (s *TransactionService) TransactionDetails(accNo int) string {

	res, v := util.DetailsValidate(accNo)
	if v == false {
		return res
	}
	return util.DetailsPerform(accNo)

}
