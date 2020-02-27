package db

import (
	"fmt"
	"time"

	pg "github.com/go-pg/pg"
)

func (t *Transaction) DoTransaction(db *pg.DB) (string, bool) {

	if t.Type == "credit" {
		tx, txerr := db.Begin()
		if txerr != nil {
			return "Cannot go in begin block", false
		}
		t.Time = time.Now()
		_, updateErr1 := tx.Model(t).Insert()
		if updateErr1 != nil {
			tx.Rollback()
			return "Cannot update in transaction table (1)", false
		}
		data := Transaction{
			Type:    "debit",
			Amount:  t.Amount,
			AccFrom: t.AccTo,
			AccTo:   t.AccFrom,
			Time:    t.Time,
		}
		_, updateErr2 := tx.Model(&data).Insert()
		if updateErr2 != nil {
			tx.Rollback()
			fmt.Println(updateErr2)
			return "Cannot update in transaction table (2)", false
		}
		var temp1 AccountDetail
		err1 := tx.Model(&temp1).Where("account = ?0", t.AccFrom).Select()
		if err1 != nil {
			return "something went wrong", false
		}
		_, updateErr3 := tx.Model(&temp1).Set("balance = ?0,updated_at = ?1 ", temp1.Balance-int(t.Amount), time.Now()).Where("account = ?0", t.AccFrom).Update()
		if updateErr3 != nil {
			tx.Rollback()
			fmt.Println(updateErr3)
			return "Cannot debit balance", false
		}
		var temp2 AccountDetail
		err2 := tx.Model(&temp2).Where("account = ?0", t.AccTo).Select()
		if err2 != nil {
			return "something went wrong", false
		}
		_, updateErr4 := tx.Model(&temp2).Set("balance = ?0,updated_at = ?1 ", temp2.Balance+int(t.Amount), time.Now()).Where("account = ?0", t.AccTo).Update()
		if updateErr4 != nil {
			tx.Rollback()
			fmt.Println(updateErr4)
			return "Cannot Credit balance", false
		}
		tx.Commit()
	} else if t.Type == "debit" {
		tx, txerr := db.Begin()
		if txerr != nil {
			return "Cannot go in begin block", false
		}
		t.Time = time.Now()
		_, updateErr1 := tx.Model(t).Insert()
		if updateErr1 != nil {
			tx.Rollback()
			return "Cannot update in transaction table (1)", false
		}
		data := Transaction{
			Type:    "credit",
			Amount:  t.Amount,
			AccFrom: t.AccTo,
			AccTo:   t.AccFrom,
			Time:    t.Time,
		}
		_, updateErr2 := tx.Model(&data).Insert()
		if updateErr2 != nil {
			tx.Rollback()
			return "Cannot update in transaction table (2)", false
		}
		var temp1 AccountDetail
		err1 := tx.Model(&temp1).Where("account = ?0", t.AccFrom).Select()
		if err1 != nil {
			return "something went wrong", false
		}
		_, updateErr3 := tx.Model(&temp1).Set("balance = ?0,updated_at = ?1 ", temp1.Balance+int(t.Amount), time.Now()).Where("account = ?0", t.AccFrom).Update()
		if updateErr3 != nil {
			tx.Rollback()
			return "Cannot credit balance", false
		}
		var temp2 AccountDetail
		err2 := tx.Model(&temp2).Where("account = ?0", t.AccTo).Select()
		if err2 != nil {
			return "something went wrong", false
		}
		_, updateErr4 := tx.Model(&temp2).Set("balance = ?0,updated_at = ?1 ", temp2.Balance-int(t.Amount), time.Now()).Where("account = ?0", t.AccTo).Update()
		if updateErr4 != nil {
			tx.Rollback()
			return "Cannot debit balance", false
		}
		tx.Commit()
	}
	return "Transaction Done Successfully", true
}

func (t *Transaction) CheckBalance(db *pg.DB, acc int) (string, bool) {

	var out AccountDetail
	err := db.Model(&out).Where("account = ?0", acc).Select()
	if err != nil {
		return "cannot get details ", false
	}
	t.Amount += 500.00
	fmt.Println("the balance is ", out.Balance)
	fmt.Println("balance to  be detucted id ", t.Amount)
	if float64(out.Balance) < t.Amount {
		return "insufficient balance", false
	}
	return "", true

}

func (t *Transaction) CheckActive(db *pg.DB, acc int) bool {

	var out AccountDetail
	db.Model(&out).Column("is_active").Where("account = ?0", acc).Select()
	return out.IsActive

}

func (t *Transaction) CheckIfExist(db *pg.DB) (string, bool) {

	var out []Transaction
	err := db.Model(&out).Where("acc_from = ?0", t.AccFrom).Select()
	if err != nil {
		return "could not fetxh data form database", false
	}
	if len(out) == 0 {
		return "Account number not found", false
	}
	return "Account number found", true

}

func (t *Transaction) GetByAccFrom(db *pg.DB) ([]Transaction, bool, string) {

	var out []Transaction
	err := db.Model(&out).Where("acc_from = ?0", t.AccFrom).Select()
	if err != nil {
		return out, false, "Could not get data from database."
	}
	return out, true, "Successfully fetched data"

}

func (t *Transaction) GetByAcc(db *pg.DB) []AccountDetail {

	var out []AccountDetail
	db.Model(&out).Where("account = ?0", t.AccFrom).WhereOr("account = ?0", t.AccTo).Select()
	return out

}
