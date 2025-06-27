ARCH := x86_64
ifeq ($(shell uname -m), arm64)
ARCH := aarch64
endif

.PHONY: all
all: goimports lint build docker-image

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

.PHONY: goimports
goimports:
	@echo "--> Running goimports"
	goimports -w .

.PHONY: lint
# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

# Build the project with release flags
.PHONY: build
build: go.sum
		go build -o ./bin/indexer -mod=readonly ./main.go

# Development commands
.PHONY: dev
dev: build
	set -a; source .env; set +a; ./bin/indexer

# Local development with PostgreSQL
.PHONY: db-up
db-up:
	@echo "--> Starting PostgreSQL with Docker Compose"
	docker compose -f compose.yml up -d indexer-db
	@echo "--> Waiting for PostgreSQL to be ready..."
	@until docker compose -f compose.yml exec -T indexer-db pg_isready -U vault_indexer; do sleep 2; done
	@echo "--> PostgreSQL is ready!"

.PHONY: db-down
db-down:
	@echo "--> Stopping PostgreSQL"
	docker compose -f compose.yml down --volumes indexer-db

.PHONY: local-run
local-run:
	@echo "--> Running evms-indexer locally with go run"
	set -a; source .env; set +a; go run main.go

.PHONY: local-dev
local-dev: db-up local-run

.PHONY: local-clean
local-clean: db-down
	@echo "--> Cleaning up local development environment"

# Development utilities
.PHONY: deps
deps:
	@echo "--> Installing dependencies"
	go mod download
	go mod tidy

.PHONY: test
test:
	@echo "--> Running tests"
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	@echo "--> Running tests with coverage"
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "--> Coverage report generated: coverage.html"

.PHONY: clean
clean:
	@echo "--> Cleaning build artifacts"
	rm -rf ./bin/
	rm -f coverage.out coverage.html

.PHONY: logs
logs:
	@echo "--> Showing PostgreSQL logs"
	docker compose -f compose.yml logs -f indexer-db

.PHONY: db-shell
db-shell:
	@echo "--> Opening PostgreSQL shell"
	docker compose -f compose.yml exec indexer-db psql -U vault_indexer -d vault-indexer-db

# Build a release image
.PHONY: docker-image
docker-image:
	@DOCKER_BUILDKIT=1 docker build \
		--build-arg ARCH="${ARCH}" \
		-t scalarorg/vault-indexer .		

compose-up:
	docker compose -f compose.yml up -d
compose-down:
	docker compose -f compose.yml down --volumes

