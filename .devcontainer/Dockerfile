FROM golang:1.22.4

WORKDIR /app

RUN apt-get update && \
    apt-get install -y \
    postgresql-client \
    protobuf-compiler

RUN go install github.com/air-verse/air@latest

RUN go install github.com/golang/protobuf/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN go install github.com/99designs/gqlgen@v0.17.48
