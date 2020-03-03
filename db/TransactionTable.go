package db

import "time"

type ttype string

const (
	credit ttype = "credit"
	debit  ttype = "debit"
)

type Transaction struct {
	tableName struct{}  `json:"transactions"`
	TID       int       `json:"tid,pk"`
	AccFrom   int       `json:"acc_from"`
	AccTo     int       `json:"acc_to"`
	Time      time.Time `json:"time"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type,type=ttype"`
}
