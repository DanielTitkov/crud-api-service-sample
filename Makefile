NAME := app
DEV_CONFIG_PATH := ./configs/dev.yml

# Docker
DOCKER_APP_FILENAME ?= deployments/docker/Dockerfile
DOCKER_COMPOSE_FILE ?= deployments/docker-compose/docker-compose.yml

CONFIG_PATH ?= ./configs/new.yml

# Build
BUILD_CMD ?= CGO_ENABLED=0 go build -o bin/${NAME} -ldflags '-v -w -s' ./cmd/${NAME}

.PHONY: run
run: entgen
	go run cmd/$(NAME)/main.go ${DEV_CONFIG_PATH}

.PHONY: db
db:
	cd deployments/dev && docker-compose up -d --force-recreate --build --remove-orphans --always-recreate-deps --renew-anon-volumes

.PHONY: entgen
entgen:
	cd internal/repository/entgo && go generate ./ent

.PHONY: build
build:
	echo "building"
	${BUILD_CMD}
	echo "build done"

.PHONY: up
up:
	docker-compose -f ${DOCKER_COMPOSE_FILE} up --build

.PHONY: lint
lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.29.0
	./bin/golangci-lint run -v


.PHONY: check
check: lint test
	echo "check done"

.PHONY: test
test:
	go test ./... 
