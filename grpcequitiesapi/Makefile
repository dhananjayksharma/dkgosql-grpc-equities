.PHONY: all

BINARY_NAME := dkgosql-merchant-service-v4

include .envrc

all: exportvar deps clean build run
	#run no use as 30Mar2022 runasdocker		test

.PHONY: clean

clean:
	go clean -modcache
	go clean
	@-rm ${BINARY_NAME}-linux
	@-rm -f cover.html cover.out
	@-rm -f ${BINARY_NAME}
	@echo "[OK] Cover file and old build removed!"

.PHONY: exportvar

exportvar:
	# @export MICROSERVICECDEMONEWAPI=local
	@echo "[OK] Export var $(MICROSERVICECDEMONEWAPI)"

.PHONY: deps

deps:
	@-rm -f go.sum
	go mod tidy

.PHONY: run

run:
	@./${BINARY_NAME}
	
	#.PHONY: runasdocker

	#runasdocker:
	#	@echo docker build -t microservice-mediaphoto-service:v1 . && docker run --name services -p 8587:8587 -d microservice-mediaphoto-service:v1
.PHONY: makedoc

makedoc:
	#@dummy


.PHONY: build

build:
	go clean
	@go get -v
	@go build -o ${BINARY_NAME} main.go
	# go build -o ${BINARY_NAME}-linux main.go GOARCH=amd64 GOOS=linux 
	@echo "[OK] Server was build!"

.PHONY: test
 
test:
	go test ./...
	@go test -v	
	@echo "[OK] Test and coverage file was created!"

	#@go tool cover -html=cover.out -o cover.html 
	#-coverprofile cover.out ./...

.PHONY: show_coverage

show_coverage:
	#@go tool cover -html=cover.out
	@echo "[OK] Coverage file opened!"
