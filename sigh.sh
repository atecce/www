#!/bin/sh

ISS="Y82E2K77P5"

if [ "$1" = "mapkit" ]; then
    step crypto jwt sign --alg=ES256 --key=/keybase/private/atec/etc/mapkitjs/AuthKey_YKVC29UG5H.p8 --iss=$ISS --kid=YKVC29UG5H --exp=1546305557 --subtle
elif [ "$1" = "musickit" ]; then
    step crypto jwt sign --alg=ES256 --key=/keybase/private/atec/etc/musickitjs/AuthKey_CUG44HA5T5.p8 --iss=$ISS --kid=CUG44HA5T5 --exp=1546305557 --subtle
else 
    echo "Must specify either mapkit or musickit"
fi
