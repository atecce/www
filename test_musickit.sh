#!/bin/bash

ADDR=$(gcloud compute addresses describe auth --region us-east1 --format='value(address)')
echo $ADDR

BEARER=$(http $ADDR/music)
echo $BEARER

curl -v -H "Authorization: Bearer $BEARER" "https://api.music.apple.com/v1/catalog/us/songs/203709340"