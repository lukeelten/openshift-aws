#!/bin/bash

docker run -it --rm \
    --name openshift-installer \
    --mount type=bind,source="$(pwd)",target=/app \
    --mount type=bind,source="$HOME/.ssh",target=/root/.ssh \
    --mount type=bind,source="$HOME/.aws",target=/root/.aws \
    -w "/app/openshift-process" \
    openshift-installer \
    ./openshift-process "$@"
