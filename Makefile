BIN := "./.bin/app"
DOCKER_COMPOSE_FILE := "./deployments/docker-compose.yml"
DOCKER_COMPOSE_TEST_FILE := "./deployments/docker-compose.tests.yml"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -a -o $(BIN) -ldflags "$(LDFLAGS)" cmd/main.go

run: build 
	 $(BIN)

test: 
	go test --short -race ./internal/... ./pkg/...

.PHONY: build test

generate:
	protoc -I=api --go_out=internal/transport/grpc/generated --go-grpc_out=internal/transport/grpc/generated api/CollectorService.proto
