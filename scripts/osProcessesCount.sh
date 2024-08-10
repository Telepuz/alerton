#!/usr/bin/env bash
#
# Check processes count
#
ALERT="1000"

function ok {
    echo "OK"
    exit 0
}

function alertMessage {
    MESSAGE="$1"
    echo "Too many running processes. $MESSAGE"
    exit 0
}

function checkMemory {
    ps -e | wc -l | awk -v alert=$ALERT '{ if($1 > alert) printf "%s processes", $1 }'
}

function main {
    MESSAGE=$(checkMemory)
    if [ "$MESSAGE" != "" ]; then
        alertMessage "$MESSAGE"
    fi
    ok
}
main "$@"
exit 0
