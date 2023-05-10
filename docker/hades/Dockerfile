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

FROM ubuntu:22.04 as doctl
RUN apt-get update && apt-get install -y wget
WORKDIR /app

RUN wget https://github.com/digitalocean/doctl/releases/download/v1.85.0/doctl-1.85.0-linux-amd64.tar.gz -O doctl.tar.gz
RUN tar xf doctl.tar.gz

FROM ubuntu:22.04 as gcp
RUN apt-get update && apt-get install -y wget
WORKDIR /app

RUN wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-429.0.0-linux-x86_64.tar.gz -O google-cloud-cli.tar.gz

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /app/hades /usr/bin/hades
COPY --from=doctl /app/doctl /usr/local/bin
COPY --from=gcp /app/google-cloud-cli.tar.gz /

RUN tar xf google-cloud-cli.tar.gz
RUN rm google-cloud-cli.tar.gz
RUN ./google-cloud-sdk/install.sh
RUN /google-cloud-sdk/bin/gcloud components install gke-gcloud-auth-plugin
RUN chmod +x /google-cloud-sdk/bin/gcloud
RUN chmod +x /google-cloud-sdk/bin/gke-gcloud-auth-plugin

ENV PATH $PATH:/google-cloud-sdk/bin

EXPOSE 8888