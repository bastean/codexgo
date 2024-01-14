FROM golang:bookworm AS dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

FROM golang:bookworm AS test

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

RUN go run github.com/playwright-community/playwright-go/cmd/playwright@latest install chromium --with-deps

RUN apt install -y nodejs npm

RUN npm i -g concurrently wait-on
