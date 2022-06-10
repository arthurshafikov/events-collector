include .env

BIN := "./.bin/app"
DOCKER_COMPOSE_FILE := "./deployments/docker-compose.yml"

APP_NAME := "events-sender"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

buildbin:
	go build -a -o $(BIN) -ldflags "$(LDFLAGS)" cmd/main.go

run: buildbin 
	 $(BIN)

test: 
	go test --short -race ./internal/... ./pkg/...

.PHONY: buildbin test

build:
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p ${APP_NAME} up -d --build

up:
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p ${APP_NAME} up -d

stop:
	docker-compose -f ${DOCKER_COMPOSE_FILE} -p ${APP_NAME} stop
