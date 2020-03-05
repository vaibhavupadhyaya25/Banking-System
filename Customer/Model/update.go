package Model

import "time"

type CustomerUpdate struct {
	tableName struct{}  `pg:"customer"`
	Cid       int       `pg:"cid,pk" json:"cid"`
	Fname     string    `json:"fname"`
	Lname     string    `json:"lname"`
	Address   string    `json:"address"`
	Phoneno   string    `json:"phoneno"`
	Emailid   string    `json:"emailid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}
