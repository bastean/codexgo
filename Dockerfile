FROM golang:bookworm AS dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

FROM golang:bookworm AS test

WORKDIR /app
