Copy the config-sample.yaml and rename to config.yaml and fill in the values to config your usage or use the cli flags directly
See config yaml and/or main.go cmd flags for more details and groupings of flags

```text
######################################
##    Automation Step Selection     ##
######################################

# you can replace all with a comma separated list of steps to run
# e.g. only run step 7 to verify the lambda function
# or only step 9 to send the validator deposits

# STEP "1", "createSecretsAndStoreInAWS"
# STEP "2", "createInternalLambdaUser":
# STEP "3", "generateValidatorDeposits":
# STEP "4", "createLambdaFunctionKeystoresLayer":
# STEP "5", "createLambdaFunction":
# STEP "6", "createExternalLambdaUser":
# STEP "7", "verifyLambdaFunction":
# STEP "8", "createValidatorServiceRequestOnZeus":
# STEP "9", "sendValidatorDeposits":

# ACTIONS keywords
# all        - will run steps 1-9
# serverless - will run steps 1-7

# HELPERS: keywords
# use these keywords to fetch the secrets from aws secret manager and print them to the console

# getAgeEncryptionKeySecret
# getMnemonicHDWalletPasswordSecret
# getExternalLambdaAccessKeys
```

```text
Usage:
  Validator Key Generation and AWS Lambda Serverless Setup Automation [flags]

Flags:
      --age-private-key string           AGE_PRIVKEY: age private key
      --age-public-key string            AGE_PUBKEY: age public key
      --automation-steps string          AUTOMATION_STEPS: select which steps to automate and which order, using a comma separated list of numbers. default is all steps in order (default "all")
      --aws-access-key string            AWS_ACCESS_KEY: your private aws access key, which needs permissions to create iam users, roles, policies, secrets, and lambda functions and layers 
      --aws-account-number string        AWS_ACCOUNT_NUMBER: aws account number
      --aws-automation-on                AWS_AUTOMATION: automate the entire setup process on aws, requires you provide aws credentials
      --aws-secret-key string            AWS_SECRET_KEY: your private aws secret key 
      --bearer string                    BEARER: bearer token for validator service on zeus
      --eth1-addr-priv-key string        ETH1_PRIVATE_KEY: eth1 address private key for submitting deposits
      --ext-aws-access-key string        AWS_EXTERNAL_ACCESS_KEY: external access token for validator service on zeus
      --ext-aws-age-secret-name string   AWS_AGE_DECRYPTION_SECRET_NAME: the name of the secret that holds your age decryption keys for validator service on zeus
      --ext-aws-lambda-url string        AWS_LAMBDA_FUNC_URL: your lambda func url for validator service on zeus
      --ext-aws-secret-key string        AWS_EXTERNAL_SECRET_KEY: external secret token for validator service on zeus
      --fee-recipient string             FEE_RECIPIENT_ADDR: fee recipient address for validators service on zeus
      --hd-offset int                    HD_OFFSET_VALIDATORS: offset to start generating keys from hd wallet
      --hd-wallet-pw string              HD_WALLET_PASSWORD: hd wallet password (default "password")
  -h, --help                             help for Validator
      --key-group-name string            KEY_GROUP_NAME: name for validator service group on zeus
      --keygen                           KEYGEN_SECRETS: generates secrets for validator encryption and generation (default true)
      --keygen-validators                KEYGEN_VALIDATORS: generates validator deposits, with additional encrypted age keystore (default true)
      --keystores-dir-in string          KEYSTORE_DIR_IN: keystores directory in location (relative to builds dir) (default "./serverless/keystores")
      --keystores-dir-out string         KEYSTORE_DIR_OUT: keystores directory out location (relative to builds dir) (default "./serverless/keystores")
      --mnemonic string                  MNEMONIC_24_WORDS: twenty four word mnemonic to generate keystores
      --network string                   NETWORK: network to run on mainnet, goerli, ephemery, etc (default "ephemery")
      --node-url string                  NODE_URL: beacon for getting network data for validator deposit generation & submitting deposits (default "https://eth.ephemeral.zeus.fyi")
      --submit-deposits                  SUBMIT_DEPOSITS: submits validator deposits in keystore directory to the network for activation
      --submit-validator-service-req     SUBMIT_SERVICE_REQUEST: sends a request to zeus to setup a validator service
      --validator-count int              VALIDATORS_COUNT: number of keys to generate (default 3)
```
