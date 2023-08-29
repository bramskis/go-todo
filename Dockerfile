FROM golang:1.21-alpine3.18 AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    git \
    && update-ca-certificates

FROM base AS builder
WORKDIR /app

COPY . /app
RUN go mod download \
    && go mod verify

RUN go build -o todo -a .

FROM alpine:latest as prod

COPY --from=builder /app/todo /usr/local/bin/todo
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/todo"]
