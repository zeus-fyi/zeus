#syntax=docker/dockerfile:1.4
FROM golang:1.20.1-buster as snapshot_builder

WORKDIR /app

ARG GOMODCACHE
ARG GOCACHE
ARG GOOS
ARG GOARCH
# manage dependencies
COPY . .
RUN --mount=type=cache,target=${GOMODCACHE} go mod download
COPY . .
RUN --mount=type=cache,target=${GOCACHE} CGO_ENABLED=1 GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="-s -w" -o snapshots apps/snapshots/main.go

FROM ubuntu:22.04 AS rclone_builder
RUN apt-get update && apt-get install -y ca-certificates curl unzip

WORKDIR /app

RUN curl -O https://downloads.rclone.org/rclone-current-linux-amd64.zip
RUN unzip rclone-current-linux-amd64.zip
RUN cd rclone-*-linux-amd64 && cp rclone /usr/bin/
RUN chown root:root /usr/bin/rclone
RUN chmod 755 /usr/bin/rclone

FROM ubuntu:22.04 AS gaiad-builder
RUN apt-get update && apt-get install -y wget
WORKDIR /app

ARG VERSION

RUN wget https://github.com/cosmos/gaia/releases/download/${VERSION}/gaiad-${VERSION}-linux-amd64 -O gaiad
RUN chmod +x gaiad

FROM ubuntu:22.04 AS gaiad-genesis
RUN apt-get update && apt-get install -y wget
WORKDIR /app

RUN wget https://github.com/cosmos/testnets/raw/master/public/genesis.json.gz -O genesis.json.gz
RUN gzip -d genesis.json.gz

FROM golang:1.18.5-buster
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=snapshot_builder /app/snapshots /usr/bin/snapshots
COPY --from=rclone_builder /usr/bin/rclone /usr/bin/rclone

COPY --from=gaiad-builder /app/gaiad /usr/local/go/bin/
COPY --from=gaiad-genesis /app/genesis.json /etc/gaiad/genesis.json

RUN mkdir -p ~/.config/rclone
