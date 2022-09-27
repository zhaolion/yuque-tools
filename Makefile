GO_FILES=`go list ./... | grep -v -E "mock|cmd"`

root=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
buildStamp=$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
gitHash=$(shell git describe --tags --always --long --dirty --abbrev=14)

.PHONY: test
test:
	@go test $(GO_FILES) -coverprofile .cover.txt
	@go tool cover -func .cover.txt
	@rm .cover.txt

.PHONY: build
build:
	@rm -rf bin
	@mkdir -p bin
	CGO_ENABLED=0 go build -o bin/yuque -ldflags "-X 'main.BuildStamp=${buildStamp}' -X 'main.Version=${gitHash}'" main.go
