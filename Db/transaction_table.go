package db

import "time"

type ttype string

const (
	credit ttype = "credit"
	debit  ttype = "debit"
)

//Transaction ..
type Transaction struct {
	tableName struct{}  `pg:"transactions`
	TID       int       `pg:"tid,pk" json:"tid"`
	AccFrom   int       `pg:"acc_from" json:"acc_from" validate:"required"`
	AccTo     int       `pg:"acc_to" json:"acc_to" validate:"required"`
	Time      time.Time `pg:"time" json:"time" validate:"isdefault=time.now()"`
	Amount    float64   `pg:"amount" json:"amount" validate:"required"`
	Type      string    `pg:"type,type=ttype" json:"type,type=ttype" validate:"required"`
}
