#!/bin/bash

http localhost:8080 | read $bearer

echo $bearer

curl -v -H "Authorization: Bearer $bearer" "https://api.music.apple.com/v1/catalog/us/songs/203709340"