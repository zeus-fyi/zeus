REPO	:= zeusfyi
NAME    := hercules
GIT_SHA := $(shell git rev-parse HEAD)
IMG     := ${REPO}/${NAME}:${GIT_SHA}
LATEST  := ${REPO}/${NAME}:latest

docker.pubbuildx:
	@ docker buildx build -t ${IMG} -t ${LATEST} --platform=linux/amd64 -f ./docker/hercules/Dockerfile ./apps/hercules/ --push

docker.pull:
	@ docker pull zeusfyi/hercules:latest