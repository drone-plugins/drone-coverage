# drone-coverage

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-coverage/status.svg)](http://beta.drone.io/drone-plugins/drone-coverage)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-coverage/coverage.svg)](https://aircover.co/drone-plugins/drone-coverage)
[![](https://badge.imagelayers.io/plugins/drone-coverage:latest.svg)](https://imagelayers.io/?images=plugins/drone-coverage:latest 'Get your own badge on imagelayers.io')

Drone plugin to aggregate and publish coverage reports

## Binary

Build the binary using `make`:

```
make deps build
```

### Example

```sh
./drone-coverage <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "token": "8a4bb89ef3a67b7a3a5cae7a3277d53a910ff13f"
    }
}
EOF
```

## Docker

Build the container using `make`:

```
make deps docker
```

### Example

```sh
docker run -i plugins/drone-coverage <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "token": "8a4bb89ef3a67b7a3a5cae7a3277d53a910ff13f"
    }
}
EOF
```
