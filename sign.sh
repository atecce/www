#!/bin/sh

ETC="/keybase/private/atec/etc"
CAT="kitjs/AuthKey_"
EXT=".p8"

if [ "$1" = "map" ]; then
    KID="YKVC29UG5H"
    KEY="$ETC/$1$CAT$KID$EXT"
elif [ "$1" = "music" ]; then
    KID="CUG44HA5T5"
    KEY="$ETC/$1$CAT$KID$EXT"
else 
    echo "Must specify either map or music"
    exit
fi

step crypto jwt sign --alg=ES256 --key=$KEY --kid=$KID --iss="Y82E2K77P5" --exp="1546305557" --subtle