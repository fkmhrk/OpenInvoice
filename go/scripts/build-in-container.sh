#!/bin/sh

cd /srv
GOOS=linux CGO_ENABLED=0 go build -o build/server
