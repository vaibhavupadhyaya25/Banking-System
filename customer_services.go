package main

import (
	"encoding/json"

	db "./db"
	"github.com/rightjoin/fuel"
)

type myservices struct {
	fuel.Service
	addCustomer    fuel.POST
	getCustomer    fuel.GET `route:"{cid}"`
	updateCustomer fuel.PUT
	deleteCustomer fuel.DELETE `route:"delete/{cid}"`
}

func (c *myservices) DeleteCustomer(cid int) interface{} {
	pg_db := db.Connect()
	del(pg_db, cid)
	return nil
}

func (c *myservices) UpdateCustomer(x fuel.Aide) interface{} {
	res := x.Post()
	var b interface{}
	b = CustomerValidate(res)
	bs, err := json.Marshal(b)
	if err != nil {
		return err
	}
	pg_db := db.Connect()
	var data db.Customerinfo
	json.Unmarshal(bs, &data)
	if res["Fname"] != "" {
		updfname(pg_db, data)
	}

	if res["Lname"] != "" {
		updlname(pg_db, data)
	}

	if res["Address"] != "" {
		updaddress(pg_db, data)
	}

	if res["Phoneno"] != "" {
		updphone(pg_db, data)
	}

	if res["Emailid"] != "" {
		updemail(pg_db, data)
	}
	return data
}

func (c *myservices) AddCustomer(x fuel.Aide) interface{} {
	res := x.Post()
	var d interface{}
	d = CustomerValidate(res)
	bs, err := json.Marshal(d)
	if err != nil {
		return err
	}
	pg_db := db.Connect()
	var data db.Customerinfo
	json.Unmarshal(bs, &data)
	addcus(pg_db, data)
	return data
}

func (c *myservices) GetCustomer(cid int) interface{} {
	pg_db := db.Connect()
	var customer db.Customerinfo
	customer = Read(pg_db, cid)
	bs, err := json.Marshal(customer)
	if err != nil {
		return err
	}
	return string(bs)
}
