package handlers

import (
	"grpcequitiesapi/internals/util"
	"grpcequitiesapi/pkg/v1/models/companies"

	"github.com/gin-gonic/gin"
)

// CompanyHandler
type CompanyHandler interface {
	GetCompanyList(c *gin.Context)
}

// companyHandler
type companyHandler struct {
	service companies.CompanyService
}

// NewCompanyHandler
func NewCompanyHandler(service companies.CompanyService) CompanyHandler {
	return &companyHandler{service: service}
}

// GetCompanyList
func (srv *companyHandler) GetCompanyList(c *gin.Context) {

	// err := middleware.Claim(c)
	// if err != nil {
	// 	util.HandleError(c, err)
	// 	return
	// }

	resp, err := srv.service.GetCompanyList(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

// CreateCompany
/*func (srv *companyHandler) CreateCompany(c *gin.Context) {
	resp, err := srv.service.CreateCompany(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}*/

// UpdateCompanyByID
/*func (srv *companyHandler) UpdateCompanyByID(c *gin.Context) {
	resp, err := srv.service.UpdateCompanyByID(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}*/
