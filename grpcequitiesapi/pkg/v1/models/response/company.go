package response

type CompanyResponse struct {
	CreatedAt string `json:"created_at"`
	Name      string `json:"name"`
	Code      string `json:"code"`
}

var _table_css = "company"

// TableName get sql table name merchants
func (m *CompanyResponse) TableName() string {
	return _table_css
}
