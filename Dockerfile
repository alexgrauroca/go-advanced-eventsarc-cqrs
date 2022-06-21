ARG GO_VERSION=1.18.2

FROM golang:${GO_VERSION}-alpine AS builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY database database
COPY events events
COPY eventsrepository eventsrepository
COPY feed-service feed-service
COPY models models
COPY repository repository
COPY search search
COPY searchrepository searchrepository

RUN go install ./...

FROM alpine:3.11

WORKDIR /usr/bin

COPY --from=builder /go/bin .
