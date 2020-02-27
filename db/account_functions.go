package db

import (
	//"fmt"
	"log"

	pg "github.com/go-pg/pg"
)

func (pi *AccountDetail) Save(db *pg.DB) error {
	insertErr := db.Insert(pi)
	if insertErr != nil {
		log.Printf("Error while inserting new item into DB, Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("AccountDetail %s inserted successfully.\n", pi.AccountNumber)
	return nil
}
func (pi *AccountDetail) Read(db *pg.DB) AccountDetail {
	var item AccountDetail
	getErr := db.Model(&item).Where("account = ?0", pi.AccountNumber).Select()
	if getErr != nil {
		log.Printf("Error while getting the account details, Reason: %v\n", getErr)
	}
	log.Printf("Read successful for %v\n", item)
	return item
}

func (pi *AccountDetail) Update(db *pg.DB) error {

	tx, txErr := db.Begin()
	if txErr != nil {
		log.Printf("Error while opening tx, Reason: %v\n", txErr)
		return txErr
	}

	_, updateErr := tx.Model(pi).Set("balance = ?balance").Where("account = ?account").Update()
	if updateErr != nil {
		log.Printf("Error while updating item, Reason: %v\n", updateErr)
		tx.Rollback()
		return updateErr
	}
	_, updateErr2 := tx.Model(pi).Set("is_active = ?is_active").Where("account = ?account").Update()
	if updateErr != nil {
		log.Printf("Error while updating item, Reason: %v\n", updateErr)
		tx.Rollback()
		return updateErr2
	}
	tx.Commit()
	log.Printf("Update successful for acc_num: %d\n", pi.AccountNumber)
	return nil
}

func (pi *AccountDetail) Delete(db *pg.DB) error {
	_, deleteErr := db.Model(pi).Where("account = ?account").Delete()
	if deleteErr != nil {
		log.Printf("Error while deleting item, Reason: %v\n", deleteErr)
		return deleteErr
	}
	ni := &Transaction{
		AccFrom: pi.AccountNumber,
		AccTo:   pi.AccountNumber,
	}
	//fmt.Println(ni)
	_, deleteErr2 := db.Model(ni).Where("acc_from = ?0", ni.AccFrom).Delete()
	if deleteErr2 != nil {
		log.Printf("Error while deleting item2, Reason: %v\n", deleteErr2)
		return deleteErr2
	}
	log.Printf("Delete successful for %s,Item\n", pi.AccountNumber)
	return nil
}