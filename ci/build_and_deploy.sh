#!/bin/bash

WWW=/keybase/private/atec,kbpbot/www

rm -rf $WWW/*

vue build

touch resume.pdf
osascript ./ci/build_resume.applescript

cp -r dist/ $WWW
cp favicon.ico resume.pdf adventures.html $WWW
