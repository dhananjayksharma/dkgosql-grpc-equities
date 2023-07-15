package request

type AddCompanyInputRequest struct {
	Name string `json:"name" binding:"required,min=5,max=55"`
	Code string `json:"code" binding:"required,min=16,max=24,alphanum"`
}

type UpdateCompanyInputRequest struct {
	Name string `json:"name" binding:"required,min=5,max=55"`
}
