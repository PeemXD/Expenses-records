package model

import "github.com/lib/pq"

type Expenses struct {
	ID     uint     `gorm:"primaryKey" json:"id"`
	Title  string   `json:"title"`
	Amount float32  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `gorm:"type:text[]" json:"tags"`
}

type ExpensesForPg struct {
	ID     uint           `gorm:"primaryKey" json:"id"`
	Title  string         `json:"title"`
	Amount float32        `json:"amount"`
	Note   string         `json:"note"`
	Tags   pq.StringArray `gorm:"type:text[]" json:"tags"`
}
