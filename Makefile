.PHONY: up down stop fmt

RSRC?="resourcemanager nodemanager datanode namenode kafka dev-ingestion"

up:
	@ echo ${RSRC}
	@cd docker; docker-compose up -d `echo ${RSRC}`

down:
	@cd docker; docker-compose down

stop:
	@cd docker; docker-compose stop ${RSRC}

fmt:
	@echo "Formatting golang files"
	@cd ./app/api; gofmt -l -s -w .

proto:
	@cd ./app/api/pkg ; protoc proto/KafkaService.proto --go_out=. --go-grpc_out=.

start-api:
	@ /bin/bash -c "source ./.env ; cd ./app/api ; go run main.go"
