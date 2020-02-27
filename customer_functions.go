package main

import (
	"strconv"
	"time"

	db "./db"
	pg "github.com/go-pg/pg"
)

func del(dbRef *pg.DB, cid int) {
	new1 := &db.Customerinfo{
		Cid: cid,
	}
	new1.DeleteRecord(dbRef)
}

func updfname(dbRef *pg.DB, res db.Customerinfo) {
	res.UpdateCustomer(dbRef)
}

func updlname(dbRef *pg.DB, res db.Customerinfo) {
	res.Updatelname(dbRef)
}

func updaddress(dbRef *pg.DB, res db.Customerinfo) {
	res.Updateaddress(dbRef)
}

func updphone(dbRef *pg.DB, res db.Customerinfo) {
	res.Updatephone(dbRef)
}

func updemail(dbRef *pg.DB, res db.Customerinfo) {
	res.Updateemail(dbRef)
}

func CustomerValidate(res map[string]string) interface{} {
	var data db.Customerinfo
	val, err := strconv.Atoi(res["Cid"])
	if err != nil {
		return "invalid customer id"
	}

	data.Cid = val
	data.Fname = res["Fname"]
	data.Lname = res["Lname"]
	data.Address = res["Address"]
	data.Emailid = res["Emailid"]

	val2, err2 := strconv.Atoi(res["Phoneno"])
	if err2 != nil {
		return "invalid phone no entered"
	}
	// if len(res["phoneno"]) != 10 {
	// 	return "invalid phone no entered"
	// }
	data.Phoneno = val2

	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	data.IsActive = true

	return data
}

func addcus(dbRef *pg.DB, res db.Customerinfo) {
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
