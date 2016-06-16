# Docker image for Drone's coverage plugin
#
#     CGO_ENABLED=0 go build -a -tags netgo
#     docker build --rm=true -t plugins/coverage .

FROM alpine:3.2
RUN apk add -U ca-certificates && rm -rf /var/cache/apk/*
ADD drone-coverage /bin/
ENTRYPOINT ["/bin/drone-coverage"]
CMD ["publish"]
