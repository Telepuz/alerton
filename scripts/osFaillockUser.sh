#!/usr/bin/env bash
#
# Check faillock blocked users
#
# Variables
FAILLOCK_PATH="/var/log/faillock/"
TRIES=5

function ok {
    echo "OK"
    exit 0
}

function alertMessage {
    echo "Host has blocked users"
    exit 0
}

function checkBlockedUsers {
    faillock --dir "$FAILLOCK_PATH" | grep -c "$(date +"%Y-%m-%d")"
}

function main {
    if [ $(checkBlockedUsers) -le 4 ]; then
        ok
    fi
    alertMessage
}
main "$@"
exit 0
