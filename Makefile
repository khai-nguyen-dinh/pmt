SHELL := /bin/bash
PLATFORM := $(shell go env GOOS)
EXECUTABLE=pmt
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64
VERSION=$(shell git describe --tags --always)
DEP_VERSION:=$(shell dep version | sed -n 2p 2>/dev/null)
.PHONY: all build clean

all: install-dep build-all
build: auto windows linux darwin build-all## Build binaries

build-all: windows linux darwin## Build all plaform binaries

install-dep:
    $(info Checking version dep ...)
    ifndef DEP_VERSION
       $(info go dep is not install. Installing...)
       $(shell go get -u github.com/golang/dep/cmd/dep)
    endif
    $(info dep version: $(shell dep version | sed -n 2p| awk '{print $$3}'))
auto: $(PLATFORM)

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOARCH=amd64 go build -i -v -o ./build/$(WINDOWS)

$(LINUX):
	env GOOS=linux CGO_ENABLED=1 GOARCH=amd64 go build -i -v -o ./build/$(LINUX)

$(DARWIN):linux
	env GOOS=darwin CGO_ENABLED=1 GOARCH=amd64 go build -i -v -o ./build/$(DARWIN)

clean: ## Remove previous build
	rm -rf ./build

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
