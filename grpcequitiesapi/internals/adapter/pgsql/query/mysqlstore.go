package query

import (
	"context"
	"fmt"
	"grpcequitiesapi/internals/adapter/pgsql/entities"
	"grpcequitiesapi/internals/consts"
	"grpcequitiesapi/internals/util"
	"grpcequitiesapi/pkg/v1/models/request"
	"grpcequitiesapi/pkg/v1/models/response"
	"log"
	"strings"

	"gorm.io/gorm"
)

type MySQLDBStore struct {
	db *gorm.DB
}

func NewMySQLDBStore(db *gorm.DB) *MySQLDBStore {
	return &MySQLDBStore{db: db}
}

type MySQLDBStoreAccess interface {
	GetMerchantList(ctx context.Context, merchantData *[]response.MerchantResponse) error
	CreateMerchant(ctx context.Context, merchantData *entities.Merchant) error
	ListMerchantByID(ctx context.Context, merchantData *[]response.MerchantResponse, code string) error
	UpdateMerchantByID(ctx context.Context, user *entities.Merchant, updateTypeData map[string]interface{}, code string) error

	CreateMerchantMember(ctx context.Context, user *entities.Users) error

	ListMembersByCode(ctx context.Context, user *[]response.MerchantsMembersResponse, queryParams request.QueryMembersInputRequest) error

	LoginUserByEmailID(ctx context.Context, userData *[]response.UserLoginResponse, queryParams request.LoginUserInputRequest) error

	//
	GetOrderProcessedList(ctx context.Context, OrderProcessedData *[]response.OrdersProcessedResponse) error
	CreateOrderProcessed(ctx context.Context, OrderProcessedData *entities.OrdersProcessed) error
	ListOrderProcessedByID(ctx context.Context, OrderProcessedData *[]response.OrdersProcessedResponse, userID string) error
	UpdateOrderProcessedByID(ctx context.Context, orderProcess *entities.OrdersProcessed, updateTypeData map[string]interface{}, orderProcessRequest request.UpdateOrderProcessedInputRequest) error
}

func (ms *MySQLDBStore) GetOrderProcessedList(ctx context.Context, OrderProcessedData *[]response.OrdersProcessedResponse) error {

	return nil
}
func (ms *MySQLDBStore) CreateOrderProcessed(ctx context.Context, OrderProcessedData *entities.OrdersProcessed) error {

	return nil
}

func (ms *MySQLDBStore) ListOrderProcessedByID(ctx context.Context, orderProcessedData *[]response.OrdersProcessedResponse, userID string) error {
	result := ms.db.Debug().WithContext(ctx).Model(&response.OrdersProcessedResponse{}).Select("id, user_id, order_id, company_id, quantity, status, order_type, created_dt, updated_dt").Where("user_id=?", userID).Scan(&orderProcessedData)
	if result.RowsAffected == 0 {
		return &util.DataNotFound{ErrMessage: fmt.Sprintf(consts.ErrorOrderDataNotFoundCode, userID)}
	}
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}
func (ms *MySQLDBStore) UpdateOrderProcessedByID(ctx context.Context, orderProcess *entities.OrdersProcessed, updateTypeData map[string]interface{}, orderProcessRequest request.UpdateOrderProcessedInputRequest) error {
	var updateFields = make(map[string]interface{})
	for key, val := range updateTypeData {
		updateFields[key] = val
	}

	result := ms.db.Debug().WithContext(ctx).Model(&orderProcess).Where("user_id=? AND order_id=? and status=?", orderProcessRequest.UserID, orderProcessRequest.OrderID, consts.OrderPendingtatus).Omit("user_id", "id", "order_id").Updates(updateFields)

	log.Println("UpdateOrderProcessedByID updated rows: ", result.RowsAffected)
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	} else if result.RowsAffected == 0 {
		err := fmt.Sprintf(consts.ErrorOrderProcessedUpdateType, orderProcessRequest.UserID, orderProcessRequest.OrderID)
		return &util.InternalServer{ErrMessage: err}
	}
	return nil
}

// CreateMerchantMember
func (ms *MySQLDBStore) CreateMerchantMember(ctx context.Context, user *entities.Users) error {
	result := ms.db.Debug().WithContext(ctx).Create(&user)
	err := result.Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates") {
			_userMsg := fmt.Sprintf(consts.ErrUserAlreadyExists, user.FkCode, user.Email)
			return &util.BadRequest{ErrMessage: _userMsg}
		} else {
			return &util.InternalServer{ErrMessage: err.Error()}
		}
	}

	return nil
}

// UpdateMerchantByID
func (ms *MySQLDBStore) UpdateMerchantByID(ctx context.Context, user *entities.Merchant, updateTypeData map[string]interface{}, code string) error {

	var updateFields = make(map[string]interface{})
	for key, val := range updateTypeData {
		updateFields[key] = val
	}

	result := ms.db.Debug().WithContext(ctx).Model(&user).Where("code=?", code).Omit("code", "id").Updates(updateFields)

	log.Println("UpdateByID updated rows: ", result.RowsAffected)
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	} else if result.RowsAffected == 0 {
		err := fmt.Sprintf(consts.ErrorUpdateType, code)
		return &util.InternalServer{ErrMessage: err}
	}
	return nil
}

// CreateMerchant
func (ms *MySQLDBStore) CreateMerchant(ctx context.Context, merchant *entities.Merchant) error {
	result := ms.db.Debug().WithContext(ctx).Create(&merchant)
	err := result.Error
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			_userMsg := fmt.Sprintf(consts.ErrMerchantAlreadyExists, merchant.Code)
			return &util.BadRequest{ErrMessage: _userMsg}
		} else {
			return &util.InternalServer{ErrMessage: err.Error()}
		}
	}

	return nil
}

// ListMerchantByID
func (ms *MySQLDBStore) ListMerchantByID(ctx context.Context, merchantData *[]response.MerchantResponse, code string) error {

	log.Println("ListMerchantByID ")
	result := ms.db.Debug().WithContext(ctx).Model(&response.MerchantResponse{}).Select("code, name, address, status, created_at, updated_at").Where("code=?", code).Scan(&merchantData)
	if result.RowsAffected == 0 {
		return &util.DataNotFound{ErrMessage: fmt.Sprintf(consts.ErrorDataNotFoundCode, code)}
	}
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

// ListMembersByCode
func (ms *MySQLDBStore) LoginUserByEmailID(ctx context.Context, userData *[]response.UserLoginResponse, queryParams request.LoginUserInputRequest) error {

	result := ms.db.Debug().WithContext(ctx).Model(&response.UserLoginResponse{}).Select("users.fk_code, users.first_name, users.last_name, users.email, users.mobile, users.password, users.is_active, users.created_at, merchants.name as MerchantName").Joins("left join merchants on merchants.code = users.fk_code").Where("fk_code=? AND users.email=?", queryParams.Code, queryParams.Email).Scan(&userData)

	if result.RowsAffected == 0 {
		return &util.DataNotFound{ErrMessage: fmt.Sprintf(consts.ErrorUserNotFoundCode, queryParams.Code)}
	}

	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

// ListMembersByCode
func (ms *MySQLDBStore) ListMembersByCode(ctx context.Context, merchant *[]response.MerchantsMembersResponse, queryParams request.QueryMembersInputRequest) error {

	result := ms.db.Debug().WithContext(ctx).Model(&response.MerchantsMembersResponse{}).Select("users.fk_code, users.first_name, users.last_name, users.email, users.mobile, users.is_active, users.created_at, merchants.name as MerchantName").Joins("left join merchants on merchants.code = users.fk_code").Where("fk_code=?", queryParams.Code).Limit(queryParams.Limit).Offset(queryParams.Skip).Scan(&merchant)
	if result.RowsAffected == 0 {
		return &util.DataNotFound{ErrMessage: fmt.Sprintf(consts.ErrorDataNotFoundCode, queryParams.Code)}
	}
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

// GetMerchantList
func (ms *MySQLDBStore) GetMerchantList(ctx context.Context, merchantData *[]response.MerchantResponse) error {
	result := ms.db.WithContext(ctx).Model(&response.MerchantResponse{}).Select("code,  name, address, status, created_at, updated_at").Find(&merchantData)
	err := result.Error
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}
