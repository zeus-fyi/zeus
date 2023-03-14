package bls_serverless_signing

import spec "github.com/attestantio/go-eth2-client/spec/phase0"

type BlsKeyGenRequests struct {
	MnemonicAndHDWalletSecretName string `json:"mnemonicAndHDWalletSecretName,omitempty"`
	Mnemonic                      string `json:"mnemonic,omitempty"`
	HdWalletPassword              string `json:"hdWalletPassword,omitempty"`

	AgeSecretName string `json:"ageSecretName,omitempty"`
	AgePubKey     string `json:"agePubKey,omitempty"`
	AgePrivKey    string `json:"agePrivKey,omitempty"`
}

type EthereumValidatorEncryptedZipKeysRequests struct {
	AgeSecretName                 string `json:"ageSecretName,omitempty"`
	MnemonicAndHDWalletSecretName string `json:"mnemonicAndHDWalletSecretName,omitempty"`
	ValidatorCount                int    `json:"validatorCount,omitempty"`
	HdOffset                      int    `json:"hdOffset,omitempty"`
}

type EthereumValidatorDepositsGenRequests struct {
	MnemonicAndHDWalletSecretName string `json:"mnemonicAndHDWalletSecretName"`
	WithdrawalAddress             string `json:"withdrawalAddress,omitempty"`
	ValidatorCount                int    `json:"validatorCount"`
	HdOffset                      int    `json:"hdOffset,omitempty"`

	Network     string        `json:"network"`
	ForkVersion *spec.Version `json:"forkVersion"`
	BeaconURL   string        `json:"beaconURL,omitempty"`
}
