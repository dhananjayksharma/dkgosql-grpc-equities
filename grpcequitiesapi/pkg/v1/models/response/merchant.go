package response

type MerchantResponse struct {
	CreatedAt string `json:"created_at"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Code      string `json:"code"`
}

var _table_cs = "merchants"

// TableName get sql table name merchants
func (m *MerchantResponse) TableName() string {
	return _table_cs
}
