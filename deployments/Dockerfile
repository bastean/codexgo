# syntax=docker/dockerfile:1

FROM golang:1.24-bookworm AS dev

WORKDIR /app

RUN apt update \
 && apt upgrade -y \
 && apt install -y sudo unzip \
 && go install github.com/go-task/task/v3/cmd/task@latest

COPY Taskfile.yml .

RUN task upgrade-managers \
 && task install-tools-dev

FROM dev AS test

RUN task install-tools-test

FROM dev AS build

COPY . .

RUN task build

FROM golang:1.24-bookworm AS prod

WORKDIR /app

COPY --from=build app/deployments/run.sh run

RUN chmod +x run

COPY --from=build app/build/codexgo .
