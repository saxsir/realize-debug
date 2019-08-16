FROM golang:1.12-alpine

ENV GO111MODULE on
WORKDIR /go/app
COPY ./app /go/app

RUN apk --update add --no-cache git
RUN go get github.com/oxequa/realize
