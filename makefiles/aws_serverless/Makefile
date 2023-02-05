# shared variables
GIT_SHA := $(shell git rev-parse HEAD)
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64

AWS_REPO := $(shell echo $AWS_REPO)

SERVERLESS_REPO 	:= AWS_REPO
SERVERLESS_NAME 	:= ethereumbls
SERVERLESS_IMG  	:= ${SERVERLESS_REPO}/${SERVERLESS_NAME}:${GIT_SHA}
SERVERLESS_LATEST := ${SERVERLESS_REPO}/${SERVERLESS_NAME}:latest

docker.pub.buildx:
	@ docker buildx build -t ${SERVERLESS_IMG} -t ${SERVERLESS_LATEST} --build-arg GOMODCACHE=${GOMODCACHE} --build-arg GOCACHE=${GOCACHE} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --platform=${GOOS}/${GOARCH} -f ../../docker/serverless/Dockerfile ../../

docker.debug:
	docker run -it --entrypoint /bin/bash ${SERVERLESS_REPO}/ethereumbls:latest

docker cp $(shell docker container ls | awk 'NR==2 {print $1}') /usr/bin/ethereumsignbls .


docker cp 480391564655.dkr.ecr.us-west-1.amazonaws.com/ethereumbls:latest:/usr/bin/ethereumsignbls ~/ethereumsignbls

	docker run -it --entrypoint /bin/bash 480391564655.dkr.ecr.us-west-1.amazonaws.com/ethereumbls:latest

 docker container ls -aqf "name=480391564655.dkr.ecr.us-west-1.amazonaws.com/ethereumbls:latest" | xargs docker inspect --format '{{.Id}}'

docker ps -aqf "name=480391564655.dkr.ecr.us-west-1.amazonaws.com/ethereumbls:latest" | xargs docker inspect --format '{{.Id}}'

docker cp $(docker container ls | awk 'NR==2 {print $1}'):/usr/bin/ethereumsignbls .