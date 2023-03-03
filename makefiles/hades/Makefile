# shared variables
GIT_SHA := $(shell git rev-parse HEAD)
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64

# for HADES cookbook echo server template
HADES_REPO 	:= zeusfyi
HADES_NAME 	:= hades
HADES_IMG  	:= ${HADES_REPO}/${HADES_NAME}:${GIT_SHA}
HADES_LATEST := ${HADES_REPO}/${HADES_NAME}:latest

docker.pubbuildx:
	@ docker buildx build -t ${HADES_IMG} -t ${HADES_LATEST} --build-arg GOMODCACHE=${GOMODCACHE} --build-arg GOCACHE=${GOCACHE} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ../../docker/hades/Dockerfile ../../ --push

