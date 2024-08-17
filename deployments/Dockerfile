# syntax=docker/dockerfile:1

FROM golang:1.23-bookworm AS dev

WORKDIR /app

RUN apt update \
 && apt upgrade -y \
 && apt install -y sudo nodejs npm

COPY Makefile .

RUN make install-tools-dev

FROM dev AS test

RUN make install-tools-test

FROM dev AS build

COPY . .

RUN make build

FROM golang:1.23-bookworm AS prod

WORKDIR /app

COPY --from=build app/deployments/run.sh run

RUN chmod +x run

COPY --from=build app/build/codexgo .
