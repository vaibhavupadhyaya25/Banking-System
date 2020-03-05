package Model

import "time"

type CustomerReq struct {
	tableName struct{}  `pg:"customer"`
	Cid       int       `pg:"cid,pk" json:"cid"`
	Fname     string    `json:"fname" validate:"required"`
	Lname     string    `json:"lname" validate:"required"`
	Address   string    `json:"address" validate:"required"`
	Phoneno   string    `json:"phoneno" validate:"required,min=10,max=12"`
	Emailid   string    `json:"emailid" validate:"required,email"`
	CreatedAt time.Time `json:"created_at" validate:"isdefault=time.now()"`
	UpdatedAt time.Time `json:"updated_at" validate:"isdefault=time.now()"`
	IsActive  bool      `json:"is_active" validate:"isdefault=true"`
}
