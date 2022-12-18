package users

import (
	"context"
	"errors"
	"fmt"
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

	auth "github.com/dhananjayksharma/dkgo-auth/auth"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	ListMembersByCode(c *gin.Context) (models.Response, error)
	CreateMerchantMember(c *gin.Context) (models.Response, error)
	LoginMember(c *gin.Context) (models.Response, error)
	RefreshToken(c *gin.Context) (models.Response, error)
}

type userService struct {
	db query.MySQLDBStoreAccess
}

func NewUserService(db query.MySQLDBStoreAccess) UserService {
	return &userService{db: db}
}

func (service userService) RefreshToken(c *gin.Context) (models.Response, error) {
	var resp = models.Response{}
	tokenString := c.GetHeader("RefreshToken")
	if tokenString == "" {
		err := errors.New("request does not contain a refresh token")
		if err != nil {
			return resp, err
		}
	}
	err := auth.ValidateRefreshToken(tokenString)
	if err != nil {
		fmt.Printf("auth.ValidateRefreshToken %v", err)
		return resp, err
	}

	claims, err := auth.GetRefreshClaim(tokenString)
	fmt.Printf("Username refresh token %v", claims.Username)
	fmt.Printf("Email refresh token %v", claims.Email)

	var responseUserLogin []response.UserLoginResponse

	userName := claims.Username
	token, refreshToken, err := auth.GenerateJWT(claims.Email, userName)
	if err != nil {
		return resp, err
	}
	responseUserLogin = append(responseUserLogin, response.UserLoginResponse{Token: token, ResetToken: refreshToken})
	// responseUserLogin[0].Token = token
	// responseUserLogin[0].ResetToken = refreshToken
	resp.Data = responseUserLogin
	resp.Message = consts.TokenRegeneatedSuccess
	return resp, nil
}

func (service userService) LoginMember(c *gin.Context) (models.Response, error) {
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var resp = models.Response{}

	var loginUserRequest request.LoginUserInputRequest
	if err := c.BindJSON(&loginUserRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}
	var responseMerchant []response.MerchantResponse
	err := service.db.ListMerchantByID(ctx, &responseMerchant, loginUserRequest.Code)
	if err != nil {
		return resp, err
	}
	// fmt.Printf("%v", len(responseMerchant))

	if len(responseMerchant) == 0 {
		err = errors.New(fmt.Sprintf(consts.ErrorDataNotFoundCode, loginUserRequest.Code))
		return resp, err
	}

	// hashpassword, _ := util.HashPassword(loginUserRequest.Password)

	var responseUserLogin []response.UserLoginResponse
	err = service.db.LoginUserByEmailID(ctx, &responseUserLogin, loginUserRequest)
	if err != nil {
		return resp, err
	}

	passsword_match := util.CheckPasswordHash(loginUserRequest.Password, responseUserLogin[0].Password)
	// fmt.Println("passsword_match status:", passsword_match)

	if !passsword_match {
		err = errors.New("Password did not matching")
		return resp, err
	}
	userName := fmt.Sprintf("%s %s", responseUserLogin[0].FirstName, responseUserLogin[0].LastName)
	token, refreshToken, err := auth.GenerateJWT(responseUserLogin[0].Email, userName)
	if err != nil {
		return resp, err
	}
	responseUserLogin[0].Token = token
	responseUserLogin[0].ResetToken = refreshToken
	resp.Data = responseUserLogin
	resp.Message = consts.UserLoginSuccess
	return resp, nil
}

func (service userService) CreateMerchantMember(c *gin.Context) (models.Response, error) {
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var resp = models.Response{}

	code := strings.Trim(c.Param("code"), "")
	if len(code) == 0 {
		err := errors.New(consts.InvalidMerchantCode)
		return resp, err
	}

	var responseMerchant []response.MerchantResponse
	err := service.db.ListMerchantByID(ctx, &responseMerchant, code)
	if err != nil {
		return resp, err
	}
	if len(responseMerchant) == 0 {
		err = errors.New(fmt.Sprintf(consts.ErrorDataNotFoundCode, code))
		return resp, err
	}

	var addUserRequest request.AddUserInputRequest
	if err := c.BindJSON(&addUserRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}
	hashpassword, _ := util.HashPassword(addUserRequest.Password)
	var status uint8
	status = uint8(consts.ActiveStatus)
	addUser := entities.Users{
		IsActive:  &status,
		FirstName: addUserRequest.FirstName,
		LastName:  addUserRequest.LastName,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		FkCode:    code,
		Password:  hashpassword,
		Email:     addUserRequest.Email,
	}

	err = service.db.CreateMerchantMember(ctx, &addUser)
	if err != nil {
		return resp, err
	}

	var newSpotlightMaster []response.UsersResponse
	newSpotlightMaster = append(newSpotlightMaster, response.UsersResponse{
		IsActive:       &status,
		FirstName:      addUserRequest.FirstName,
		LastName:       addUserRequest.LastName,
		FkMerchantCode: addUserRequest.Code,
		Email:          addUserRequest.Email,
	})

	resp.Data = newSpotlightMaster
	resp.Message = consts.UserAddedSuccess
	return resp, nil
}

func (srv *userService) ListMembersByCode(c *gin.Context) (models.Response, error) {
	var err error
	// set context
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var userData []response.MerchantsMembersResponse
	var resp = models.Response{}

	code := strings.Trim(c.Param("code"), "")
	if len(code) == 0 {
		err = errors.New(consts.InvalidUserDocId)
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

	var queryParams = request.QueryMembersInputRequest{Code: code, Limit: int(page_limit), Skip: int(skip_number)}

	err = srv.db.ListMembersByCode(ctx, &userData, queryParams)
	if err != nil {
		return resp, err
	}
	var responseUser []response.MerchantsMembersResponse
	for _, row := range userData {
		responseUser = append(responseUser, response.MerchantsMembersResponse{
			IsActive:     row.IsActive,
			FirstName:    row.FirstName,
			LastName:     row.LastName,
			Email:        row.Email,
			Mobile:       row.Mobile,
			MerchantName: row.MerchantName,
			FkCode:       row.FkCode,
			CreatedAt:    row.CreatedAt,
		})
	}

	resp.Data = responseUser
	return resp, nil
}
