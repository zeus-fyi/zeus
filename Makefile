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

#Usage:
#  Validator Key Generation and AWS Lambda Serverless Setup Automation [flags]

# Flags:
#      --age-private-key string           AGE_PRIVKEY: age private key
#      --age-public-key string            AGE_PUBKEY: age public key
#      --aws-access-key string            AWS_ACCESS_KEY: aws access key, which needs permissions to create iam users, roles, policies, secrets, and lambda functions and layers
#      --aws-account-number string        AWS_ACCOUNT_NUMBER: aws account number
#      --aws-automation-on                AWS_AUTOMATION: automate the entire setup process on aws, requires you provide aws credentials
#      --aws-secret-key string            AWS_SECRET_KEY: aws secret key
#      --bearer string                    BEARER: bearer token for validator service on zeus
#      --eth1-addr-priv-key string        ETH1_PRIVATE_KEY: eth1 address private key for submitting deposits
#      --ext-aws-access-key string        AWS_EXTERNAL_ACCESS_KEY: bearer token for validator service on zeus
#      --ext-aws-age-secret-name string   AWS_AGE_DECRYPTION_SECRET_NAME: bearer token for validator service on zeus
#      --ext-aws-lambda-url string        AWS_LAMBDA_FUNC_URL: bearer token for validator service on zeus
#      --ext-aws-secret-key string        AWS_EXTERNAL_SECRET_KEY: bearer token for validator service on zeus
#      --fee-recipient string             FEE_RECIPIENT_ADDR: fee recipient address for validators service on zeus
#      --hd-offset int                    HD_OFFSET_VALIDATORS: offset to start generating keys from hd wallet
#      --hd-wallet-pw string              HD_WALLET_PASSWORD: hd wallet password
#      --key-group-name string            KEY_GROUP_NAME: name for validator service group on zeus
#      --keygen                           KEYGEN_SECRETS: generates secrets for validator encryption and generation (default true)
#      --keygen-validators                KEYGEN_VALIDATORS: generates validator deposits, with additional encrypted age keystore (default true)
#      --keystores-dir-in string          KEYSTORE_DIR_IN: keystores directory in location (relative to builds dir) (default "./serverless/keystores")
#      --keystores-dir-out string         KEYSTORE_DIR_OUT: keystores directory out location (relative to builds dir) (default "./serverless/keystores")
#      --mnemonic string                  MNEMONIC_24_WORDS: twenty four word mnemonic to generate keystores
#      --network string                   NETWORK: network to run on (mainnet, goerli, ephemery, etc (default "ephemery")
#      --node-url string                  NODE_URL: beacon for getting network data for validator deposit generation & submitting deposits (default "https://eth.ephemeral.zeus.fyi")
#      --submit-deposits                  SUBMIT_DEPOSITS: submits validator deposits in keystore directory to the network for activation
#      --submit-validator-service-req     SUBMIT_SERVICE_REQUEST: sends a request to zeus to setup a validator service
#      --validator-count int              VALIDATORS_COUNT: number of keys to generate (default 3)


