# shared variables
GIT_SHA := $(shell git rev-parse HEAD)
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64

SERVERLESS_REPO 	:= zeusfyi
SERVERLESS_NAME 	:= ethereumbls
SERVERLESS_IMG  	:= ${SERVERLESS_REPO}/${SERVERLESS_NAME}:${GIT_SHA}
SERVERLESS_LATEST := ${SERVERLESS_REPO}/${SERVERLESS_NAME}:latest

docker.pub.buildx:
	@ docker buildx build -t ${SERVERLESS_IMG} -t ${SERVERLESS_LATEST} --build-arg GOMODCACHE=${GOMODCACHE} --build-arg GOCACHE=${GOCACHE} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ../../docker/serverless/Dockerfile ../../

docker.debug:
	docker run -it --entrypoint /bin/bash ${SERVERLESS_REPO}/ethereumbls:latest
