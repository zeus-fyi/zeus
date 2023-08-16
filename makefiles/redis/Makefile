REPO	:= zeusfyi
NAME    := redis
GIT_SHA := $(shell git rev-parse HEAD)
IMG     := ${REPO}/${NAME}:${GIT_SHA}
LATEST  := ${REPO}/${NAME}:latest
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64
T_DIGEST_VERSION := v0.4.2

docker.pubbuildx:
	@ docker buildx build -t ${IMG} -t ${LATEST} --build-arg VERSION=${T_DIGEST_VERSION} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ../../docker/redis/Dockerfile ../../
