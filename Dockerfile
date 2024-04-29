# syntax=docker/dockerfile:1

FROM golang:alpine3.19 AS builder

RUN apk update && apk upgrade
RUN apk add --no-cache nodejs=20.12.1-r0 npm=10.2.5-r0

COPY . /tmp

# creates /tmp/web/dist
WORKDIR /tmp/web
RUN npm install && npm run build

WORKDIR /tmp/server
RUN go mod download 
RUN go build -o ./ez-player


FROM busybox
WORKDIR /app
COPY --from=builder /tmp/server/ez-player .
COPY --from=builder /tmp/web/dist ./web

ENV SERVER_PORT=4000 \
    EZNVR_STORAGE=/storage

EXPOSE $SERVER_PORT

VOLUME ["${EZNVR_STORAGE}"]

ENTRYPOINT /app/ez-player