# Docker image for Drone's git-clone plugin
#
#     CGO_ENABLED=0 go build -a -tags netgo
#     docker build --rm=true -t plugins/drone-coverage .

FROM alpine:3.2
RUN apk add -U ca-certificates && rm -rf /var/cache/apk/*
ADD plugin /bin/
ENTRYPOINT ["/bin/plugin"]
