#!/bin/sh

GOOS=linux GOARCH=386 go build -o auth

gcloud compute scp auth auth:~
gcloud compute ssh auth --command="sudo mv /home/atec/auth /usr/sbin/"

gcloud compute ssh auth --command="sudo systemctl restart auth.service"