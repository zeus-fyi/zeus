package ecdsa

import "github.com/zeus-fyi/gochain/web3/accounts"

type Account struct {
	*accounts.Account
}

func NewAccount(pk string) Account {
	a, err := accounts.ParsePrivateKey(pk)
	if err != nil {
		panic(err)
	}

	acc := Account{a}
	return acc
}
