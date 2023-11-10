REPO	:= zeusfyi
NAME    := docusaurus-template
GIT_SHA := $(shell git rev-parse HEAD)
IMG     := ${REPO}/${NAME}:${GIT_SHA}
LATEST  := ${REPO}/${NAME}:latest
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64

docker.pubbuildx:
	@ docker buildx build -t ${IMG} -t ${LATEST} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ../../docker/docusaurus/Dockerfile ../../ --push

docker.debug:
	docker run -it -p 3000:3000 --entrypoint /bin/bash zeusfyi/docusaurus:latest
