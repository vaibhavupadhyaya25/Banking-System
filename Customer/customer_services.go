package Customer

import (
	"encoding/json"
	"strconv"

	md "github.com/Banking-System/Customer/Model"
	util "github.com/Banking-System/Customer/Util"
	db "github.com/Banking-System/Db"
	"github.com/rightjoin/fuel"
	"gopkg.in/go-playground/validator.v9"
)

type Myservices struct {
	fuel.Service
	addCustomer    fuel.POST
	getCustomer    fuel.GET `route:"{cid}"`
	updateCustomer fuel.PUT
	deleteCustomer fuel.DELETE `route:"delete/{cid}"`
}

func (c *Myservices) DeleteCustomer(cid int) interface{} {
	pg_db := db.Connect()
	util.Del(pg_db, cid)
	return nil
}

func (c *Myservices) UpdateCustomer(x fuel.Aide) interface{} {
	res := x.Post()
	var b interface{}

	dat, _ := json.Marshal(res)
	var result md.CustomerUpdate
	json.Unmarshal(dat, &result)

	validate := validator.New()
	err := validate.Struct(result)
	if err != nil {
		return err.Error()
	}

	b = util.CustomerValidate(res)
	bs, err := json.Marshal(b)
	if err != nil {
		return err
	}
	pg_db := db.Connect()
	var data db.Customerinfo
	var dn md.CustomerRes

	json.Unmarshal(bs, &data)
	data.Cid, _ = strconv.Atoi(res["Cid"])

	if res["Fname"] != "" {
		util.Updfname(pg_db, data)
		dn.Fname = data.Fname
	}

	if res["Lname"] != "" {
		util.Updlname(pg_db, data)
		dn.Lname = data.Lname
	}

	if res["Address"] != "" {
		util.Updaddress(pg_db, data)
		dn.Address = data.Address
	}

	if res["Phoneno"] != "" {
		util.Updphone(pg_db, data)
		dn.Phoneno = data.Phoneno
	}

	if res["Emailid"] != "" {
		util.Updemail(pg_db, data)
		dn.Emailid = data.Emailid
	}
	return dn
}

func (c *Myservices) AddCustomer(x fuel.Aide) interface{} {
	res := x.Post()
	dat, _ := json.Marshal(res)
	var result md.CustomerReq
	json.Unmarshal(dat, &result)

	validate := validator.New()
	err1 := validate.Struct(result)
	if err1 != nil {
		return err1.Error()
	}

	var d interface{}
	d = util.CustomerValidate(res)
	bs, err := json.Marshal(d)
	if err != nil {
		return err
	}
	pg_db := db.Connect()
	var data db.Customerinfo
	json.Unmarshal(bs, &data)
	util.Addcus(pg_db, data)

	var dn md.CustomerRes
	dn.Fname = data.Fname
	dn.Lname = data.Lname
	dn.Address = data.Address
	dn.Phoneno = data.Phoneno
	dn.Emailid = data.Emailid

	return dn
}

func (c *Myservices) GetCustomer(cid int) interface{} {
	pg_db := db.Connect()
	var customer db.Customerinfo
	customer = util.Read(pg_db, cid)
	bs, err := json.Marshal(customer)
	if err != nil {
		return err
	}
	return string(bs)
}
