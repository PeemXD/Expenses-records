package model

type Expenses struct {
	ID     uint     `gorm:"primaryKey column:id" json:"id"`
	Title  string   `gorm:"column:title" json:"title"`
	Amount float32  `gorm:"column:amount" json:"amount"`
	Note   string   `gorm:"column:note" json:"note"`
	Tags   []string `gorm:"type:text[];column:tags" json:"tags"`
}

type ExpensesForPg struct {
	ID     uint        `gorm:"primaryKey" json:"id"`
	Title  string      `json:"title"`
	Amount float32     `json:"amount"`
	Note   string      `json:"note"`
	Tags   interface{} `gorm:"type:text[]" json:"tags"`
}
