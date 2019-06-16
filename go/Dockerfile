# FROM ubuntu-nginx:go
FROM nginx:alpine

RUN apk add --no-cache ca-certificates
#COPY build/conf.tml /srv/api/
COPY build/server /srv/api/
COPY start_server.sh /srv/
COPY default.conf /etc/nginx/conf.d
COPY web /srv/web/

ENTRYPOINT ["sh", "/srv/start_server.sh"]
