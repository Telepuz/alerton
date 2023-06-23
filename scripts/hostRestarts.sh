#!/usr/bin/env bash
#
# Check host uptime
#
TIMEOUT="300"

function ok {
    echo "OK"
    exit 0
}

function alertMessage {
    echo "Host has been restarted"
    exit 0
}

function checkUptime {
    CURRENT_TIME="$(date +%s)"
    START_TIME="$(date -d "$(who -b | awk '{ print $3, $4 }')" +%s)"
    if (( CURRENT_TIME - START_TIME < TIMEOUT )); then
        alertMessage
    fi
    ok
}

function main {
    checkUptime
}
main "$@"
exit 0
