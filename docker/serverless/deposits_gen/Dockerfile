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

RUN --mount=type=cache,target=${GOCACHE} CGO_ENABLED=1 GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="-s -w" -o validatorsdepositsgen apps/cookbooks/serverless/ethereum/validators_deposits_gen/main.go

FROM ubuntu:22.04
RUN apt-get update

COPY --from=builder /app/validatorsdepositsgen /usr/bin/validatorsdepositsgen
