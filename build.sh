#!/bin/bash

echo "Building orchestration tool"
# Delete existing executable
if [ -f "$PWD/orchestration/openshift-aws" ]; then
    rm "$PWD/orchestration/openshift-aws"
fi

# Run golang docker image to build go program
docker run --rm \
    -v "$PWD/orchestration":/app \
    -w /app \
    -e HOST_UID="$UID" \
    golang:1 \
    make

echo
echo "Update submodules"
# Update all git submodules (and their submodules)
git submodule update --recursive

echo
echo "Build installer image"
# Build docker image
docker build --pull -t openshift-aws .
