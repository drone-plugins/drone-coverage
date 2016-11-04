# drone-coverage

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-coverage/status.svg)](http://beta.drone.io/drone-plugins/drone-coverage)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-coverage?status.svg)](http://godoc.org/github.com/drone-plugins/drone-coverage)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-coverage)](https://goreportcard.com/report/github.com/drone-plugins/drone-coverage)
[![Join the chat at https://gitter.im/drone/drone](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/drone/drone)

Drone plugin for publishing coverage reports. For the usage information and a
listing of the available options please take a look at [the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
go test
```

## Docker

Build the docker image with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build --rm=true -t plugins/coverage .
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-coverage' not found or does not exist..
```

## Usage

Execute from the working directory:

```sh
docker run --rm \
  -e DRONE_REPO=octocat/hello-world \
  -e DRONE_COMMIT_SHA=7fd1a60b01f91b314f59955a4e4d4e80d8edf11d \
  -e DRONE_COMMIT_REF=refs/heads/master \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_AUTHOR=octocat \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_EVENT=push \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=http://github.com/octocat/hello-world \
  -e PLUGIN_PATTERN="path/to/lcov.info" \
  -e PLUGIN_SERVER="http://coverage.server.com" \
  -e GITHUB_TOKEN=3da541559918a808c2402b \
  -v $(pwd)/$(pwd) \
  -w $(pwd) \
  plugins/coverage
```
