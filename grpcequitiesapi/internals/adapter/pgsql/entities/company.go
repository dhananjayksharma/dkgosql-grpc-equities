// Copyright x go-swagger maintainers
//
// ...
package entities

import (
	"time"
)

var _table_company = "company"

// Users represents the company for this application
// swagger:model
type Company struct {
	ID        int       `gorm:"column:id;primary_key"`
	Code      string    `gorm:"column:code;uniqueIndex:Code_UniqueIndex"`
	Name      string    `gorm:"column:name"`
	IsActive  *uint8    `gorm:"column:is_active"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName get sql table name companies
func (m *Company) TableName() string {
	return _table_company
}
