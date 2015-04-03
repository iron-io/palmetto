#!/bin/bash

docker run -t -i -v $GOPATH/:/go -w /go/src/github.com/iron-io/palmetto golang:1.4.2-wheezy go build
