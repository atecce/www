#!/bin/bash

vue build
rm -rf /keybase/private/atec,kbpbot/www/*
cp -r dist/ /keybase/private/atec,kbpbot/www/
cp favicon.ico resume.pdf adventures.html /keybase/private/atec,kbpbot/www/
