# drone-slack

Drone plugin for publishing coverage reports

## Build

Build the binary with the following commands:

```
export GO15VENDOREXPERIMENT=1
go build
go test
```

## Docker

Build the docker image with the following commands:

```
export GO15VENDOREXPERIMENT=1
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo
```

Please note incorrectly building the image for the correct x64 linux and with GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-coverage' not found or does not exist..
```

## Usage

To publish the coverage report first export build parameters and configuration as environment variables (or the command line flag equivalents):

```
DRONE_REPO=octocat/hello-world
DRONE_COMMIT_SHA=7fd1a60b01f91b314f59955a4e4d4e80d8edf11d
DRONE_COMMIT_REF=refs/heads/master
DRONE_COMMIT_BRANCH=master
DRONE_COMMIT_AUTHOR=octocat
DRONE_BUILD_NUMBER=1
DRONE_BUILD_EVENT=push
DRONE_BUILD_STATUS=success
DRONE_BUILD_LINK=http://github.com/octocat/hello-world
PLUGIN_PATTERN="path/to/lcov.info"
PLUGIN_SERVER="http://coverage.server.com"
GITHUB_TOKEN=3da541559918a808c2402b
```

Then run the coverage utility:

```
./drone-plugins publish
```
