REPO	:= zeusfyi
NAME    := keydb
GIT_SHA := $(shell git rev-parse HEAD)
IMG     := ${REPO}/${NAME}:${GIT_SHA}
LATEST  := ${REPO}/${NAME}:latest
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64
T_DIGEST_VERSION := v1.0.0-rc.1

# --no-cache

docker.pubbuildx:
	@ docker buildx build -t ${IMG} -t ${LATEST} --build-arg VERSION=${T_DIGEST_VERSION} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ../../docker/keydb/Dockerfile ../../ --push

# redis-server --loadmodule /usr/local/lib/libredis_percentile.so --protected-mode no --bind 0.0.0.0 --dir /data
docker.debug:
	docker run -it -p 6379:6379 --entrypoint /bin/bash zeusfyi/keydb:latest
