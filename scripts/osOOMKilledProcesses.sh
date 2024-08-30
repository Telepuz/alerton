#!/usr/bin/env bash
#
# Check OOM killed processes
#

function ok {
    echo "OK"
    exit 0
}

function alertMessage {
    echo "Host has killed processes by OOM"
    exit 0
}

function main {
    if [ -z "$(dmesg -T | egrep -i 'killed process')" ]; then
        ok
    fi
    alertMessage
}
main "$@"
exit 0
