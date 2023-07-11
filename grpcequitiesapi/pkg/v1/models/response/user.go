package response

type UsersResponse struct {
	Email             string `json:"email"`
	FirstName         string `json:"firstName"`
	FkMerchantCode    string `json:"fkMerchantCode"`
	FkMerchantAppCode string `json:"fkMerchantAppCode"`
	Gender            int    `json:"gender"`
	ID                int    `json:"-"`
	IsActive          *uint8 `json:"isActive"`
	IsVerified        int    `json:"isVerified"`
	LastName          string `json:"lastName"`
	LoginType         int    `json:"loginType"`
	MiddleName        string `json:"middleName"`
	Mobile            string `json:"mobile"`
	PassResetToken    string `json:"passResetToken"`
	PasswordHash      string `json:"-"`
	Source            int    `json:"source"`
	Updateddate       string `json:"updateddate"`
	CAuthOtp          string `json:"-"`
	OtpValidity       string `json:"otpValidity"`
	Createddate       string `json:"createddate"`
}

var _table_users = "users"

type UserLoginResponse struct {
	Email             string `json:"email"`
	FirstName         string `json:"firstName"`
	FkMerchantCode    string `json:"fkMerchantCode"`
	FkMerchantAppCode string `json:"fkMerchantAppCode"`
	Gender            int    `json:"gender"`
	ID                int    `json:"-"`
	IsActive          *uint8 `json:"isActive"`
	IsVerified        int    `json:"isVerified"`
	LastName          string `json:"lastName"`
	LoginType         int    `json:"loginType"`
	MiddleName        string `json:"middleName"`
	Mobile            string `json:"mobile"`
	Password          string `json:"password"`
	Token             string `json:"token"`
	ResetToken        string `json:"restToken"`
}

// TableName get sql table name users
func (m *UserLoginResponse) TableName() string {
	return _table_users
}

// TableName get sql table name users
func (m *UsersResponse) TableName() string {
	return _table_users
}
