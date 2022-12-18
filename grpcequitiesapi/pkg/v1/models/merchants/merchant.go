package merchants

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"dkgosql-merchant-service-v4/internals/adapter/mysql/entities"
	"dkgosql-merchant-service-v4/internals/adapter/mysql/query"
	"dkgosql-merchant-service-v4/internals/consts"
	"dkgosql-merchant-service-v4/internals/util"
	"dkgosql-merchant-service-v4/pkg/v1/models"
	"dkgosql-merchant-service-v4/pkg/v1/models/request"
	"dkgosql-merchant-service-v4/pkg/v1/models/response"

	"github.com/gin-gonic/gin"
)

type MerchantService interface {
	GetMerchantList(c *gin.Context) (models.Response, error)
	CreateMerchant(c *gin.Context) (models.Response, error)
	UpdateMerchantByID(c *gin.Context) (models.Response, error)
}

type merchantService struct {
	db query.MySQLDBStoreAccess
}

func NewMerchantService(db query.MySQLDBStoreAccess) MerchantService {
	return &merchantService{db: db}
}

// UpdateMerchantByID
func (service merchantService) UpdateMerchantByID(c *gin.Context) (models.Response, error) {
	var err error
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var merchantData entities.Merchant
	var resp = models.Response{}

	var updateMerchantRequest request.UpdateMerchantInputRequest
	if err := c.BindJSON(&updateMerchantRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}

	code := strings.Trim(c.Param("code"), "")
	if len(code) == 0 {
		err = errors.New(consts.InvalidCode)
		return resp, err
	}

	var responseMerchant []response.MerchantResponse
	err = service.db.ListMerchantByID(ctx, &responseMerchant, code)
	if err != nil {
		return resp, err
	}
	if len(responseMerchant) == 0 {
		err = errors.New(fmt.Sprintf(consts.ErrorDataNotFoundCode, code))
		return resp, err
	}

	var updateTypeData = make(map[string]interface{})
	updateTypeData["address"] = updateMerchantRequest.Address
	updateTypeData["name"] = updateMerchantRequest.Name

	err = service.db.UpdateMerchantByID(ctx, &merchantData, updateTypeData, code)
	if err != nil {
		return resp, err
	}

	err = service.db.ListMerchantByID(ctx, &responseMerchant, code)
	if err != nil {
		return resp, err
	}
	resp.Data = responseMerchant
	resp.Message = fmt.Sprintf(consts.MerchantUpdatedSuccess, code)
	return resp, nil
}

// CreateMerchant
func (srv merchantService) CreateMerchant(c *gin.Context) (models.Response, error) {
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var resp = models.Response{}
	var addMerchantRequest request.AddMerchantInputRequest
	if err := c.BindJSON(&addMerchantRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}

	currTime := time.Now()
	var status uint8
	status = uint8(consts.ActiveStatus)
	addMerchant := entities.Merchant{
		UpdatedAt: currTime,
		CreatedAt: currTime,
		Code:      addMerchantRequest.Code,
		Status:    &status,
		Name:      addMerchantRequest.Name,
		Address:   addMerchantRequest.Address,
	}

	err := srv.db.CreateMerchant(ctx, &addMerchant)
	if err != nil {
		return resp, err
	}

	var newMerchantMaster []response.MerchantResponse
	newMerchantMaster = append(newMerchantMaster, response.MerchantResponse{
		Name:      addMerchantRequest.Name,
		Code:      addMerchantRequest.Code,
		CreatedAt: currTime.String(),
		Address:   addMerchantRequest.Address,
	})

	resp.Data = newMerchantMaster
	resp.Message = consts.MerchantAddedSuccess
	return resp, nil
}

// GetMerchantList merchants
func (srv *merchantService) GetMerchantList(c *gin.Context) (models.Response, error) {
	var err error
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var merchantData []response.MerchantResponse
	var resp = models.Response{}
	if err != nil {
		return resp, err
	}

	err = srv.db.GetMerchantList(ctx, &merchantData)
	if err != nil {
		return resp, err
	}

	var outMSM []response.MerchantResponse
	for _, row := range merchantData {
		outMSM = append(outMSM, response.MerchantResponse{
			Name:      row.Name,
			Code:      row.Code,
			Address:   row.Address,
			CreatedAt: row.CreatedAt,
		})
	}
	resp.Data = outMSM
	return resp, nil
}
