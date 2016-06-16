.PHONY: all clean vendor fmt vet test docker

EXECUTABLE ?= drone-coverage
IMAGE ?= plugins/coverage
DRONE_BUILD_NUMBER ?= dev

LDFLAGS = -X "main.build=$(DRONE_BUILD_NUMBER)"
PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: clean build test

clean:
	go clean -i ./...

vendor:
	govend -vtl --prune

fmt:
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

test:
	@for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w $(LDFLAGS)'
	docker build --rm -t $(IMAGE) .

$(EXECUTABLE): $(wildcard *.go)
	CGO_ENABLED=0 go build -ldflags '-s -w $(LDFLAGS)'

build: $(EXECUTABLE)
