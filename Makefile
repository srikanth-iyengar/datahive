.PHONY: up down stop fmt

SHELL=/bin/bash

RSRC?="resourcemanager nodemanager datanode namenode kafka dev-ingestion spark minio elastic kibana"
OPTS?=

up:
	@ echo ${RSRC}
	@ cd docker; docker-compose up ${OPTS} -d `echo ${RSRC}`
	@ sleep 6
	@ docker exec -i -t docker_datanode_1 hadoop fs -chmod 777 /

down:
	@cd docker; docker-compose down

stop:
	@cd docker; docker-compose stop ${RSRC}

fmt:
	@echo "Formatting golang files"
	@cd ./app/api; gofmt -l -s -w .

proto:
	@cd ./app/api/pkg ; protoc proto/* --go_out=. --go-grpc_out=.

start-api:
	@ /bin/bash -c "source ./.env ; cd ./app/api ; go run main.go"
