#!/bin/sh

docker run --rm -v `pwd`:/srv -it golang:1.12.6 sh /srv/scripts/build-in-container.sh
