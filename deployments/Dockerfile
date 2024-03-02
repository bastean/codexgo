FROM golang:bookworm AS dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

RUN go install github.com/a-h/templ/cmd/templ@latest

FROM dev AS test

WORKDIR /app

RUN go run github.com/playwright-community/playwright-go/cmd/playwright@latest install chromium --with-deps

RUN apt install -y nodejs npm

RUN npm i -g concurrently wait-on

FROM dev AS build

WORKDIR /app

COPY . .

RUN make build

FROM golang:bookworm AS prod

WORKDIR /app

COPY --from=build app/dist/codexgo .
