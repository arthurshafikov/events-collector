include .env

BIN := "./.bin/app"
DOCKER_COMPOSE_FILE := "./deployments/docker-compose.yml"
DOCKER_COMPOSE_TEST_FILE := "./deployments/docker-compose.tests.yml"

CLICKHOUSE_DSN := "clickhouse://clickhouse:9000?username=${CLICKHOUSE_USERNAME}&password=${CLICKHOUSE_PASSWORD}&database=${CLICKHOUSE_DATABASE}&x-multi-statement=true"
APP_NAME := "events-collector"
MIGRATIONS_FOLDER := "/go/src/migrations"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

buildbin:
	go build -a -o $(BIN) -ldflags "$(LDFLAGS)" cmd/main.go

run: buildbin 
	 $(BIN)

test: 
	go test --short -race ./internal/... ./pkg/...

.PHONY: buildbin test

generate:
	protoc -I=api --go_out=internal/transport/grpc/generated --go-grpc_out=internal/transport/grpc/generated api/CollectorService.proto

build:
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p ${APP_NAME} up -d --build

up:
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p ${APP_NAME} up -d

stop:
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p ${APP_NAME} stop

migrate:
	docker exec ${APP_NAME}_app_1 migrate -database ${CLICKHOUSE_DSN} -path ${MIGRATIONS_FOLDER} up
