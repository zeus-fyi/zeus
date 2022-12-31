FROM ubuntu:22.04 as web3signer
RUN apt-get update && apt-get install -y wget
WORKDIR /app

ARG VERSION

RUN wget https://artifacts.consensys.net/public/web3signer/raw/names/web3signer.tar.gz/versions/${VERSION}/web3signer-${VERSION}.tar.gz
RUN mv web3signer-${VERSION}.tar.gz web3signer
RUN tar xf web3signer

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=web3signer /app/web3signer /usr/local/bin

EXPOSE 9000