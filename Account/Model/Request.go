package model

//AccountRequest ...
type AccountRequest struct {
	Cid           string `json:"Cid" validate:"required"`
	AccountNumber string `json:"AccountNumber" validate:"required"`
	AccountType   string `json:"AccountType" validate:"required"`
	Balance       string `json:"Balance" validate:"required"`
	Interest      string `json:"Interest" validate:"required"`
	IsActive      string `json:"IsActive" `
}

//UpdateRequest ...
type UpdateRequest struct {
	AccountNumber string `json:"AccountNumber" validate:"required"`
	Balance       string `json:"Balance" validate:"required"`
	IsActive      string `json:"IsActive" validate:"required"`
}
