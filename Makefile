# shared variables
GIT_SHA := $(shell git rev-parse HEAD)
GOMODCACHE := $(shell go env GOMODCACHE)
GOCACHE := $(shell go env GOCACHE)
GOOS 	:= linux
GOARCH  := amd64
VERSION := 0.2.0

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

#Flags:
#      --age-private-key string      age private key
#      --age-public-key string       age public key
#      --aws-access-key string       aws access key, which needs permissions to create iam users, roles, policies, secrets, and lambda functions and layers
#      --aws-account-number string   aws account number
#      --aws-automation-on           automate the entire setup process on aws, requires you provide aws credentials
#      --aws-secret-key string       aws secret key
#      --eth1-addr-priv-key string   eth1 address private key for submitting deposits
#      --hd-offset int               offset to start generating keys from hd wallet
#      --hd-wallet-pw string         hd wallet password
#  -h, --help                        help for Web3
#      --keygen                      generates secrets for validator encryption and generation (default true)
#      --keygen-validators           generates validator deposits, with additional encrypted age keystore (default true)
#      --keystores-dir-in string     keystores directory in location (relative to builds dir) (default "./serverless/keystores")
#      --keystores-dir-out string    keystores directory out location (relative to builds dir) (default "./serverless/keystores")
#      --mnemonic string             twenty four word mnemonic to generate keystores
#      --network string              network to run on (mainnet, goerli, ephemery, etc (default "ephemery")
#      --node-url string             beacon for getting network data for validator deposit generation & submitting deposits (default "https://eth.ephemeral.zeus.fyi")
#      --num-keys int                number of keys to generate (default 3)
#      --submit-deposits             submits validator deposits in keystore directory to the network for activation

