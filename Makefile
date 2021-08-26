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

.PHONY: mocks
mocks:
	rm -rf ./internal/mocks/mock_*
	mockgen -source=./internal/utils/repo/repo.go -destination=./internal/utils/mocks/repo_mock.go -package mocks