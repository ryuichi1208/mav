NAME     := mav
VERSION  := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)

GOBIN=$(shell which go)
GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

SRC=src

.PHONY: run_srv
run_srv:
	$(GOBIN) run $(SRC)/server/main.go

.PHONY: run_cli
run_cli:
	$(GOBIN) run $(SRC)/client/main.go

.PHONY: test
test:
	:

.PHONY: install
install:
	:

.PHONY: clean
clean:
	:
