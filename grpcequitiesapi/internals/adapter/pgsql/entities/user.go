// Copyright x go-swagger maintainers
//
// ...
package entities

import (
	"time"
)

var _table_users = "users"

// Users represents the user for this application
// swagger:model
type Users struct {
	ID        int       `gorm:"column:id;primary_key"`
	FkCode    string    `gorm:"column:fk_code;uniqueIndex:Code_Email_UniqueIndex"`
	Email     string    `gorm:"column:email;uniqueIndex:Code_Email_UniqueIndex"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Mobile    string    `gorm:"column:mobile"`
	Password  string    `gorm:"column:password"`
	IsActive  *uint8    `gorm:"column:is_active"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName get sql table name users
func (m *Users) TableName() string {
	return _table_users
}
