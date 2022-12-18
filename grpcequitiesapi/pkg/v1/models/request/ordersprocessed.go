package request

type AddOrderProcessedInputRequest struct {
	Email     string `json:"email" binding:"required,min=7,max=251,email"`
	FirstName string `json:"firstName" binding:"required,min=1,max=25,alphanum"`
	LastName  string `json:"lastName" binding:"required,min=1,max=25,alphanum"`
	Password  string `json:"password" binding:"required,min=7,max=251"`
	Mobile    string `json:"mobile" binding:"omitempty,min=10,max=10,alphanum"`
	Code      string `json:"code"`
}

type ListOrderProcessedInputRequest struct {
	Email string `json:"email" binding:"required,min=7,max=251,email"`
	Code  string `json:"code" binding:"required,min=16,max=24,alphanum"`
}

type UpdateOrderProcessedInputRequest struct {
	UserID  string `json:"userID" binding:"required,min=1"`
	OrderID string `json:"orderID" binding:"required,min=1"`
}

type QueryOrderProcessedInputRequest struct {
	Limit  int
	Skip   int
	UserID string
}

type LoginOrderProcessedInputRequest struct {
	Email    string `json:"email" binding:"required,min=7,max=251,email"`
	Password string `json:"password" binding:"required,min=7,max=251"`
	Code     string `json:"code" binding:"required,min=16,max=24,alphanum"`
}
