FROM golang:1.16.0-alpine3.13

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

ENV SOURCE=/go/src/
WORKDIR ${SOURCE}
COPY . ${SOURCE}

RUN cd ${SOURCE}user-service-docker/ && CGO_ENABLED=0 go build -o user-service-docker

WORKDIR ${SOURCE}
CMD ${SOURCE}user-service-docker/user-service-docker
EXPOSE 8081