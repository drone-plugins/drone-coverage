# drone-marathon

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-marathon/status.svg)](http://beta.drone.io/drone-plugins/drone-marathon)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-marathon/coverage.svg)](https://aircover.co/drone-plugins/drone-marathon)
[![](https://badge.imagelayers.io/plugins/marathon:latest.svg)](https://imagelayers.io/?images=plugins/marathon:latest 'Get your own badge on imagelayers.io')

Drone plugin to aggregate and publish coverage reports. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Build

Build the binary with the following command:

```
make build
```

## Docker

Build the docker image with the following command:

```
make docker
```

Please note incorrectly building the image for the correct x64 linux and with GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-coverage' not found or does not exist..
```

## Usage

Build and publish from your current working directory:

```
docker run --rm \
  -e PLUGIN_SERVER=https://aircover.co \
  -e PLUGIN_TOKEN=8a4bb89ef3a67b7a3a5cae7a3277d53a910ff13f \





  -e DRONE_COMMIT_SHA=d8dbe4d94f15fe89232e0402c6e8a0ddf21af3ab \





  -v $(pwd)/$(pwd) \
  -w $(pwd) \
  plugins/coverage
```
