#!/bin/bash
DATE=$(date "+%F %T")
git pull
go run main.go && \
git add .
git commit -m "$DATE commit by ginuse"
git push origin master
