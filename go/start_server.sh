#!/bin/sh

cd /srv/api
./server &
nginx -g 'daemon off;'


