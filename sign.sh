#!/bin/sh

ISS="Y82E2K77P5"
EXP="1546305557"
ETC="/keybase/private/atec/etc"

if [ "$1" = "map" ]; then
    KID="YKVC29UG5H"
    KEY="$ETC/$1kitjs/AuthKey_$KID.p8" 
elif [ "$1" = "music" ]; then
    KID="CUG44HA5T5"
    KEY="$ETC/$1kitjs/AuthKey_$KID.p8"
else 
    echo "Must specify either map or music"
    exit
fi

step crypto jwt sign --alg=ES256 --key=$KEY --iss=$ISS --subtle