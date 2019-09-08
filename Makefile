NAME     := mav
VERSION  := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)

GOBIN=$(shell which go)
GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

SRC=src
BIN=bin

.PHONY: run_srv
run_srv:
	$(GOBIN) run $(SRC)/server/main.go

.PHONY: run_cli
run_cli:
	$(GOBIN) run $(SRC)/client/main.go

.PHONY: build
build:
	cd $(SRC)/client && $(GOBIN) build -o ../../bin/client
	cd $(SRC)/server && $(GOBIN) build -o ../../bin/server
	cd $(SRC)/tmp && $(GOBIN) build -o ../../bin/echo

.PHONY: test
test:
	:

.PHONY: install
install:
	:

.PHONY: clean
clean:
	rm -f $(BIN)/client && rm -f $(BIN)/server && rm -f $(BIN)/echo
