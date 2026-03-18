.PHONY: all build test clean tidy

APP_NAME = app
BUILD_DIR = .

all: clean tidy build

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) cmd/server/main.go

test:
	go test -v ./...

clean:
	rm -f $(BUILD_DIR)/$(APP_NAME)
	go clean

tidy:
	go mod tidy
