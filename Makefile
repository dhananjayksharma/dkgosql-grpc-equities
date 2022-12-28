.PHONY: all
all: genproto

.PHONY: clean

clean: 
	@echo "[OK] removed nothing!"
 
.PHONY: genproto

genproto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative equities/equities.proto  

.PHONY: gserver

gserver:
	go mod tidy
	go run equities_server/equities_server.go 
	