#syntax=docker/dockerfile:1.4
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

FROM ubuntu:22.04 as aws_eks
RUN apt-get update && apt-get install -y wget curl
WORKDIR /app

ARG PLATFORM=Linux_amd64
RUN curl -sLO "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$PLATFORM.tar.gz"

# (Optional) Verify checksum
RUN curl -sL "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_checksums.txt" | grep $PLATFORM | sha256sum --check

RUN tar -xzf eksctl_$PLATFORM.tar.gz -C /tmp && rm eksctl_$PLATFORM.tar.gz

RUN mv /tmp/eksctl /app/eksctl

FROM ubuntu:22.04 as aws_cli
RUN apt-get update && apt-get install -y wget curl
WORKDIR /app

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y ca-certificates unzip less

COPY --from=aws_eks /app/eksctl /usr/local/bin
COPY --from=aws_cli /app/awscliv2.zip /
COPY --from=builder /app/hades /usr/bin/hades
COPY --from=doctl /app/doctl /usr/local/bin
COPY --from=gcp /app/google-cloud-cli.tar.gz /

RUN unzip awscliv2.zip
RUN rm awscliv2.zip
RUN chmod +x ./aws/install

RUN ./aws/install

RUN tar xf google-cloud-cli.tar.gz
RUN rm google-cloud-cli.tar.gz
RUN ./google-cloud-sdk/install.sh
RUN /google-cloud-sdk/bin/gcloud components install gke-gcloud-auth-plugin
RUN chmod +x /google-cloud-sdk/bin/gcloud
RUN chmod +x /google-cloud-sdk/bin/gke-gcloud-auth-plugin

ENV PATH $PATH:/google-cloud-sdk/bin

EXPOSE 8888