.PHONY: build test lint install clean

build:
	go build -o bin/exact-online-activities-pp-cli ./cmd/exact-online-activities-pp-cli

test:
	go test ./...

lint:
	golangci-lint run

install:
	go install ./cmd/exact-online-activities-pp-cli

clean:
	rm -rf bin/

build-mcp:
	go build -o bin/exact-online-activities-pp-mcp ./cmd/exact-online-activities-pp-mcp

install-mcp:
	go install ./cmd/exact-online-activities-pp-mcp

build-all: build build-mcp
