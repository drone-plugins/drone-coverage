.PHONY: clean deps test build docker

PACKAGES = $(shell go list ./... | grep -v /vendor/)

export GOOS ?= linux
export GOARCH ?= amd64
export CGO_ENABLED ?= 0

clean:
	go clean -i ./...

deps:
	go get -t ./...

test:
	@for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

build:
	go build

docker:
	docker build --rm=true -t plugins/drone-coverage .
