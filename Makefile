# shared variables
GIT_SHA := $(shell git rev-parse HEAD)
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64
VERSION := 0.3.0-rc.0

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

# arch -x86_64
# export CGO_CFLAGS_ALLOW="-D__BLST_PORTABLE__"
build.staking.cli.mac.intel:
	CGO_CFLAGS="-O -D__BLST_PORTABLE__" CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 CC="clang" CXX="zig c++ -target x86_64-macos" go build --tags extended -o ./builds/serverless/bin_mac_intel/serverless ./builds/serverless

#build.staking.cli.windows:
#	CGO_CFLAGS="-O -D__BLST_PORTABLE__" CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC="zig cc -target x86_64-windows-gnu" CXX="zig c++ -target x86_64-windows-gnu" go build -trimpath -ldflags='-H=windowsgui' -o ./builds/serverless/bin_windows/serverless ./builds/serverless

AWS_ACCOUNT_NUMBER:= ""
AWS_ACCESS_KEY := ""
AWS_SECRET_KEY := ""

HD_OFFSET_VALIDATORS := 0
VALIDATORS_COUNT := 0
AUTOMATION_STEPS := serverless

ETH1_PRIV_KEY := ""
# you will need an eth1 address and it must have 32 Eth + gas fees to deposit per validator

VALIDATORS_COUNT := 3
BEARER := ""

# cli suffix indicates you need to drive with flags, otherwise assumes your the config file is used to drive the automation

serverless.automation.cli:
	./builds/serverless/bin/serverless --aws-account-number=$(AWS_ACCOUNT_NUMBER) --aws-access-key=$(AWS_ACCESS_KEY) --aws-secret-key=$(AWS_SECRET_KEY) --validator-count=$(VALIDATORS_COUNT) --automation-steps=$(AUTOMATION_STEPS)

serverless.validator.gen.cli:
	./builds/serverless/bin/serverless --aws-account-number=$(AWS_ACCOUNT_NUMBER) --aws-access-key=$(AWS_ACCESS_KEY) --aws-secret-key=$(AWS_SECRET_KEY) --validator-count=$(VALIDATORS_COUNT) --hd-offset=$(HD_OFFSET_VALIDATORS) --automation-steps=generateValidatorDeposits

serverless.verify.cli:
	./builds/serverless/bin/serverless --aws-account-number=$(AWS_ACCOUNT_NUMBER) --aws-access-key=$(AWS_ACCESS_KEY) --aws-secret-key=$(AWS_SECRET_KEY) --automation-steps=verifyLambdaFunction

serverless.deploy.all.cli:
	./builds/serverless/bin/serverless --aws-account-number=$(AWS_ACCOUNT_NUMBER) --aws-access-key=$(AWS_ACCESS_KEY) --aws-secret-key=$(AWS_SECRET_KEY) --validator-count=$(VALIDATORS_COUNT) --eth1-addr-priv-key=$(ETH1_PRIV_KEY) --bearer=$(BEARER) --automation-steps=all

# config.yaml driven automation
serverless.submit.deposits:
	./builds/serverless/bin/serverless --automation-steps=sendValidatorDeposits

serverless.verify.config:
	./builds/serverless/bin/serverless --automation-steps=verifyLambdaFunction

serverless.service.config:
	./builds/serverless/bin/serverless --automation-steps=createValidatorServiceRequestOnZeus

serverless.deploy.all.config:
	./builds/serverless/bin/serverless --automation-steps=all

serverless.validators.gen:
	./builds/serverless/bin/serverless --automation-steps=generateValidatorDeposits

# use prefix: arch -x86_64 (for m1 macs)
serverless.validators.gen.mac.intel:
	./builds/serverless/bin_mac_intel/serverless --automation-steps=generateValidatorDeposits

serverless.deploy.all.config.mac.intel:
	./builds/serverless/bin_mac_intel/serverless --automation-steps=all

serverless.verify.config.mac.intel:
	./builds/serverless/bin_mac_intel/serverless --automation-steps=verifyLambdaFunction
