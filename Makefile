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

# Build a release image
.PHONY: docker-image
docker-image:
	@DOCKER_BUILDKIT=1 docker build \
		--build-arg ARCH="${ARCH}" \
		-t scalarorg/evms-indexer .		

docker-up:
	docker compose up -d

docker-down:
	docker compose down

