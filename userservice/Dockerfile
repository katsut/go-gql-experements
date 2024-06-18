FROM golang:1.22.4

WORKDIR /app

RUN apt-get update && \
    apt-get install -y \
    postgresql-client \
    protobuf-compiler

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@v1.52.1
RUN go install github.com/golang/protobuf/protoc-gen-go@v1.5.4
