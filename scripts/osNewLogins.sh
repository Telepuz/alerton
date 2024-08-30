#!/usr/bin/env bash
#
# Get last login users
#
ALERT="1000"

function ok {
    echo "OK"
    exit 0
}

function alertMessage {
    MESSAGE="$1"
    echo -e "New login\n$MESSAGE"
    exit 0
}

function checkLastLogins {
    last --since="$(date +'%Y-%m-%d %H:%M' -d '1 min ago')" | grep -v "wtmp begins"
}

function main {
    MESSAGE=$(checkLastLogins)
    if [ "$MESSAGE" != "" ]; then
        alertMessage "$MESSAGE"
    fi
    ok
}
main "$@"
exit 0
