# shared variables
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64
VERSION := 0.0.5-rc.0

# hercules build info
REPO	:= zeusfyi
NAME    := hercules
GIT_SHA := $(shell git rev-parse HEAD)
IMG     := ${REPO}/${NAME}:${GIT_SHA}
LATEST  := ${REPO}/${NAME}:latest

docker.pubbuildx:
	@ docker buildx build -t ${IMG} -t ${LATEST} --build-arg GOMODCACHE=${GOMODCACHE} --build-arg GOCACHE=${GOCACHE} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ./docker/hercules/Dockerfile . --push

docker.pull:
	@ docker pull zeusfyi/hercules:latest

tag:
	git tag v${VERSION}

tag.push:
	git push origin v${VERSION}

# for microservice cookbook echo server template
MICROSERVICE_REPO 	:= zeusfyi
MICROSERVICE_NAME 	:= microservice
MICROSERVICE_IMG  	:= ${MICROSERVICE_REPO}/${MICROSERVICE_NAME}:${GIT_SHA}
MICROSERVICE_LATEST := ${MICROSERVICE_REPO}/${MICROSERVICE_NAME}:latest

docker.microservice.buildx:
	@ docker buildx build -t ${MICROSERVICE_IMG} -t ${LATEST} --build-arg GOMODCACHE=${GOMODCACHE} --build-arg GOCACHE=${GOCACHE} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ./docker/microservices/server_template/Dockerfile .
