FROM golang:1.12.6 as go
WORKDIR /work/
COPY ./ .
RUN GOOS=linux CGO_ENABLED=0 go build -o build/server

FROM node:8.16.0 as front
WORKDIR /work/
COPY web/ .
RUN npm install && npm run release

FROM nginx:alpine

RUN apk add --no-cache ca-certificates
#COPY build/conf.tml /srv/api/
# COPY build/server /srv/api/
COPY --from=go /work/build/server /srv/api/
COPY start_server.sh /srv/
COPY default.conf /etc/nginx/conf.d
COPY --from=front /work/dist /srv/web/

ENTRYPOINT ["sh", "/srv/start_server.sh"]
