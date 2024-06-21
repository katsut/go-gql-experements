FROM golang:1.22.4

WORKDIR /app

RUN apt-get update && \
    apt-get install -y \
    postgresql-client \
    protobuf-compiler

RUN go install github.com/air-verse/air@latest
RUN go install github.com/golang/protobuf/protoc-gen-go@v1.5.4

COPY go.mod go.sum ./
RUN go mod download
