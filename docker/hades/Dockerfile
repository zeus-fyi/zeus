#syntax=docker/dockerfile:1.4
FROM golang:1.20.1-buster as builder

WORKDIR /app

ARG GOMODCACHE
ARG GOCACHE
ARG GOOS
ARG GOARCH
# manage dependencies
COPY . .
RUN --mount=type=cache,target=${GOMODCACHE} go mod download
COPY . .

RUN --mount=type=cache,target=${GOCACHE} GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="-s -w" -o hades apps/hades/main.go

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /app/hades /usr/bin/hades

EXPOSE 8888