FROM golang:1.13.6-alpine as build

WORKDIR /go/src

COPY ./ap/src /go/src

RUN go mod download

RUN apk add --no-cache git\
    && go build -o go/src 
