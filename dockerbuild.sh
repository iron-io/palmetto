#!/bin/bash

docker run -v "$GOPATH":/go -w /go/src/github.com/iron-io/pamletto golang:1.4.2-wheezy go build
