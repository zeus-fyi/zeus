package signing_automation_ethereum

import (
	"github.com/zeus-fyi/gochain/web3/accounts"
	web3_actions "github.com/zeus-fyi/gochain/web3/client"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
)

type Web3SignerClient struct {
	web3_actions.Web3Actions
}

func NewWeb3Client(nodeUrl string, acc *accounts.Account) Web3SignerClient {
	err := bls_signer.InitEthBLS()
	if err != nil {
		panic(err)
	}
	w := web3_actions.NewWeb3ActionsClientWithAccount(nodeUrl, acc)
	return Web3SignerClient{w}
}
