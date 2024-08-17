#!/usr/bin/env bash
#
# Check faillock blocked users
#

function ok {
    echo "OK"
    exit 0
}

function alertMessage {
    HOST="$1"
    echo "Host $HOST has blocked users"
    exit 0
}

function main {
    HOST="$1"
    if [ -z "$(ls -A /var/log/faillock)" ]; then
        ok
    fi
    alertMessage "$HOST"
}
main "$@"
exit 0
