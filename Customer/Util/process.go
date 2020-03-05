package Util

import (
	db "github.com/Banking-System/Db"
	pg "github.com/go-pg/pg"
)

func Del(dbRef *pg.DB, cid int) {
	new1 := &db.Customerinfo{
		Cid: cid,
	}
	// var ab db.AccountDetail
	// dbRef.Model(&ab).Column("account").Where("cid = ?0", new1.Cid).Select()

	new1.DeleteRecord(dbRef)
	// Delete(dbRef, ab.AccountNumber)
}

func Updfname(dbRef *pg.DB, res db.Customerinfo) {
	res.UpdateCustomer(dbRef)
}

func Updlname(dbRef *pg.DB, res db.Customerinfo) {
	res.Updatelname(dbRef)
}

func Updaddress(dbRef *pg.DB, res db.Customerinfo) {
	res.Updateaddress(dbRef)
}

func Updphone(dbRef *pg.DB, res db.Customerinfo) {
	res.Updatephone(dbRef)
}

func Updemail(dbRef *pg.DB, res db.Customerinfo) {
	res.Updateemail(dbRef)
}

func Addcus(dbRef *pg.DB, res db.Customerinfo) {
	res.Save(dbRef)
}

func Read(dbRef *pg.DB, cid int) db.Customerinfo {
	new1 := &db.Customerinfo{
		Cid: cid,
	}
	item := new1.GetById(dbRef)
	return item
}

func GetById(dbRef *pg.DB, cid int) db.Customerinfo {
	newPI := &db.Customerinfo{
		Cid: cid,
	}
	item := newPI.GetById(dbRef)
	return item
}
