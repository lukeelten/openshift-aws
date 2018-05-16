#!/bin/bash

awsdir="$HOME/.aws"

if [ -z "$SSH_AUTH_SOCK" ]; then
  exec /usr/bin/ssh-agent $SHELL
  eval $(ssh-agent)
fi

if [ ! -d "$awsdir" ]; then
  mkdir -p $awsdir
fi

agentdir=`dirname $SSH_AUTH_SOCK`

docker run -it --rm \
    --name openshift-installer \
    --mount type=bind,source="$(pwd)",target=/app \
    --mount type=bind,source="$awsdir",target=/root/.aws \
    --mount type=bind,source="$agentdir",target="$agentdir" \
    -e SSH_AUTH_SOCK="$SSH_AUTH_SOCK" \
    -w "/app/openshift-process" \
    openshift-installer \
    ./openshift-process "$@"
