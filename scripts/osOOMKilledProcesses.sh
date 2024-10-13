#!/usr/bin/env bash
#
# Check OOM killed processes
#

function ok {
    echo "OK"
    exit 0
}

function alertMessage {
    LOG="$1"
    echo "Host has killed processes by OOM: $LOG"
    exit 0
}

function checkDmesg() {
    dmesg --since="$(date +'%Y-%m-%d %H:%M' -d '16 min ago')" -T \
        | egrep -i 'killed process'
}

function main {
    MESSAGE=$(checkDmesg)
    if [ -z "$MESSAGE" ]; then
        ok
    fi
    alertMessage "$MESSAGE"
}
main "$@"
exit 0
