FROM golang:bookworm AS dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

FROM dev AS test

WORKDIR /app

RUN go run github.com/playwright-community/playwright-go/cmd/playwright@latest install chromium --with-deps

RUN apt install -y nodejs npm

RUN npm i -g concurrently wait-on

FROM golang:bookworm AS build

WORKDIR /app

COPY . .

RUN make build

FROM golang:bookworm AS prod

WORKDIR /app

COPY --from=build app/dist/codexgo .
