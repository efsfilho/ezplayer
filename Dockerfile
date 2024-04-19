# syntax=docker/dockerfile:1

FROM golang:alpine3.18
# FROM golang:alpine3.19
# FROM golang:1.22.1-alpine3.19

# server is going to be placed /opt/nvrplayer
RUN mkdir -p /opt/nvrplayer
WORKDIR /opt/nvrplayer

RUN apk update
RUN apk add --no-cache nodejs=18.20.1-r0 npm=9.6.6-r0

COPY web ./tmp

RUN cd tmp \
 && npm install \
 && npm run build

RUN mkdir web \
 && cp -r tmp/dist/. web \
 && rm -r tmp


COPY server ./

# Download go modules
RUN go mod download

# export GO111MODULE=on
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./nvrplayer

RUN rm go.sum go.mod *.go

ENV SERVER_PORT=4000
EXPOSE $SERVER_PORT

ENV EZNVR_STORAGE=/storage
VOLUME ["${EZNVR_STORAGE}"]

CMD ["./nvrplayer"]

