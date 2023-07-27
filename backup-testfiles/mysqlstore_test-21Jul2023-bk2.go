ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDBConnector := mocks.NewMockDBConnector(ctrl)
	mockDBStoreAccess := mocks.NewMockMySQLDBStoreAccess(ctrl)
	mySQLDBStore, err := pgsql.NewDbConnector("")
	// Create a new instance of MySQLDBStore using the mocked DBConnector
	dbStore := query.NewMySQLDBStore(mySQLDBStore.DB)

	// Create a context for testing (can be an empty context in this case)
	ctx := context.TODO()

	// Set up the expected response from the database query
	expectedMerchantList := []response.MerchantResponse{
		{Code: "MERCHANT_001", Name: "Merchant 1", Address: "Address 1"},
		{Code: "MERCHANT_002", Name: "Merchant 2", Address: "Address 2"},
	}

	// Set up the expected call to the mockDBStoreAccess.GetMerchantList function
	mockDBStoreAccess.EXPECT().GetMerchantList(ctx, gomock.Any()).DoAndReturn(
		func(ctx context.Context, merchantData *[]response.MerchantResponse) error {
			// Copy the expected data to the merchantData pointer
			*merchantData = expectedMerchantList
			return nil
		},
	)

	// Set up the expected call to the mockDBConnector.DBConn function
	mockDBConnector.EXPECT().DBConn(gomock.Any()).Return(nil, nil)

	// Replace the original DBStoreAccess with the mockDBStoreAccess
	dbStore.DB = mockDBStoreAccess

	// Call the GetMerchantList function
	var merchantData []response.MerchantResponse
	err = dbStore.GetMerchantList(ctx, &merchantData)

	// Check for any errors
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Compare the returned data with the expected data
	if !reflect.DeepEqual(merchantData, expectedMerchantList) {
		t.Fatalf("Returned data does not match the expected data")
	}

	// // Check if the returned merchants match the expected merchants
	// if len(merchants) != len(expectedMerchants) {
	// 	t.Fatalf("Expected %d merchants, but got %d", len(expectedMerchants), len(merchants))
	// }

	// for i := 0; i < len(expectedMerchants); i++ {
	// 	if merchants[i].Code != expectedMerchants[i].Code || merchants[i].Name != expectedMerchants[i].Name {
	// 		t.Fatalf("Expected merchant: %v, but got: %v", expectedMerchants[i], merchants[i])
	// 	}
	// }
