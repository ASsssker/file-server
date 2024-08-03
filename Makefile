MAIN_PACKAGE_PATH := ./cmd/
BINARY_NAME := file-server

.PHONY: build
build:
	go build -o=./${BINARY_NAME} ${MAIN_PACKAGE_PATH}

.PHONY: run
run: build
	./${BINARY_NAME}