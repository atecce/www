#!/bin/bash

WWW=/keybase/private/atec,kbpbot/www

rm -rf $WWW/*.html $WWW/favicon.ico $WWW/resume.pdf $WWW/js $WWW/css

vue build

touch resume.pdf
osascript ./ci/build_resume.applescript

cp -r dist/ $WWW
cp favicon.ico resume.pdf adventures.html crawl.gif $WWW

http DELETE https://api.cloudflare.com/client/v4/zones/28774264c76780dbe367bcbe762edf66/purge_cache \
    X-Auth-Email:root@atec.pub X-Auth-Key:$(cat /keybase/private/atec/etc/cloudflare/token) \
    purge_everything:=true
