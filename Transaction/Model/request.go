package model

type ttype string

const (
	credit ttype = "credit"
	debit  ttype = "debit"
)

//TransactionRequest ..
type TransactionRequest struct {
	AccFrom string `json:"acc_from" validate:"required"`
	AccTo   string `json:"acc_to" validate:"required"`
	Amount  string `json:"amount" validate:"required"`
	Type    string `json:"type,type=ttype" validate:"required"`
}
