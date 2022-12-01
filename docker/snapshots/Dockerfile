FROM golang:1.19-buster as snapshot_builder

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

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=snapshot_builder /app/snapshots /usr/bin/snapshots
