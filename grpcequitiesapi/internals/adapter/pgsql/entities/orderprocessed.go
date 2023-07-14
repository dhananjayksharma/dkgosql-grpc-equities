package entities

import (
	"time"
)

var _table_ordersprocessed = "ordersprocessed"

// OrdersProcessed represents the processed order for users
// swagger:model
type OrdersProcessed struct {
	ID                int       `gorm:"column:id;primary_key"`
	Status            uint8     `gorm:"column:status"`
	UserId            int       `gorm:"column:user_id"`
	CompanyID         int       `gorm:"column:company_id"`
	OrderID           int       `gorm:"column:order_id"`
	OrderType         int32     `gorm:"column:order_type"`
	Quantity          int64     `gorm:"column:quantity"`
	QuantityProcessed int64     `gorm:"column:quantity_processed"`
	CreatedDt         time.Time `gorm:"column:created_dt"`
	UpdatedDt         time.Time `gorm:"column:updated_dt"`
}

// TableName get sql table name ordersprocessed
func (m *OrdersProcessed) TableName() string {
	return _table_ordersprocessed
}
