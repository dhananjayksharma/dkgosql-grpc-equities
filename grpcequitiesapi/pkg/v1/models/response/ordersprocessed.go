package response

type OrdersProcessedResponse struct {
	ID        int    `json:"-"`
	Status    uint8  `json:"status"`
	UserId    int    `json:"user_id"`
	CompanyID int    `json:"company_id"`
	OrderID   int    `json:"order_id"`
	OrderType int32  `json:"order_type"`
	Quantity  int64  `json:"quantity"`
	UpdatedDt string `json:"-"`
	CreatedDt string `json:"created_dt"`
}

var _table_opr = "ordersprocessed"

// TableName get sql table name ordersprocessed
func (m *OrdersProcessedResponse) TableName() string {
	return _table_opr
}
