package db

import "time"

type AccountDetail struct {
	RefPointer    int       `json:"-"`
	tableName     struct{}  `json:"account_collection"`
	Cid           int       `json:"cid"`
	AccountNumber int       `json:"account,pk"`
	AccountType   string    `json:"type"`
	Balance       int       `json:"balance"`
	Interest      int       `json:"interest"`
	CreatedAt     time.Time `"json:created_at"`
	UpdatedAt     time.Time `"json:updated_at"`
	IsActive      bool      `"json:is_active"`
}
