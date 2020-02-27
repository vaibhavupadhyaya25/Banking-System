package main

import (
	"github.com/rightjoin/fuel"
)

type TransactionService struct {
	fuel.Service
	transactionDetails fuel.GET `route:"{accNo}"`
	performTransaction fuel.POST
}

func (s *TransactionService) PerformTransaction(a fuel.Aide) string {

	res := a.Post()
	data, b1, out1 := ExtractData(res)
	if b1 == false {
		return out1
	}
	out2, b2 := TransactionValidate(data)
	if b2 == false {
		return out2
	}
	out3, b3 := TransactionPerform(data)
	if b3 == false {
		return out3
	}
	return "Transaction Done Successfully."

}

func (s *TransactionService) TransactionDetails(accNo int) string {

	res, v := DetailsValidate(accNo)
	if v == false {
		return res
	}
	return DetailsPerform(accNo)

}
