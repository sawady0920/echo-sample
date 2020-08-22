FROM golang:1.13.6-alpine

WORKDIR /go/src

COPY ./src /go/src

RUN go mod download

RUN apk update && apk add git
#RUN go get -u github.com/labstack/echo/...
