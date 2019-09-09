NAME     := mav
VERSION  := v0.0.2
REVISION := $(shell git rev-parse --short HEAD)

GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOBUILD_ENV=$(GOOS=linux GOARCH=amd64)

GO_CMD=$(shell which go)
GO_SRC=$(shell pwd)/src
GO_BIN=$(shell pwd)/bin

export GO111MODULE=on

.PHONY: init
init: clean

.PHONY: build
build: init
	cd $(GO_SRC)/client && $(GO_CMD) build -o ../../bin/client
	cd $(GO_SRC)/server && $(GO_CMD) build -o ../../bin/server
	cd $(GO_SRC)/tmp && $(GO_CMD) build -o ../../bin/echo

.PHONY: run_srv
run_srv:
	$(GO_CMD) run $(GO_SRC)/server/main.go

.PHONY: run_cli
run_cli:
	$(GO_CMD) run $(GO_SRC)/client/main.go

.PHONY: test
test:
	:

.PHONY: install
install:
	:

.PHONY: env
env:
	$(GO_CMD) env

.PHONY: clean
clean:
	rm -f $(GO_BIN)/^[a-z]+
