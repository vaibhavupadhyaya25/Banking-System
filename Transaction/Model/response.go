package model

import "time"

//TransactionResposnse ..
type TransactionResposnse struct {
	TID     int       `json:"tid"`
	AccFrom int       `json:"acc_from"`
	AccTo   int       `json:"acc_to"`
	Time    time.Time `json:"time"`
	Amount  float64   `json:"amount"`
	Type    string    `json:"type,type=ttype"`
}
