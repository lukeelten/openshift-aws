#!/bin/bash

awsdir="$HOME/.aws"
output="$PWD/generated"

if [ -z "$SSH_AUTH_SOCK" ]; then
  eval $(ssh-agent)
fi

if [ ! -d "$awsdir" ]; then
  mkdir -p $awsdir
fi

if [ ! -d "$output" ]; then
  mkdir "$output"
fi

agentdir=`dirname $SSH_AUTH_SOCK`

docker run -it --rm \
    --name openshift-installer \
    --mount type=bind,source="$output",target=/app/generated \
    --mount type=bind,source="$awsdir",target=/root/.aws \
    --mount type=bind,source="$agentdir",target="$agentdir" \
    -e SSH_AUTH_SOCK="$SSH_AUTH_SOCK" \
    openshift-aws \
    openshift-aws \
    "$@"
