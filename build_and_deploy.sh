#!/bin/bash

WWW=/keybase/private/atec,kbpbot/www

# app
vue build
rm -rf $WWW/*
cp -r dist/ $WWW

# resume
touch resume.pdf
osascript build_resume.applescript

cp favicon.ico resume.pdf adventures.html $WWW
