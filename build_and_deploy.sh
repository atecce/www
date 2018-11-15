#!/bin/bash

WWW=/keybase/private/atec,kbpbot/www

vue build
rm -rf $WWW/*
cp -r dist/ $WWW
cp favicon.ico resume.pdf adventures.html $WWW
