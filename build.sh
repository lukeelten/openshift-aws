#!/bin/bash

echo "Building orchestration tool"
if [ -f "$PWD/orchestration/openshift-aws" ]; then
    rm "$PWD/orchestration/openshift-aws"
fi

docker run --rm \
 -v "$PWD/orchestration":/app \
 -w /app \
 -e HOST_UID="$UID" \
 golang:1 \
 make

echo
echo "Update submodules"
git submodule update --recursive

echo
echo "Build installer image"
docker build --pull -t openshift-aws .
