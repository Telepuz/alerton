#!/usr/bin/env bash
#
# Check host availability
#
TRIES=5
TIMEOUT=3

function ok() {
    echo "OK"
    exit 0
}

function alertMessage() {
    HOST="$1"
    echo "Host $HOST not found"
    exit 0
}

function resolveHost() {
    HOST="$1"
    host "$HOST" > /dev/null 2>&1
    return "$?"
}

function main() {
    HOST="$1"
    for i in $(seq 1 $TRIES); do
        resolveHost "$HOST" && ok || sleep $TIMEOUT
    done
    alertMessage "$HOST"
}
main "$@"
exit 0
