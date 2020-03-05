package Util

import (
	"time"

	db "github.com/code/Banking-System/Db"
)

func CustomerValidate(res map[string]string) interface{} {
	var data db.Customerinfo
	// val, err := strconv.Atoi(res["Cid"])
	// if err != nil {
	// 	return "invalid customer id"
	// }

	// data.Cid = val
	data.Fname = res["Fname"]
	data.Lname = res["Lname"]
	data.Address = res["Address"]
	data.Emailid = res["Emailid"]

	// val2, err2 := strconv.Atoi(res["Phoneno"])
	// if err2 != nil {
	// 	return "invalid phone no entered"
	// }
	// if len(res["phoneno"]) != 10 {
	// 	return "invalid phone no entered"
	// }
	data.Phoneno = res["Phoneno"]

	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	data.IsActive = true

	return data
}
