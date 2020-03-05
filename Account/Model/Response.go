package model

//AccountResponse ...
type AccountResponse struct {
	Cid           int  `json:"Cid"`
	AccountNumber int  `json:"AccountNumber"`
	IsActive      bool `json:"IsActive"`
}

//UpdateResponse ...
type UpdateResponse struct {
	AccountNumber int     `json:"AccountNumber" `
	Balance       float64 `json:"Balance" `
	IsActive      bool    `json:"IsActive" `
}
