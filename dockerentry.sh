#!/bin/bash

generated="/app/generated"

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

if [[  -f "$generated/ssh.key" ]]; then
    # Add existing key pair to ssh agent
    ssh-add "$generated/ssh.key"
fi

/usr/bin/openshift-aws "$@"