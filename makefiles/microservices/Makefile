# shared variables
GIT_SHA := $(shell git rev-parse HEAD)
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64

# for microservice cookbook echo server template
MICROSERVICE_REPO 	:= zeusfyi
MICROSERVICE_NAME 	:= microservice
MICROSERVICE_IMG  	:= ${MICROSERVICE_REPO}/${MICROSERVICE_NAME}:${GIT_SHA}
MICROSERVICE_LATEST := ${MICROSERVICE_REPO}/${MICROSERVICE_NAME}:latest

docker.microservice.buildx:
	@ docker buildx build -t ${MICROSERVICE_IMG} -t ${MICROSERVICE_LATEST} --build-arg GOMODCACHE=${GOMODCACHE} --build-arg GOCACHE=${GOCACHE} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ../../docker/microservices/server_template/Dockerfile ../../

CHOREOGRAPHY_NAME 	:= choreography
CHOREOGRAPHY_IMG  	:= ${MICROSERVICE_REPO}/${CHOREOGRAPHY_NAME}:${GIT_SHA}
CHOREOGRAPHY_LATEST := ${MICROSERVICE_REPO}/${CHOREOGRAPHY_NAME}:latest

docker.choreography.buildx:
	@ docker buildx build -t ${CHOREOGRAPHY_IMG} -t ${CHOREOGRAPHY_LATEST} --build-arg GOMODCACHE=${GOMODCACHE} --build-arg GOCACHE=${GOCACHE} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ../../docker/microservices/choreography/Dockerfile ../../
