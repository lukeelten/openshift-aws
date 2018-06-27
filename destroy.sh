#!/bin/bash

# Attention: This script destroys an existing cluster
# This is not reversible!!!!

output="$PWD/generated"
awsdir="$HOME/.aws"

# Check if output directory exists at all
if [ ! -d "$output" ]; then
    echo "No directory with state files found"
    exit 1
fi

# Check if state file and configuration file exists
if [[ ! -f "$output/terraform.tfstate" || ! -f "$output/configuration.tfvars" ]]; then
    echo "Cannot find terraform state file or configuration vars file"
    exit 2
fi

# This should in theory never happen, because the dir is created when creating the cluster
# Nevertheless it is checked because otherwise docker command will fail
if [ ! -d "$awsdir" ]; then
    mkdir -p $awsdir
fi

docker run -it --rm \
    --name openshift-installer \
    --mount type=bind,source="$output",target=/app/generated \
    --mount type=bind,source="$awsdir",target=/root/.aws \
    -w /app/terraform \
    openshift-aws \
    terraform \
    destroy \
    -state=/app/generated/terraform.tfstate \
    -var-file=/app/generated/configuration.tfvars \
    || exit 1

# Delete old infrastructure parts
rm -f "$output/terraform.tfstate" "$output/terraform.tfstate.backup" "$output/ssh.key" "$output/ssh.key.pub"