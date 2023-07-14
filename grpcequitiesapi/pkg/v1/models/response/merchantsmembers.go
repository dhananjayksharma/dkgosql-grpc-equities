package response

type MerchantsMembersResponse struct {
	ID           int    `json:"-"`
	MerchantName string `json:"merchantName"`
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	FkCode       string `json:"code"`
	IsActive     *uint8 `json:"isActive"`
	Mobile       string `json:"mobile"`
	UpdatedAt    string `json:"-"`
	CreatedAt    string `json:"createdAt"`
}

var _table_mmr = "users"

// TableName get sql table name users
func (m *MerchantsMembersResponse) TableName() string {
	return _table_mmr
}
