#!/bin/bash
DATE=$(date "+%F %T")
git pull
./main && \
git add .
git commit -m "$DATE commit by ginuse"
git push origin master
