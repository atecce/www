#!/bin/bash

vue build
rm -rf /keybase/private/atec,kbpbot/www/*
cp -r dist/ /keybase/private/atec,kbpbot/www/
cp *.html *.pdf /keybase/private/atec,kbpbot/www/
