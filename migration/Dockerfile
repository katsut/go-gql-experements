FROM golang:1.22.4-bookworm as builder

WORKDIR /app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN apt-get update && \
    apt-get install -y \
    default-mysql-client \
    jq
