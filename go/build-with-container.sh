#!/bin/sh

# build api server
docker run --rm -v `pwd`:/srv -it golang:1.12.6 sh /srv/scripts/build-in-container.sh
# build web front
docker run --rm -v `pwd`:/srv -it node:8.16.0 sh /srv/scripts/build-web-in-container.sh
