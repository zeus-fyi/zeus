REPO	:= zeusfyi
NAME    := staking-deposit-cli
GIT_SHA := $(shell git rev-parse HEAD)
IMG     := ${REPO}/${NAME}:${GIT_SHA}
LATEST  := ${REPO}/${NAME}:latest
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64
VERSION := v2.3.0

docker.pubbuildx:
	@ docker buildx build -t ${IMG} -t ${LATEST} --build-arg VERSION=${VERSION} --platform=${GOOS}/${GOARCH} -f ../../docker/eth_staking_deposits/Dockerfile .

docker.debug:
	docker run -it --entrypoint /bin/sh zeusfyi/staking-deposit-cli:latest

