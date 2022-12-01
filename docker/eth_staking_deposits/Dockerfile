FROM ubuntu:22.04 as eth-deposits-src
WORKDIR /app

ARG VERSION
RUN apt-get update && apt-get install -y ca-certificates git

RUN git clone https://github.com/ethereum/staking-deposit-cli.git --branch ${VERSION} --single-branch
RUN cd staking-deposit-cli && git checkout ${VERSION} && git pull

COPY . .

FROM python:alpine3.14 as eth-deposits
WORKDIR /app

RUN apk add --update gcc libc-dev linux-headers

COPY --from=eth-deposits-src /app/staking-deposit-cli .

RUN pip3 install -r requirements.txt
RUN python3 setup.py install
