package db

import "time"

// Customerinfo ..
type Customerinfo struct {
	RefPointer int       `json:"-"`
	tableName  struct{}  `json:"customer"`
	Cid        int       `json:"id, pk, serial`
	Fname      string    `json:"fname, unique"`
	Lname      string    `json:"lname"`
	Address    string    `json:"address`
	Phoneno    int       `json:"phoneno"`
	Emailid    string    `json:"emailid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsActive   bool      `json:"is_active"`
}
