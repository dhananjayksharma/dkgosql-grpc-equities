package orderprocessed

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"grpcequitiesapi/internals/adapter/pgsql/entities"
	"grpcequitiesapi/internals/adapter/pgsql/query"
	"grpcequitiesapi/internals/consts"
	"grpcequitiesapi/internals/util"
	"grpcequitiesapi/pkg/v1/models"
	"grpcequitiesapi/pkg/v1/models/request"
	"grpcequitiesapi/pkg/v1/models/response"

	gGPCEquities "github.com/dhananjayksharma/dkgosql-grpc-equities/equities"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var MAX_ALLOWED_QTY int64 = 1600

type OrderProcessedService interface {
	GetOrderProcessedList(c *gin.Context) (models.Response, error)
	CreateOrderProcessed(c *gin.Context) (models.Response, error)
	ListOrderProcessedByID(c *gin.Context) (models.Response, error)
	UpdateOrderProcessedByID(c *gin.Context) (models.Response, error)

	BulkOrderProcessedByUserId(c *gin.Context) (models.Response, error)
}

type orderProcessedService struct {
	db       query.MySQLDBStoreAccess
	conngRPC *grpc.ClientConn
}

func NewOrderProcessedService(db query.MySQLDBStoreAccess, conngRPC *grpc.ClientConn) OrderProcessedService {
	return &orderProcessedService{db: db, conngRPC: conngRPC}
}

func (service orderProcessedService) BulkOrderProcessedByUserId(c *gin.Context) (models.Response, error) {
	// set context

	var resp = models.Response{}

	clientgRCP := gGPCEquities.NewOrderClient(service.conngRPC)

	// Define the context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// var deadlineMs = flag.Int("deadline_ms", 20*1000, "Default deadline in milliseconds.")

	// clientDeadline := time.Now().Add(time.Duration(*deadlineMs) * time.Millisecond)
	// ctx, cancel = context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	if ctx.Err() == context.Canceled {
		return resp, errors.New("client cancelled, abandoning.")
	}

	stream, err := clientgRCP.ProcessOrder(ctx)

	if err != nil {
		log.Fatalln("Opening stream", err)
	}

	userid := strings.Trim(c.Param("userid"), "")
	if len(userid) == 0 {
		err := errors.New(consts.InvalidUserId)
		return resp, err
	}

	var orderProcessData []response.OrdersProcessedResponse
	err = service.db.ListOrderProcessedByID(ctx, &orderProcessData, userid)
	if err != nil {
		return resp, err
	}

	for _, row := range orderProcessData {
		if err := stream.Send(&gGPCEquities.OrderRequest{
			Orderid:    fmt.Sprintf("%d", row.OrderID),
			Userid:     userid,
			Allowedqty: MAX_ALLOWED_QTY,
			Quantity:   row.Quantity,
		}); err != nil {
			log.Fatalln("Send", err)
		}
		log.Printf("Sending userid, orderid:%v, %v", userid, row.OrderID)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalln("CloseSend", err)
	}

	for {
		readRow, err := stream.Recv()
		if grpcError := readRow.GetError(); grpcError != nil {
			log.Printf("error found for processing order err:%v, errCode:%v, err Details: %v", err, grpcError.Code, grpcError.GetDetails())
			continue
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Recv", err)
		}
		if resp := readRow.GetOrderresponse(); resp != nil {
			fmt.Printf("Processed GetUserid:%s, GetOrderid:%s, Org Qty:%d, Processed Qty:%d, NOT Processed Qty:%d, Status:%t, OrderProcessed Dt:%s\n", resp.GetUserid(), resp.GetOrderid(), resp.GetQuantity(), resp.GetProcessedQuantity(), resp.GetNotProcessedQuantity(), resp.GetStatus(), resp.Orderprocessedupdatedt.String())
			fmt.Printf("readRow :%v\n\n", resp)
		}

	}

	orderProcessLen := len(orderProcessData)
	resp.Data = fmt.Sprintf("Number of data processed: %d", orderProcessLen)
	return resp, nil
}

func (service orderProcessedService) GetOrderProcessedList(c *gin.Context) (models.Response, error) {
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var resp = models.Response{}
	_ = ctx

	return resp, nil
}

func (service orderProcessedService) CreateOrderProcessed(c *gin.Context) (models.Response, error) {
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var resp = models.Response{}
	_ = ctx

	return resp, nil
}
func (service orderProcessedService) ListOrderProcessedByID(c *gin.Context) (models.Response, error) {
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var orderProcessData []response.OrdersProcessedResponse
	var resp = models.Response{}

	userid := strings.Trim(c.Param("userid"), "")
	if len(userid) == 0 {
		err := errors.New(consts.InvalidUserId)
		return resp, err
	}
	skip_number, err := strconv.ParseUint(c.Query("skip"), 10, 64)
	if skip_number < 0 || err != nil {
		if err != nil {
			return resp, err
		}
		err = errors.New(consts.SkipMessage)
		return resp, err
	}

	page_limit, _ := strconv.ParseUint(c.Query("limit"), 10, 64)

	if page_limit < 1 {
		err = errors.New(consts.PageLimitMessage)
		return resp, err
	}

	var queryParams = request.QueryOrderProcessedInputRequest{UserID: userid, Limit: int(page_limit), Skip: int(skip_number)}
	_ = queryParams

	err = service.db.ListOrderProcessedByID(ctx, &orderProcessData, userid)
	if err != nil {
		return resp, err
	}
	var responseOrder []response.OrdersProcessedResponse
	for _, row := range orderProcessData {
		responseOrder = append(responseOrder, response.OrdersProcessedResponse{
			Status:    row.Status,
			OrderID:   row.OrderID,
			Quantity:  row.Quantity,
			UserId:    row.UserId,
			CreatedDt: row.CreatedDt,
		})
	}

	resp.Data = responseOrder
	return resp, nil
}

func (service orderProcessedService) UpdateOrderProcessedByID(c *gin.Context) (models.Response, error) {
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var resp = models.Response{}
	var orderProcessData entities.OrdersProcessed

	var updateOrderProcessRequest request.UpdateOrderProcessedInputRequest
	if err := c.BindJSON(&updateOrderProcessRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}

	userid := updateOrderProcessRequest.UserID
	if len(userid) == 0 {
		err := errors.New(consts.InvalidUserId)
		return resp, err
	}

	orderid := updateOrderProcessRequest.UserID
	if len(orderid) == 0 {
		err := errors.New(consts.InvalidUserId)
		return resp, err
	}

	var responseOrderProcess []response.OrdersProcessedResponse
	err := service.db.ListOrderProcessedByID(ctx, &responseOrderProcess, userid)
	if err != nil {
		return resp, err
	}
	if len(responseOrderProcess) == 0 {
		err = errors.New(fmt.Sprintf(consts.ErrorOrderDataNotFoundCode, userid))
		return resp, err
	}

	var updateTypeData = make(map[string]interface{})
	// updateTypeData["order_id"] = updateOrderProcessRequest.OrderId
	// updateTypeData["user_id"] = updateOrderProcessRequest.UserID
	updateTypeData["status"] = consts.OrdrActiveStatus
	updateTypeData["updated_dt"] = time.Now()
	fmt.Printf("updateOrderProcessRequest: %#v", updateOrderProcessRequest)
	err = service.db.UpdateOrderProcessedByID(ctx, &orderProcessData, updateTypeData, updateOrderProcessRequest)
	if err != nil {
		return resp, err
	}

	err = service.db.ListOrderProcessedByID(ctx, &responseOrderProcess, userid)
	if err != nil {
		return resp, err
	}
	resp.Data = responseOrderProcess
	resp.Message = fmt.Sprintf(consts.OrderProcessedUpdatedSuccess, userid)
	return resp, nil
}
