package handlers

import (
	"bytes"
	"dkgosql-merchant-service-v4/internals/adapter/mysql/query"
	"dkgosql-merchant-service-v4/pkg/v1/models/merchants"
	"dkgosql-merchant-service-v4/pkg/v1/models/response"
	"dkgosql-merchant-service-v4/pkg/v1/models/users"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetMerchantList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var (
		router *gin.Engine
	)
	beforeEach := func(t *testing.T) {
		srvMerchant := merchants.NewMerchantService(&query.MockMySQLDBStore{})
		srvUser := users.NewUserService(&query.MockMySQLDBStore{})
		router = SetupRouter(srvMerchant, srvUser, nil)
	}
	t.Run("get-merchant-list", func(t *testing.T) {
		beforeEach(t)
		merchantsList := []response.MerchantResponse{
			{
				Code: "9876541",
				Name: "Sony",
			},
			{
				Code: "123456789",
				Name: "Cisca",
			},
		}
		rr := httptest.NewRecorder()

		request, err := http.NewRequest(http.MethodGet, "/merchants/merchants", nil)
		request.Header.Set("Token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRoYW5hbmpheSBzaGFybWEiLCJlbWFpbCI6ImRoYW5hbmpheTMzMzMxQGdtYWlsLmNvbSIsImV4cCI6MTY2MDgyMzE5MH0.1GjfH0aq5qFF-Mp9x83Gz9X_B0ua2PyhX01uAVjiiOQ")
		assert.NoError(t, err)
		router.ServeHTTP(rr, request)
		assert.Equal(t, http.StatusOK, rr.Code)

		var resp map[string]interface{}
		err = json.NewDecoder(rr.Body).Decode(&resp)
		data, _ := resp["data"]
		jsonStr, err := json.Marshal(data)
		assert.NoError(t, err)

		// Convert json string to struct
		var out []response.MerchantResponse
		err = json.Unmarshal(jsonStr, &out)
		assert.NoError(t, err)

		assert.Equal(t, len(out), len(merchantsList))
	})

}

func TestCreateMerchant(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var (
		router *gin.Engine
	)
	beforeEach := func(t *testing.T) {
		srvMerchant := merchants.NewMerchantService(&query.MockMySQLDBStore{})
		srvUser := users.NewUserService(&query.MockMySQLDBStore{})
		router = SetupRouter(srvMerchant, srvUser, nil)
	}
	t.Run("create-merchant", func(t *testing.T) {
		beforeEach(t)
		testCases := []struct {
			body   []byte
			want   string
			status int
		}{
			{
				body: []byte(`{
					"name":"Sony New ltd", "address":"Mumbai", "code":"cadjq02gqpmvljdra98"
					}`),
				want:   "Merchant added successfully",
				status: http.StatusOK,
			},
		}

		for _, tt := range testCases {
			rr := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodPost, "/merchants/merchants", bytes.NewBufferString(string(tt.body)))
			assert.NoError(t, err)

			router.ServeHTTP(rr, request)
			assert.Equal(t, tt.status, rr.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(rr.Body).Decode(&resp)
			assert.NoError(t, err)
			// fmt.Printf("created data: %v", resp)
			msg, _ := resp["message"]
			assert.Equal(t, tt.want, msg)
		}
	})

}
