package handlers

import (
	"dkgosql-merchant-service-v4/internals/util"
	"dkgosql-merchant-service-v4/pkg/v1/models/orderprocessed"

	"github.com/gin-gonic/gin"
)

type OrderProcessedHandler interface {
	GetOrderProcessedList(c *gin.Context)
	CreateOrderProcessed(c *gin.Context)

	ListOrderProcessedByID(c *gin.Context)
	UpdateOrderProcessedByID(c *gin.Context)
	BulkOrderProcessedByUserId(c *gin.Context)
}

type orderProcessedHandler struct {
	service orderprocessed.OrderProcessedService
}

func NewOrderProcessedHandler(service orderprocessed.OrderProcessedService) OrderProcessedHandler {
	return &orderProcessedHandler{service: service}
}

func (srv *orderProcessedHandler) BulkOrderProcessedByUserId(c *gin.Context) {

	resp, err := srv.service.BulkOrderProcessedByUserId(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

func (srv *orderProcessedHandler) GetOrderProcessedList(c *gin.Context) {
	resp, err := srv.service.GetOrderProcessedList(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

func (srv *orderProcessedHandler) CreateOrderProcessed(c *gin.Context) {
	resp, err := srv.service.CreateOrderProcessed(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}
func (srv *orderProcessedHandler) ListOrderProcessedByID(c *gin.Context) {
	resp, err := srv.service.ListOrderProcessedByID(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}

func (srv *orderProcessedHandler) UpdateOrderProcessedByID(c *gin.Context) {
	resp, err := srv.service.UpdateOrderProcessedByID(c)
	if err != nil {
		util.HandleError(c, err)
		return
	}
	util.JSON(c, resp.Data, resp.Message)
}
