.PHONY: all
all: genproto

.PHONY: clean

clean: 
	@echo "[OK] removed nothing!"
 
.PHONY: genproto

genproto:
	protoc -I equities/ --go_out=equities/. --go_opt=paths=source_relative --go-grpc_out=equities/. --go-grpc_opt=paths=source_relative equities/*.proto

	@#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative equities/equities.proto  

	@#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto  
	@# equities/orders.proto
	
.PHONY: gserver

gserver:
	go mod tidy
	go run equities_server/equities_server.go 
	
