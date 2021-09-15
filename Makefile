# https://gist.github.com/thomaspoignant/5b72d579bd5f311904d973652180c705
NAME=letters-api
EXT_PORT?=4000

BUILD_DATE  := $(shell date +'%Y-%m-%dT%H:%M:%S%z')
BUILD_HOST  := $(shell hostname)
GIT_URL  	:= $(shell git config --get remote.origin.url)
BRANCH  	:= $(shell git rev-parse --abbrev-ref HEAD)
VERSION  	:= $(shell echo "$(shell git rev-parse --abbrev-ref HEAD)" | cut -c1-2)-$(shell git rev-parse --short=8 HEAD)

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test build vendor

all: help

## Build:
build: ## Build your project
	go build -ldflags "\
		-X main.buildDate=$(BUILD_DATE) \
		-X main.buildHost=$(BUILD_HOST) \
		-X main.gitURL=$(GIT_URL) \
		-X main.branch=$(BRANCH) \
		-X main.version=$(VERSION)" \
		-o ./main ./cmd/main.go

upgrade: ## Upgrade module's go version and dependencies
	@echo $(shell head -n 1 go.mod) > go.mod
	@make vendor


vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	go get ./...
	go mod tidy
	go mod vendor

## Run:
run: ## Run the application
	go run -ldflags "\
		-X main.buildDate=$(BUILD_DATE) \
		-X main.buildHost=$(BUILD_HOST) \
		-X main.gitURL=$(GIT_URL) \
		-X main.branch=$(BRANCH) \
		-X main.version=$(VERSION)" \
		./cmd/main.go

## Test:
test: ## Run the tests of the project
	./scripts/unit_test.sh

## Lint:
lint: lint-go lint-helm ## Run all available linters

report-card: ## Generate a report card of various issues
	@go install github.com/gojp/goreportcard/cmd/goreportcard-cli@latest
	goreportcard-cli -v

lint-go: ## Use golintci-lint on your project
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run --deadline=65s --issues-exit-code 0 ./...

lint-helm: ## Use helm lint on the helm charts of your projects
	helm lint deployment/helm

## Docker:
docker-build: ## Use the dockerfile to build the container
	docker build --rm --tag $(NAME):local -f ./Dockerfile . \
		--build-arg BUILD_DATE=$(BUILD_DATE) \
		--build-arg BUILD_HOST=$(BUILD_HOST) \
		--build-arg GIT_URL=$(GIT_URL) \
		--build-arg BRANCH=$(BRANCH) \
		--build-arg VERSION=$(VERSION)

docker-run: ## Use the dockerfile to build run container
	@make docker-build
	docker run -it --rm -e dbconn=mongodb+srv://L3O-svc-user:RXOh3AiCv3HUefu8@bedaring-cluster-01-f6szy.gcp.mongodb.net/test?retryWrites=true -v "$(shell pwd)/config/:/config" -p $(EXT_PORT):3000 --name=$(NAME) $(NAME):local

docker-clean: ## Clean up docker (Stop containers, prune network, containers and images, remove volumes)
	@docker stop $(shell docker ps -a -q) || true;
	@docker network prune -f || true;
	@docker container prune -f || true;
	@docker image prune -af || true;
	@docker volume rm $(shell docker volume ls -qf dangling=true) || true;

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
