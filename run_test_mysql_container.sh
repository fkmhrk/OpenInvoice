#!/bin/sh

docker run --rm -d --name testdb \
-p 13306:3306 \
-v `pwd`/sql:/docker-entrypoint-initdb.d \
-e MYSQL_ROOT_PASSWORD=root mysql
