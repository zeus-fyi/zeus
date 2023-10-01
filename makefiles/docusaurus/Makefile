REPO	:= zeusfyi
HOME := "home"
DOCS := "zeusfyi"
TEMPLATE := "template"
NAME    := docusaurus
GIT_SHA := $(shell git rev-parse HEAD)
IMG     := ${REPO}/${NAME}:${GIT_SHA}
IMG_HOME     := ${REPO}/${HOME}:${GIT_SHA}
LATEST  := ${REPO}/${NAME}:latest
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64

docker.pubbuildx:
	@ docker buildx build -t ${IMG} -t ${LATEST} --build-arg DIR=${DOCS} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ../../docker/docusaurus/Dockerfile ../../ --push

docker.pubbuildx.home:
	@ docker buildx build -t ${IMG_HOME} -t ${LATEST} --build-arg DIR=${HOME} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ../../docker/docusaurus/Dockerfile ../../ --push

docker.pubbuildx.template:
	@ docker buildx build -t ${IMG} -t ${LATEST} --build-arg DIR=${TEMPLATE} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ../../docker/docusaurus/Dockerfile ../../

docker.debug:
	docker run -it -p 3000:3000 --entrypoint /bin/bash zeusfyi/docusaurus:latest

