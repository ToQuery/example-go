# 从环境变量获取版本号
VERSION ?= 0.0.1
BUILD_ID ?= 0x0001
BUILD_TIME := $(shell date +%FT%T)
OUTPUT_DIR := dist
BIN_NAME := example-go


.PHONY: build
build:
	go build -ldflags "-s -w -X main.Version=$(VERSION) -X main.BuildId=$(BUILD_ID) -X main.BuildTime=$(BUILD_TIME)" -o $(OUTPUT_DIR)/$(BIN_NAME) main.go


.PHONY: build-all
build-all:
	@mkdir -p $(OUTPUT_DIR)
	@platforms="windows linux darwin"; \
	archs="amd64 arm64"; \
	for GOOS in $$platforms; do \
		for GOARCH in $$archs; do \
			EXT=""; \
			if [ "$$GOOS" = "windows" ]; then EXT=".exe"; fi; \
			OUTDIR="$(OUTPUT_DIR)/$${GOOS}_$${GOARCH}"; \
			mkdir -p $$OUTDIR; \
			OUTPUT_FILE="$$OUTDIR/$(BIN_NAME)$$EXT"; \
			echo "Building $$OUTPUT_FILE..."; \
			GOOS=$$GOOS GOARCH=$$GOARCH go build -ldflags "-s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o $$OUTPUT_FILE main.go || exit 1; \
		done; \
	done
