#syntax=docker/dockerfile:1.4
FROM rust:latest as t-digest-builder
WORKDIR /app

ARG VERSION

RUN apt-get update && apt-get install -y ca-certificates git libclang-dev libc6-dev llvm build-essential
RUN git clone https://github.com/zeus-fyi/redis-percentile.git --branch ${VERSION} --single-branch

WORKDIR /app/redis-percentile
RUN git checkout ${VERSION} && git pull
RUN cargo update
RUN cargo build --release

FROM redis:7.0.12 as redis
WORKDIR /app

COPY --from=t-digest-builder /app/redis-percentile/target/release/libredis_percentile.so /usr/local/lib/
EXPOSE 6379

