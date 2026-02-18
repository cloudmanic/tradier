BINARY_NAME=tradier
BUILD_DIR=build

.PHONY: build test clean install lint vet fmt run help

## build: Compile the binary
build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

## test: Run all unit tests
test:
	go test ./... -count=1

## test-verbose: Run all unit tests with verbose output
test-verbose:
	go test ./... -v -count=1

## test-cover: Run tests with coverage report
test-cover:
	go test ./... -coverprofile=$(BUILD_DIR)/coverage.out
	go tool cover -func=$(BUILD_DIR)/coverage.out

## test-cover-html: Open coverage report in browser
test-cover-html: test-cover
	go tool cover -html=$(BUILD_DIR)/coverage.out

## clean: Remove build artifacts
clean:
	rm -rf $(BUILD_DIR)

## install: Build and install the binary to GOPATH/bin
install:
	go install .

## vet: Run go vet
vet:
	go vet ./...

## fmt: Format all Go files
fmt:
	go fmt ./...

## lint: Run vet and check formatting
lint: vet
	@test -z "$$(gofmt -l .)" || (echo "Files need formatting:" && gofmt -l . && exit 1)

## tidy: Tidy and verify module dependencies
tidy:
	go mod tidy
	go mod verify

## run: Build and run with optional ARGS (e.g. make run ARGS="markets quotes --symbols AAPL")
run: build
	./$(BUILD_DIR)/$(BINARY_NAME) $(ARGS)

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@sed -n 's/^## //p' $(MAKEFILE_LIST) | column -t -s ':'
