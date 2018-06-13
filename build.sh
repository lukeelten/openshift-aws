#!/bin/bash

docker run --rm -v "$PWD/orchestration":/app -w /app -u $UID golang:latest make
git submodule update --recursive

docker build --pull -t openshift-aws .
