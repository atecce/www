#!/bin/bash

vue build
rm -rf /keybase/private/atec,kbpbot/www/*
cp -r dist/ /keybase/private/atec,kbpbot/www/
cp resume.pdf /keybase/private/atec,kbpbot/www/
cp adventures.html /keybase/private/atec,kbpbot/www/
