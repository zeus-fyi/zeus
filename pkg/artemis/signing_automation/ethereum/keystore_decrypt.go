package signing_automation_ethereum

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	keystorev4 "github.com/wealdtech/go-eth2-wallet-encryptor-keystorev4"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
)

func DecryptKeystoreCipher(ctx context.Context, keystoreJSON map[string]interface{}, password string) ([]byte, error) {
	ks := keystorev4.New()

	cryptoMsg := keystoreJSON["crypto"]
	b, err := json.Marshal(cryptoMsg)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, err
	}

	cryptoInput := make(map[string]interface{})
	err = json.Unmarshal(b, &cryptoInput)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, err
	}
	sk, err := ks.Decrypt(cryptoInput, password)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return nil, err
	}
	return sk, err
}

func DecryptKeystoreCipherIntoEthSignerBLS(ctx context.Context, keystoreJSON map[string]interface{}, password string) (bls_signer.EthBLSAccount, error) {
	sk, err := DecryptKeystoreCipher(ctx, keystoreJSON, password)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return bls_signer.EthBLSAccount{}, err
	}
	acc := bls_signer.NewEthSignerBLSFromExistingKeyBytes(sk)
	return acc, nil
}
