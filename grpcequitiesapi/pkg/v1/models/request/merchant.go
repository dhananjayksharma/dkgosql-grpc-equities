package request

type AddMerchantInputRequest struct {
	Name    string `json:"name" binding:"required,min=5,max=55"`
	Code    string `json:"code" binding:"required,min=16,max=24,alphanum"`
	Address string `json:"address" binding:"required,min=5,max=465"`
}

type UpdateMerchantInputRequest struct {
	Name    string `json:"name" binding:"required,min=5,max=55"`
	Address string `json:"address" binding:"required,min=5,max=465"`
}
