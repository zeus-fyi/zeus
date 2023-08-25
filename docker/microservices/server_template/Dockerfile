FROM golang:1.21-bookworm as builder

WORKDIR /app

ARG GOMODCACHE
ARG GOCACHE
ARG GOOS
ARG GOARCH
# manage dependencies
COPY . .
RUN --mount=type=cache,target=${GOMODCACHE} go mod download
COPY . .

RUN --mount=type=cache,target=${GOCACHE} CGO_ENABLED=1 GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="-s -w" -o api apps/cookbooks/microservices/server_template/main.go

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /app/api /usr/bin/api

EXPOSE 9090