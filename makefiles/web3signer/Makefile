REPO	:= zeusfyi
NAME    := web3signer
GIT_SHA := $(shell git rev-parse HEAD)
IMG     := ${REPO}/${NAME}:${GIT_SHA}
LATEST  := ${REPO}/${NAME}:latest
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64
VERSION := 22.11.0

docker.pubbuildx:
	@ docker buildx build -t ${IMG} -t ${LATEST} --build-arg VERSION=${VERSION} --build-arg GOMODCACHE=${GOMODCACHE} --build-arg GOCACHE=${GOCACHE} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ../../docker/web3signer/Dockerfile ../../

FOLDER_NAME=web3signer
VOLUME_NAME=${FOLDER_NAME}_postgres
ROOT_PATH=../../

pg.fresh:
	docker volume rm ${VOLUME_NAME}
	docker-compose -f ${ROOT_PATH}docker/datastores/local_postgres/web3signer/docker-compose-postgres.yml up

pg.start:
	docker-compose -f ${ROOT_PATH}docker/datastores/local_postgres/web3signer/docker-compose-postgres.yml up

pg.stop:
	docker-compose -f ${ROOT_PATH}docker/datastores/local_postgres/web3signer/docker-compose-postgres.yml down

docker.volume.rm:
	docker volume rm ${VOLUME_NAME}

docker.volume.ls:
	docker volume ls