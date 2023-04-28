#syntax=docker/dockerfile:1.4
FROM ubuntu:22.04 AS gaiad-builder
RUN apt-get update && apt-get install -y wget
WORKDIR /app

ARG VERSION

RUN wget https://github.com/cosmos/gaia/releases/download/${VERSION}/gaiad-${VERSION}-linux-amd64 -O gaiad

FROM golang:1.18.5-buster
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=gaiad-builder /app/gaiad $HOME/go/bin/
EXPOSE 26656 26657 1317 9090

ENTRYPOINT ["gaiad", "start"]