#syntax=docker/dockerfile:1.4
FROM golang:1.21-bookworm as snapshot_builder

WORKDIR /app

ARG GOMODCACHE
ARG GOCACHE
ARG GOOS
ARG GOARCH
# manage dependencies
COPY . .
RUN GO_ENABLED=1 GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags="-s -w" -o snapshots apps/snapshots/main.go

FROM ubuntu:22.04 AS rclone_builder
RUN apt-get update && apt-get install -y ca-certificates curl unzip

WORKDIR /app

RUN curl -O https://downloads.rclone.org/rclone-current-linux-amd64.zip
RUN unzip rclone-current-linux-amd64.zip
RUN cd rclone-*-linux-amd64 && cp rclone /usr/bin/
RUN chown root:root /usr/bin/rclone
RUN chmod 755 /usr/bin/rclone

FROM ubuntu:22.04 AS aws-cli
RUN apt-get update && apt-get install -y curl unzip
WORKDIR /app

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y ca-certificates unzip less

COPY --from=snapshot_builder /app/snapshots /usr/bin/snapshots
COPY --from=rclone_builder /usr/bin/rclone /usr/bin/rclone
COPY --from=aws-cli /app/awscliv2.zip /

RUN unzip awscliv2.zip
RUN rm awscliv2.zip
RUN chmod +x ./aws/install

RUN ./aws/install

RUN mkdir -p ~/.config/rclone
