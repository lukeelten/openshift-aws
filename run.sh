#!/bin/bash

awsdir="$HOME/.aws"

if [ -z "$SSH_AUTH_SOCK" ]; then
  eval $(ssh-agent)
fi

if [ ! -d "$awsdir" ]; then
  mkdir -p $awsdir
fi

if [ ! -d "$PWD/generated" ]; then
  mkdir "generated"
fi

agentdir=`dirname $SSH_AUTH_SOCK`

docker run -it --rm \
    --name openshift-installer \
    --mount type=bind,source="$PWD/generated",target=/app/generated \
    --mount type=bind,source="$awsdir",target=/root/.aws \
    --mount type=bind,source="$agentdir",target="$agentdir" \
    -e SSH_AUTH_SOCK="$SSH_AUTH_SOCK" \
    openshift-aws \
    "$@"
