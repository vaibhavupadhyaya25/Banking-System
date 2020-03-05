package db

import "time"

//AccountDetail ...
type AccountDetail struct {
	RefPointer    int       `pg:"-",`
	tableName     struct{}  `pg:"account_collection"`
	Cid           int       `pg:"cid" json:"Cid"`
	AccountNumber int       `pg:"account,pk" json:"AccountNumber"`
	AccountType   string    `pg:"type" json:"AccountType"`
	Balance       float64   `pg:"balance" json:"Balance"`
	Interest      float64   `pg:"interest" json:"Interest"`
	CreatedAt     time.Time `pg:"created_at" json:"CreatedAt" validate:"isdefault=time.now()"`
	UpdatedAt     time.Time `pg:"updated_at" json:"UpdatedAt" validate:"isdefault=time.now()"`
	IsActive      bool      `pg:"is_active" json:"IsActive" validata:"isdefault=true"`
}
