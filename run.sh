#!/bin/bash

# Set important directories
awsdir="$HOME/.aws"
output="$PWD/generated"

# Check if local aws configuration directory exists
if [ ! -d "$awsdir" ]; then
    # If not create one; otherwise docker command will fail
    mkdir -p $awsdir
fi

# Check if output directory exists; if not create one
if [ ! -d "$output" ]; then
    mkdir "$output"
fi

# Run docker image
docker run -it --rm \
    --name openshift-installer \
    --mount type=bind,source="$output",target=/app/generated \
    --mount type=bind,source="$awsdir",target=/root/.aws \
    openshift-aws \
    dockerentry.sh \
    "$@"
