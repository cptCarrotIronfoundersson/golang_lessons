BIN_SCHEDULER := "./bin/scheduler"
BIN_CALENDAR := "./bin/calendar"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN_CALENDAR) -ldflags "$(LDFLAGS)" ./cmd/calendar
	go build -v -o $(BIN_SCHEDULER) -ldflags "$(LDFLAGS)" ./cmd/scheduler

run: build
	$(BIN) --config ./configs/config.toml

run_docker:
	docker-compose -f ./deployments/docker-compose.yaml down --remove-orphans && docker-compose -f ./deployments/docker-compose.yaml up --build -d

migrate-db:
	docker-compose -f ./deployments/docker-compose.yaml up migration --build -d

stop_docker:
	docker-compose -f ./deployments/docker-compose.yaml down --remove-orphans

version: build
	$(BIN) version

test:
	go test -count=1 -race ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin 1.49.0

integration-tests-run:
	docker-compose -f ./deployments/docker-compose.integrationtests.yaml up integration_tests --build

lint: install-lint-deps
	golangci-lint run ./...

up: run_docker migrate-db

down: stop_docker

integration-tests: up integration-tests-run down

.PHONY: build run build-img deploy version test lint
