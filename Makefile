# shared variables
GIT_SHA := $(shell git rev-parse HEAD)
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64
VERSION := 0.2.1

# hercules build info
REPO	:= zeusfyi
NAME    := hercules
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

docker.debug:
	docker run -it --entrypoint /bin/bash zeusfyi/hercules:latest

build.staking.cli:
	go build -o ./builds/serverless/bin/serverless ./builds/serverless

# generates new mnemonic, age encryption key, uses default hd password if none provided, and creates keystores
# zipped age encrypted file for serverless app --keygen true/false will toggle new keygen creation
serverless.keygen:
	./builds/serverless/bin/serverless --keygen true --num-keys 3

ETH1_PRIV_KEY := ""
# you will need an eth1 address and it must have 32 Eth + gas fees to deposit per validator
serverless.submit.deposits:
	./builds/serverless/bin/serverless --keygen false --submit-deposits true --eth1-addr-priv-key ${ETH1_PRIV_KEY}



