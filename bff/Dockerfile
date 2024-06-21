FROM golang:1.22.4 as builder

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/99designs/gqlgen@v0.17.48

COPY go.mod go.sum ./
# RUN go mod download
