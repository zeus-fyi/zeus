REPO	:= zeusfyi
NAME    := hercules
GIT_SHA := $(shell git rev-parse HEAD)
IMG     := ${REPO}/${NAME}:${GIT_SHA}
LATEST  := ${REPO}/${NAME}:latest
VERSION  := 0.0.3-rc.0

docker.pubbuildx:
	@ docker buildx build -t ${IMG} -t ${LATEST} --platform=linux/amd64 -f ./docker/hercules/Dockerfile ./apps/hercules/ --push

docker.pull:
	@ docker pull zeusfyi/hercules:latest

tag:
	git tag v${VERSION}

tag.push:
	git tag push origin v${VERSION}
