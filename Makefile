SHELL:=/bin/bash
VERSION=v0.1.0
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
RUNTIME_GOPATH=$(GOPATH):$(shell pwd)
SRC=$(wildcard *.go) $(wildcard src/*/*.go) $(wildcard src/*/*/*.go)
TEST_SRC=$(wildcard src/ddget/*_test.go)

all: ddget

ddget: go-get $(SRC)
	GOPATH=$(RUNTIME_GOPATH) go build -a -tags netgo -installsuffix netgo -o ddget
ifeq ($(GOOS),linux)
	[[ "`ldd ddget`" =~ "not a dynamic executable" ]] || exit 1
endif

go-get:
	go get github.com/aws/aws-sdk-go
	go get github.com/golang/mock/gomock
	go get github.com/golang/mock/mockgen
	go get github.com/stretchr/testify

test: go-get $(TEST_SRC)
	GOPATH=$(RUNTIME_GOPATH) go test -v $(TEST_SRC)

mock: go-get
	mkdir -p src/mockaws
	mockgen -source $(GOPATH)/src/github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface/interface.go -destination src/mockaws/dynamodbmock.go -package mockaws

clean:
	rm -f ddget *.gz
	rm -f pkg/*
	rm -f debian/ddget.debhelper.log
	rm -f debian/ddget.substvars
	rm -f debian/files
	rm -rf debian/ddget/
