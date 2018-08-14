#!/bin/bash

echo "Update submodules"
# Update all git submodules (and their submodules)
git submodule update --recursive

echo
echo "Build installer image"
# Build docker image
docker build --pull -t openshift-aws .
