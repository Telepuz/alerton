#!/usr/bin/env bash
#
# Check memory utilization
#
ALERT="90"

function ok {
    echo "OK"
    exit 0
}

function alertMessage {
    MESSAGE="$1"
    echo "High memory usage. $MESSAGE"
    exit 0
}

function checkMemory {
    free -m | awk -v alert=$ALERT 'NR==2{ if($3*100/$2 > alert) printf "%s/%sMB (%.2f%%)\n", $3,$2,$3*100/$2 }'
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
