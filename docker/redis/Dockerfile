#syntax=docker/dockerfile:1.4
FROM ubuntu:22.04 as t-digest-builder
WORKDIR /app

ARG VERSION
RUN apt-get update && apt-get install -y ca-certificates git build-essential cmake clang-format
RUN git clone https://github.com/RedisBloom/t-digest-c.git --branch ${VERSION} --single-branch

WORKDIR /app/t-digest-c
RUN git checkout ${VERSION} && git pull && git submodule update --init --recursive
RUN cmake
RUN make library_all

FROM redis:7.0.13 as redis
WORKDIR /app

# Copy the t-digest shared and static libraries
COPY --from=t-digest-builder /app/t-digest-c/build/src/libtdigest.so /usr/local/lib/
COPY --from=t-digest-builder /app/t-digest-c/build/src/libtdigest_static.a /usr/local/lib/

EXPOSE 6379