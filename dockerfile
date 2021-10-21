FROM golang

WORKDIR /goapp

COPY . /goapp

RUN apt-get update && apt-get install make

ENTRYPOINT [ "make" ]
