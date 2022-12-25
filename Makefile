.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: lint
	go vet ./...
.PHONY:vet

cover: vet
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
.PHONY:cover

test: vet
	go test -v ./...
.PHONY:test

build: test
	go build -o residencia cmd/web/*.go
.PHONY:build

run:
	go build -o residencia cmd/web/*.go && ./residencia
.PHONY:run
