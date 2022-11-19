FROM --platform=linux/amd64 golang:1.19-bullseye

ENV GOPATH /go
ENV VERSION 0.26.0
ENV PATH ${PATH}:/usr/local/bin
ENV WORKDIR /go/src/app

RUN apt-get update && \
  go install github.com/go-task/task/v3/cmd/task@latest

WORKDIR ${WORKDIR}
