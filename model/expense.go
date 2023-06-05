package model

import "github.com/lib/pq"

//! ,omitempty in json, if not have field or fill null and if struct type string, it will be set to ""
type Expenses struct {
	ID     uint           `gorm:"primaryKey column:id" json:"id"`
	Title  string         `gorm:"column:title" json:"title,omitempty"`
	Amount float32        `gorm:"column:amount" json:"amount,omitempty"`
	Note   string         `gorm:"column:note" json:"note"`
	Tags   pq.StringArray `gorm:"type:text[];column:tags" json:"tags"`
}
