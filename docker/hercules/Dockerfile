FROM golang:1.19-buster as builder

WORKDIR /app

# manage dependencies
COPY . .
COPY go.mod .
COPY go.sum .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o hercules

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /app/hercules /usr/bin/hercules
EXPOSE 9003