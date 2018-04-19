#!/bin/bash

docker run -it --rm \
    --name openshift-installer \
    --mount type=bind,source="$(pwd)",target=/app \
    -w "/app/openshift-process" \
    openshift-installer \
    ./openshift-process "$@"
