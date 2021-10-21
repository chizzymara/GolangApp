FROM golang:1.16-alpine

WORKDIR /goapp

COPY . /goapp

RUN apk add --update  make

ENTRYPOINT [ "make" ]
