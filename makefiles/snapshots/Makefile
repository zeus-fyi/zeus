REPO	:= zeusfyi
NAME    := snapshots
GIT_SHA := $(shell git rev-parse HEAD)
IMG     := ${REPO}/${NAME}:${GIT_SHA}
LATEST  := ${REPO}/${NAME}:latest
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64
VERSION := v9.0.3

docker.pubbuildx:
	@ docker buildx build -t ${IMG} -t ${LATEST} --build-arg VERSION=${VERSION} --build-arg GOMODCACHE=${GOMODCACHE} --build-arg GOCACHE=${GOCACHE} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ../../docker/snapshots/Dockerfile ../../ --push

docker.debug:
	docker run -it --entrypoint /bin/bash zeusfyi/snapshots:latest
