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

	"dkgosql-merchant-service-v4/internals/adapter/mysql/entities"
	"dkgosql-merchant-service-v4/internals/adapter/mysql/query"
	"dkgosql-merchant-service-v4/internals/consts"
	"dkgosql-merchant-service-v4/internals/util"
	"dkgosql-merchant-service-v4/pkg/v1/models"
	"dkgosql-merchant-service-v4/pkg/v1/models/request"
	"dkgosql-merchant-service-v4/pkg/v1/models/response"

	gGPCEquities "github.com/dhananjayksharma/dkgosql-grpc-equities/equities"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

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
	defer cancel()
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
			Orderid:  fmt.Sprintf("%d", row.OrderId),
			Userid:   userid,
			Quantity: row.Quantity,
		}); err != nil {
			log.Fatalln("Send", err)
		}
		log.Printf("Sending userid, orderid:%v, %v", userid, row.OrderId)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalln("CloseSend", err)
	}

	for {
		readRow, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Recv", err)
		}

		fmt.Printf("Processed GetUserid:%s, GetOrderid:%s, Status:%t\n", readRow.GetUserid(), readRow.GetOrderid(), readRow.GetStatus())
		//, readRow.Orderprocessedupdatedt.AsTime()
		fmt.Printf("readRow :%v\n\n", readRow)

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
	var responseUser []response.OrdersProcessedResponse
	for _, row := range orderProcessData {
		responseUser = append(responseUser, response.OrdersProcessedResponse{
			Status:    row.Status,
			OrderId:   row.OrderId,
			Quantity:  row.Quantity,
			UserId:    row.UserId,
			CreatedDt: row.CreatedDt,
		})
	}

	resp.Data = responseUser
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
