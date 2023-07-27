package query_test

import (
	"context"
	"errors"
	"grpcequitiesapi/internals/adapter/pgsql/entities"
	"grpcequitiesapi/mocks"
	"grpcequitiesapi/pkg/v1/models/response"

	"testing"

	"github.com/golang/mock/gomock"
)

func TestMockMySQLDBStoreAccess_CreateMerchant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBStoreAccess := mocks.NewMockMySQLDBStoreAccess(ctrl)

	// Set up the expected inputs and outputs
	ctx := context.Background()
	mockMerchant := &entities.Merchant{}

	// Set up the expected behavior
	expectedError := errors.New("error creating merchant")
	mockDBStoreAccess.EXPECT().CreateMerchant(ctx, mockMerchant).Return(expectedError)

	// Call the method under test
	err := mockDBStoreAccess.CreateMerchant(ctx, mockMerchant)

	// Verify the result
	if err != expectedError {
		t.Errorf("unexpected error: got %v, want %v", err, expectedError)
	}
}

func Test_GetMerchantList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDBStorer := mocks.NewMockDBStorer(ctrl)

	mockDB := mocks.NewMockMySQLDBStoreAccess(ctrl)

	// Set up test data
	ctx := context.Background()
	merchantList := getMerchantList() //[]response.MerchantResponse{} // Define the expected response data

	// Set up expectations
	mockDB.EXPECT().GetMerchantList(ctx, &merchantList).Return(nil) // Expect the function call with the given context and merchantList pointer

	// Call the function you want to test
	// err := mockDB.GetMerchantList(ctx, &merchantList)
	// mySQLDBStoreAccess := query.NewMySQLDBStore(&pgsql.MySQLDbStore{DB: mockDBStorer})

	mockDBStorer.EXPECT().DBConn("").Return()
	// Validate the results
	// if err != nil {
	// 	t.Errorf("GetMerchantList returned an error: %v", err)
	// }

	// Add additional assertions for the expected results
	// For example, check the length of the returned merchantList
	expectedLength := 2
	actualLength := len(merchantList)
	if actualLength != expectedLength {
		t.Errorf("Expected merchantList length %d, but got %d", expectedLength, actualLength)
	}
}

func getMerchantList() []response.MerchantResponse {
	data := []response.MerchantResponse{}
	data = append(data, response.MerchantResponse{
		Name: "rediff",
		Code: "coderediff478",
	})
	data = append(data, response.MerchantResponse{
		Name: "google",
		Code: "codegoogle1245",
	})

	return data
}
