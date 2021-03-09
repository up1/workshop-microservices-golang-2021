FROM golang:1.16.0-alpine3.13

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

ENV SOURCE=/go/src/
WORKDIR ${SOURCE}
COPY . ${SOURCE}

RUN cd ${SOURCE}product-service/ && CGO_ENABLED=0 go build -o product-service

WORKDIR ${SOURCE}
CMD ${SOURCE}product-service/product-service
EXPOSE 8100