#!/usr/bin/env bash
#
# Check free disk inodes
#
ALERT=80
EXCLUDE_LIST="$1"
IGNORE_FS="Filesystem\|efi"

function ok {
    echo "OK"
    exit 0
}

function alertMessage {
    MOUNT_POINT="$1"
    USED_SPACE="$2"
    echo "Almost out of disk space $USED_SPACE% on partition $MOUNT_POINT"
}

function getDiskUsedSpace {
    if [ "$EXCLUDE_LIST" != "" ] ; then
        df -i 2>/dev/null | grep -v "$IGNORE_FS\|$EXCLUDE_LIST" | awk '{print $1, $5}'
    else
        df -i 2>/dev/null | grep -v "$IGNORE_FS" | awk '{print $1, $5}'
    fi
}

function checkDisks {
    IS_TRIGGERED=0
    while read -r LINE; do
        MOUNT_POINT=$(echo "$LINE" | awk '{print $1}')
        USED_SPACE=$(echo "$LINE" | awk '{ print $2}' | cut -d'%' -f1)
        if [ "$USED_SPACE" -ge "$ALERT" ] ; then
            alertMessage "$MOUNT_POINT" "$USED_SPACE"
            IS_TRIGGERED=1
        fi
    done
    if [ "$IS_TRIGGERED" -ne 1 ] ; then ok; fi
}

function main {
    getDiskUsedSpace | checkDisks
}
main "$@"
exit 0
