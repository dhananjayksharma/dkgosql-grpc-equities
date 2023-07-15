package companies

import (
	"grpcequitiesapi/internals/adapter/pgsql/entities"
	"grpcequitiesapi/internals/adapter/pgsql/query"
	"grpcequitiesapi/pkg/v1/models"
	"grpcequitiesapi/pkg/v1/models/response"

	"github.com/gin-gonic/gin"
)

type CompanyService interface {
	GetCompanyList(c *gin.Context) (models.Response, error)
	CreateCompany(c *gin.Context) (models.Response, error)
	UpdateCompanyByID(c *gin.Context) (models.Response, error)
}

type companyService struct {
	db query.MySQLDBStoreAccess
}

func NewCompanyService(db query.MySQLDBStoreAccess) CompanyService {
	return &companyService{db: db}
}

// UpdateCompanyByID
func (service companyService) UpdateCompanyByID(c *gin.Context) (models.Response, error) {
	// var err error
	// set context
	// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	// defer cancel()

	var companyData entities.Company
	var resp = models.Response{}

	// var updateCompanyRequest request.UpdateCompanyInputRequest
	// if err := c.BindJSON(&updateCompanyRequest); err != nil {
	// 	return resp, &util.BadRequest{ErrMessage: err.Error()}
	// }

	// code := strings.Trim(c.Param("code"), "")
	// if len(code) == 0 {
	// 	err = errors.New(consts.InvalidCode)
	// 	return resp, err
	// }

	// var responseCompany []response.CompanyResponse
	// err = service.db.ListCompanyByID(ctx, &responseCompany, code)
	// if err != nil {
	// 	return resp, err
	// }
	// if len(responseCompany) == 0 {
	// 	err = errors.New(fmt.Sprintf(consts.ErrorDataNotFoundCode, code))
	// 	return resp, err
	// }

	// var updateTypeData = make(map[string]interface{})
	// updateTypeData["address"] = updateCompanyRequest.Address
	// updateTypeData["name"] = updateCompanyRequest.Name

	// err = service.db.UpdateCompanyByID(ctx, &companyData, updateTypeData, code)
	// if err != nil {
	// 	return resp, err
	// }

	// err = service.db.ListCompanyByID(ctx, &responseCompany, code)
	// if err != nil {
	// 	return resp, err
	// }
	resp.Data = companyData
	resp.Message = "Hard code line 76" //fmt.Sprintf(consts.CompanyUpdatedSuccess, code)
	return resp, nil
}

// CreateCompany
func (srv companyService) CreateCompany(c *gin.Context) (models.Response, error) {
	// set context
	// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	// defer cancel()

	var resp = models.Response{}
	/*var addCompanyRequest request.AddMerchantInputRequest
	if err := c.BindJSON(&addCompanyRequest); err != nil {
		return resp, &util.BadRequest{ErrMessage: err.Error()}
	}*/

	/*currTime := time.Now()
	var status uint8
	status = uint8(consts.ActiveStatus)
	addCompany := entities.Company{
		UpdatedAt: currTime,
		CreatedAt: currTime,
		Code:      addCompanyRequest.Code,
		Status:    &status,
		Name:      addCompanyRequest.Name,
		Address:   addCompanyRequest.Address,
	}*/

	/*err := srv.db.CreateCompany(ctx, &addCompany)
	if err != nil {
		return resp, err
	}*/

	var newCompanyMaster []response.CompanyResponse
	/*newCompanyMaster = append(newCompanyMaster, response.CompanyResponse{
		Name:      addCompanyRequest.Name,
		Code:      addCompanyRequest.Code,
		CreatedAt: currTime.String(),
		Address:   addCompanyRequest.Address,
	})*/

	resp.Data = newCompanyMaster
	resp.Message = "CompanyMaster" //consts.CompanyAddedSuccess
	return resp, nil
}

// GetCompanyList companies
func (srv *companyService) GetCompanyList(c *gin.Context) (models.Response, error) {
	var err error
	// set context
	// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	// defer cancel()
	var companyData []response.CompanyResponse
	var resp = models.Response{}
	if err != nil {
		return resp, err
	}

	// err = srv.db.GetCompanyList(ctx, &companyData)
	// if err != nil {
	// 	return resp, err
	// }

	var outMSM []response.CompanyResponse
	for _, row := range companyData {
		outMSM = append(outMSM, response.CompanyResponse{
			Name:      row.Name,
			Code:      row.Code,
			CreatedAt: row.CreatedAt,
		})
	}
	resp.Data = outMSM
	return resp, nil
}
