REPO_ROOT = $(shell pwd)

MODULE := $(REPO_ROOT)/payments/gen
SWAGGER := $(REPO_ROOT)/swagger/swagger.yaml

MODULE_NAME := payments

export GOPATH := $(REPO_ROOT)/gopath
export PATH := $(GOPATH)/bin:$(PATH)

.PHONY: all
all: deps generate build

.PHONY: generate
generate:
	cd $(MODULE); swagger generate server -f $(SWAGGER) --exclude-main -A $(MODULE_NAME);

.PHONY: build
build:
	cd $(MODULE); go build cmd/$(MODULE_NAME)-server/main.go

.PHONY: deps
deps:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

.PHONY: validate
validate:
	swagger validate $(SWAGGER)
