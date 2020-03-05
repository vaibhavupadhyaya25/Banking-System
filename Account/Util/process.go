package util

import (
	"encoding/json"

	db "github.com/code/Banking-System/Db"
	pg "github.com/go-pg/pg"
)

//Create ...
func Create(dbRef *pg.DB, res db.AccountDetail) {

	res.Save(dbRef)
}

//ReadAccount ...
func ReadAccount(dbRef *pg.DB, accnum int) string {
	err := NegativeValidation(accnum)
	if err != "" {
		return err
	} else {
		//fmt.Println(err)
		newAD1 := &db.AccountDetail{
			AccountNumber: accnum,
		}
		item := newAD1.Readaccount(dbRef)
		result, _ := json.Marshal(item)
		return string(result)
	}
}

//Delete ...
func Delete(dbRef *pg.DB, accnum int) string {
	err := NegativeValidation(accnum)
	if err != "" {
		return err
	} else {
		newAD1 := &db.AccountDetail{
			AccountNumber: accnum,
		}
		newAD1.Delete(dbRef)
		return "Successfully Deleted"
	}
}

//Update ..
func Update(dbRef *pg.DB, res db.AccountDetail) {

	res.Update(dbRef)
}
