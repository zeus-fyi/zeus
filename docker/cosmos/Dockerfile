#syntax=docker/dockerfile:1.4
FROM ubuntu:22.04 AS gaiad-builder
RUN apt-get update && apt-get install -y wget
WORKDIR /app

ARG VERSION

RUN wget https://github.com/cosmos/gaia/releases/download/${VERSION}/gaiad-${VERSION}-linux-amd64 -O gaiad
RUN chmod +x gaiad

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y curl ca-certificates

RUN curl -LO https://go.dev/dl/go1.20.3.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz \
    && rm go1.20.3.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
COPY --from=gaiad-builder /app/gaiad /usr/local/go/bin/

CMD ["/bin/bash"]
