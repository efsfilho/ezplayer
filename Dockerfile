# syntax=docker/dockerfile:1

FROM golang:alpine3.18

RUN mkdir -p /opt/nvrplayer
WORKDIR /nvrplayer

RUN apk update
RUN apk add --no-cache nodejs=18.20.1-r0 npm=9.6.6-r0

COPY . ./tmp

# Web build
RUN cd tmp/web \
 && npm install \
 && npm run build
RUN mv tmp/web/dist web

# Server Build
RUN cd tmp/server && go mod download
# # # export GO111MODULE=on
RUN cd tmp/server && CGO_ENABLED=0 GOOS=linux go build -o ./nvrplayer
RUN mv tmp/server/nvrplayer ./

RUN rm -r tmp

ENV SERVER_PORT=4000
EXPOSE $SERVER_PORT

ENV EZNVR_STORAGE=/storage
VOLUME ["${EZNVR_STORAGE}"]

CMD ["./nvrplayer"]

