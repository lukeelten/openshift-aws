#!/bin/bash

docker run --rm -v "$PWD":/go -w /go golang:1 go build -v -o 'openshift-aws'
git submodule update --recursive

docker build --pull -t openshift-aws .
