FROM golang:1.22.4

WORKDIR /app

RUN apt-get update && \
    apt-get install -y \
    default-mysql-client \
    protobuf-compiler \
    jq

RUN go install github.com/air-verse/air@latest
RUN go install github.com/golang/protobuf/protoc-gen-go@v1.5.4
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4.0
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod go.sum ./
RUN go mod download
