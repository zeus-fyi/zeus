FROM golang:1.20.1-buster as builder

WORKDIR /app

ARG GOMODCACHE
ARG GOCACHE
ARG GOOS
ARG GOARCH
# manage dependencies
COPY go.mod go.sum ./

RUN --mount=type=cache,target=${GOMODCACHE} go mod download
COPY .. .

RUN --mount=type=cache,target=${GOCACHE} CGO_ENABLED=1 GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="-s -w" -o secretskeygen apps/serverless/ethereum/secrets_keygen/main.go

FROM ubuntu:22.04
RUN apt-get update

COPY --from=builder /app/secretskeygen /usr/bin/secretskeygen
