// database_test.go
package pgsql_test

import (
	"errors"
	"testing"

	"grpcequitiesapi/internals/adapter/pgsql"
	"grpcequitiesapi/mocks"

	"github.com/golang/mock/gomock"
)

func TestNewDbConnector(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDBConnector(ctrl)

	// Test case for successful DB connection
	mockDB.EXPECT().DBConn(gomock.Any()).Return(nil, nil).Times(1)

	_, err := pgsql.NewDbConnector("postgres://postgres:pass@123PGdb@equities.docker.internal:5441/BSE_Equity_Db?sslmode=disable")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Test case for DB connection error
	mockError := errors.New("some DB connection error")
	mockDB.EXPECT().DBConn(gomock.Any()).Return(nil, mockError).Times(1)

	_, err = pgsql.NewDbConnector("mock-dbs-connection-string")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
