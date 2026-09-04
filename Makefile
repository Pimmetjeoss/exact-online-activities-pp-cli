.PHONY: build test lint install clean

BIN_EXT := $(if $(filter windows,$(shell go env GOOS)),.exe,)

build:
	go build -o bin/exact-online-activities-pp-cli$(BIN_EXT) ./cmd/exact-online-activities-pp-cli

test:
	go test ./...

lint:
	golangci-lint run

install:
	go install ./cmd/exact-online-activities-pp-cli

clean:
	rm -rf bin/

build-mcp:
	go build -o bin/exact-online-activities-pp-mcp$(BIN_EXT) ./cmd/exact-online-activities-pp-mcp

install-mcp:
	go install ./cmd/exact-online-activities-pp-mcp

build-all: build build-mcp
