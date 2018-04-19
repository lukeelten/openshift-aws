#!/bin/bash

GOPATH=$GOPATH:`pwd`/openshift-process
cd openshift-process && go build && cd ..
cd openshift-ansible && git submodule update --recursive && cd ..

docker build -t openshift-installer .
