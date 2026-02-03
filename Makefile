APP_NAME=ysh
TARGET_DIR=bin

default: run

build:
	@go build -o $(TARGET_DIR)/$(APP_NAME) cmd/ysh/main.go

clean:
	@rm -rf $(TARGET_DIR)

run: build
	@./$(TARGET_DIR)/$(APP_NAME)

.PHONY: default build clean run
