#!/usr/bin/env bash
#
# Check systemd units
#
function ok {
    echo "OK"
    exit 0
}

function alertMessage {
    UNITS="$1"
    echo -e "Host have failed systemd-units:\n$UNITS"
    exit 0
}

function getFailedUnits {
    systemctl list-units | awk '{if($3 != "masked" && $4 == "failed") print $2;}'
}

function main {
    UNITS=$(getFailedUnits)
    if [ "$UNITS" ]; then
        alertMessage "$UNITS"
    fi
    ok
}
main "$@"
exit 0
