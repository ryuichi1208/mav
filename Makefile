NAME     := mav
VERSION  := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)

GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

SRC=src

.PHONY: run
run:
	go run $(SRC)/main.go

.PHONY: test
test:
	:

.PHONY: install
install:
	:

.PHONY: clean
clean:
	:
