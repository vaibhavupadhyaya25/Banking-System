package Db

import "time"

//Customerinfo ...
type Customerinfo struct {
	tableName struct{}  `pg:"customer"`
	Cid       int       `pg:"cid,pk" json:"cid"`
	Fname     string    `pg:"fname" json:"fname" validate:"required"`
	Lname     string    `pg:"lname" json:"lname" validate:"required"`
	Address   string    `pg:"address" json:"address" validate:"required"`
	Phoneno   string    `pg:"phoneno" json:"phoneno" validate:"required,min=10,max=12"`
	Emailid   string    `pg:"emailid" json:"emailid" validate:"required,email"`
	CreatedAt time.Time `pg:"created_at" json:"created_at" validate:"isdefault=time.now()"`
	UpdatedAt time.Time `pg:"updated_at" json:"updated_at" validate:"isdefault=time.now()"`
	IsActive  bool      `pg:"is_active" json:"is_active" validate:"isdefault=true"`
}
