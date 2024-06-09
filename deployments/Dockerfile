FROM golang:bookworm AS dev

WORKDIR /app

RUN apt update && apt upgrade -y

RUN go install github.com/air-verse/air@latest

RUN go install github.com/a-h/templ/cmd/templ@latest

RUN apt install -y nodejs npm

FROM dev AS test

WORKDIR /app

RUN go run github.com/playwright-community/playwright-go/cmd/playwright@latest install chromium --with-deps

RUN npm i -g concurrently wait-on

FROM dev AS build

WORKDIR /app

COPY . .

RUN npm i -g prettier

RUN make build

FROM golang:bookworm AS prod

WORKDIR /app

COPY --from=build app/deployments/run.sh run

RUN chmod +x run

COPY --from=build app/build/codexgo .
