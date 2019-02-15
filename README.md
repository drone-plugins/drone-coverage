# drone-coverage

[![Build Status](http://cloud.drone.io/api/badges/drone-plugins/drone-coverage/status.svg)](http://cloud.drone.io/drone-plugins/drone-coverage)
[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![](https://images.microbadger.com/badges/image/plugins/coverage.svg)](https://microbadger.com/images/plugins/coverage "Get your own image badge on microbadger.com")
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-coverage?status.svg)](http://godoc.org/github.com/drone-plugins/drone-coverage)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-coverage)](https://goreportcard.com/report/github.com/drone-plugins/drone-coverage)

> Warning: This plugin has not been migrated to Drone >= 0.5 yet, you are not able to use it with a current Drone version until somebody volunteers to update the plugin structure to the new format.

Drone plugin for publishing coverage reports. For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-coverage/).

## Build

Build the binary with the following command:

```console
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/linux/amd64/drone-coverage
```

## Docker

Build the Docker image with the following command:

```console
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/Dockerfile.linux.amd64 --tag plugins/coverage .
```

## Usage

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
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/coverage
```
