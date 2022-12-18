package response

type OrdersProcessedResponse struct {
	ID        int    `json:"-"`
	Status    uint8  `json:"status"`
	UserId    int    `json:"user_id"`
	OrderId   int    `json:"order_id"`
	Quantity  int64  `json:"quantity"`
	UpdatedDt string `json:"-"`
	CreatedDt string `json:"created_dt"`
}

var _table_opr = "ordersprocessed"

// TableName get sql table name ordersprocessed
func (m *OrdersProcessedResponse) TableName() string {
	return _table_opr
}
