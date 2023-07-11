package request

type AddUserInputRequest struct {
	Email     string `json:"email" binding:"required,min=7,max=251,email"`
	FirstName string `json:"firstName" binding:"required,min=1,max=25,alphanum"`
	LastName  string `json:"lastName" binding:"required,min=1,max=25,alphanum"`
	Password  string `json:"password" binding:"required,min=7,max=251"`
	Mobile    string `json:"mobile" binding:"omitempty,min=10,max=10,alphanum"`
	Code      string `json:"code"`
}

type ListUserInputRequest struct {
	Email string `json:"email" binding:"required,min=7,max=251,email"`
	Code  string `json:"code" binding:"required,min=16,max=24,alphanum"`
}

type UpdateUserInputRequest struct {
	FirstName string `json:"firstName" binding:"required,min=1,max=25,alphanum"`
	LastName  string `json:"lastName" binding:"required,min=1,max=25,alphanum"`
	Mobile    string `json:"mobile" binding:"omitempty,min=10,max=10,alphanum"`
}

type QueryMembersInputRequest struct {
	Limit int
	Skip  int
	Code  string
}

type LoginUserInputRequest struct {
	Email    string `json:"email" binding:"required,min=7,max=251,email"`
	Password string `json:"password" binding:"required,min=7,max=251"`
	Code     string `json:"code" binding:"required,min=16,max=24,alphanum"`
}
