FROM golang:bookworm AS dev

WORKDIR /app

RUN apt update && apt upgrade -y

RUN apt install -y nodejs npm

COPY Makefile .

RUN make dev-tools

FROM dev AS test

WORKDIR /app

RUN make test-tools

FROM dev AS build

WORKDIR /app

COPY . .

RUN make build

FROM golang:bookworm AS prod

WORKDIR /app

COPY --from=build app/deployments/run.sh run

RUN chmod +x run

COPY --from=build app/build/codexgo .
