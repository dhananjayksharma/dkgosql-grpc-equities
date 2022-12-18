Steps to run application:


Step 1:
	unzip	dkgosql-grpc-equities-v3.demo.zip
	
	cd into folder "apps"

Step 2:
	Need to create database and import 
	
	apps/docs/databases/BSE_Equity_Db.sql
	
	Change database credential:
	
	apps/config-local.yml
	
Step 3:
	cd into folder "apps"
	
	go mod tidy

	run app:
		make all
	
	test app:
		go test ./... -cover
			
		coverage: 61.8% of statements

STEP 4:

Test API with Postman collection:
	apps/docs/dkgosql-grpc-equities.postman_collection.json

OR

Test API on using curl:
 
