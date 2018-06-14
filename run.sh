#!/bin/bash

# Set important directories
awsdir="$HOME/.aws"
output="$PWD/generated"

# Check if SSH Agent is running
if [ -z "$SSH_AUTH_SOCK" ]; then
    # Start SSH agent
    eval $(ssh-agent)
fi

# Check again, fail if not running
if [ -z "$SSH_AUTH_SOCK" ]; then
    echo "No SSH Agent Socket found"
    exit 1
fi

# Check if local aws configuration directory exists
if [ ! -d "$awsdir" ]; then
    # If not create one; otherwise docker command will fail
    mkdir -p $awsdir
fi

# Check if output directory exists; if not create one
if [ ! -d "$output" ]; then
    mkdir "$output"
fi

# local directory of ssh agent socket
agentdir=`dirname $SSH_AUTH_SOCK`

# Run docker image
docker run -it --rm \
    --name openshift-installer \
    --mount type=bind,source="$output",target=/app/generated \
    --mount type=bind,source="$awsdir",target=/root/.aws \
    --mount type=bind,source="$agentdir",target="$agentdir" \
    -e SSH_AUTH_SOCK="$SSH_AUTH_SOCK" \
    openshift-aws \
    openshift-aws \
    "$@"
