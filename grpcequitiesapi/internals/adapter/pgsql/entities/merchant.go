package entities

import (
	"time"
)

var _table_mc = "merchants"

type Merchant struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;not null"`
	Code      string    `gorm:"column:code;index:Code_UniqueIndex" json:"code"`
	Name      string    `gorm:"column:name" json:"name"`
	Address   string    `gorm:"column:address" json:"address"`
	Status    *uint8    `gorm:"column:status;default:1" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

// TableName get sql table name merchants
func (m *Merchant) TableName() string {
	return _table_mc
}
