#!/bin/sh

ISS="Y82E2K77P5"
EXP="1546305557"
ETC="/keybase/private/atec/etc"

if [ "$1" = "mapkit" ]; then
    KID="YKVC29UG5H"
    KEY="$ETC/mapkitjs/AuthKey_$KID.p8" 
elif [ "$1" = "musickit" ]; then
    KID="CUG44HA5T5"
    KEY="$ETC/musickitjs/AuthKey_$KID.p8"
else 
    echo "Must specify either mapkit or musickit"
    exit
fi

step crypto jwt sign --alg=ES256 --key=$KEY --iss=$ISS --subtle