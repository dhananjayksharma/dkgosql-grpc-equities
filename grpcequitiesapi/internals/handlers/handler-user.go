package handlers

import (
	"dkgosql-merchant-service-v4/internals/util"
	"dkgosql-merchant-service-v4/pkg/v1/models/users"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	ListMembersByCode(c *gin.Context)
	CreateMerchantMember(c *gin.Context)

	LoginMember(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type userHandler struct {
	service users.UserService
}

func NewUserHandler(service users.UserService) UserHandler {
	return &userHandler{service: service}
}

func (srv *userHandler) RefreshToken(c *gin.Context) {
	resp, err := srv.service.RefreshToken(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

func (srv *userHandler) LoginMember(c *gin.Context) {
	resp, err := srv.service.LoginMember(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}
func (srv *userHandler) ListMembersByCode(c *gin.Context) {
	resp, err := srv.service.ListMembersByCode(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

func (srv *userHandler) CreateMerchantMember(c *gin.Context) {
	resp, err := srv.service.CreateMerchantMember(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}
