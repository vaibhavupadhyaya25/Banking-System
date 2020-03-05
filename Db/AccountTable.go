package db

import "time"

//AccountDetail ..
type AccountDetail struct {
	RefPointer    int       `pg:"-"`
	tableName     struct{}  `pg:"account_collection"`
	Cid           int       `pg:"cid"`
	AccountNumber int       `pg:"account,pk"`
	AccountType   string    `pg:"type"`
	Balance       int       `pg:"balance"`
	Interest      int       `pg:"interest"`
	CreatedAt     time.Time `pg:"created_at"`
	UpdatedAt     time.Time `pg:"updated_at"`
	IsActive      bool      `pg:"is_active"`
}
