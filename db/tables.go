package db

import (
	"time"
)

type ttype string

const (
	credit ttype = "credit"
	debit  ttype = "debit"
)

type AccountDetail struct {
	RefPointer    int       `pg:"-"`
	tableName     struct{}  `pg:"account_collection"`
	AccountNumber int       `pg:"account,pk"`
	AccountType   string    `pg:"type"`
	Balance       int       `pg:"balance"`
	Interest      int       `pg:"interest"`
	CreatedAt     time.Time `"pg:created_at"`
	UpdatedAt     time.Time `"pg:updated_at"`
	IsActive      bool      `"pg:is_active"`
}
type BankDetail struct {
	RefPointer    int       `pg:"-"`
	tableName     struct{}  `pg:"bank_collection"`
	Cid           int       `pg:"cid"`
	AccountNumber int       `pg:"account,pk"`
	Ifsc          int       `pg:"ifsc"`
	CreatedAt     time.Time `"pg:created_at"`
	UpdatedAt     time.Time `"pg:updated_at"`
}
type Transaction struct {
	tableName struct{}  `pg:"transactions"`
	TID       int       `pg:"tid,pk"`
	AccFrom   int       `pg:"acc_from"`
	AccTo     int       `pg:"acc_to"`
	Time      time.Time `pg:"time"`
	Amount    float64   `pg:"amount"`
	Type      string    `pg:"type,type=ttype"`
}
type Customerinfo struct {
	RefPointer int       `pg:"-"`
	tableName  struct{}  `pg:"customer"`
	Cid        int       `pg:"id, pk`
	Fname      string    `pg:"fname, unique"`
	Lname      string    `pg:"lname"`
	Address    string    `pg:"address`
	Phoneno    int       `pg:"phoneno"`
	Emailid    string    `pg:"emailid"`
	CreatedAt  time.Time `pg:"created_at"`
	UpdatedAt  time.Time `pg:"updated_at"`
	IsActive   bool      `pg:"is_active"`
}
