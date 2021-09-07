LOCAL_BIN:=$(CURDIR)/bin

.PHONY: run, lint, test, mocks
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org|direct

build:
all: generate test build

.PHONY: deps
deps:
	go get -u github.com/onsi/ginkgo
	go get -u github.com/onsi/gomega
	go get -u github.com/golang/mock
	go get -u github.com/rs/zerolog/log
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/rs/zerolog/log
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

.PHONY: build
build: deps
	CGO_ENABLED=0 go build -o $(LOCAL_BIN)/ova-conference-api cmd/ova-conference-api/main.go

.PHONY: generate
generate:
	protoc --proto_path=. -I vendor.protogen \
	--go_out=pkg/api --go_opt=paths=import \
	--go-grpc_out=pkg/api --go-grpc_opt=paths=import \
	api/api.proto

.PHONY: run
run:
	go run ./cmd/ova-conference-api

.PHONY: test
test: mocks
	go test ./...

.PHONY: migration
migration:
	goose postgres "user=postgres password=postgres  dbname=ozon sslmode=disable" up

.PHONY: lint
lint:
	golangci-lint run

.PHONY: mocks
mocks:
	rm -rf ./internal/mocks/mock_*
	mockgen -source=./internal/utils/repo/repo.go -destination=./internal/utils/mocks/repo_mock.go -package mocks
	mockgen -source=./internal/kafka/producer.go -destination=./internal/infrastructure_mocks/producer_mock.go -package infrastructure_mocks
	mockgen -source=./internal/metrics/metrics.go -destination=./internal/infrastructure_mocks/metrics_mock.go -package infrastructure_mocks

.PHONY: docker-build
docker-build:
	docker-compose build