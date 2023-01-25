package aegis_inmemdbs

import (
	"context"

	"github.com/rs/zerolog/log"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
)

type EthereumBLSKeySignatureRequests struct {
	Map map[string]EthereumBLSKeySignatureRequest
}

type EthereumBLSKeySignatureRequest struct {
	Message string `json:"message"`
}

type EthereumBLSKeySignatureResponses struct {
	Map map[string]EthereumBLSKeySignatureResponse
}

type EthereumBLSKeySignatureResponse struct {
	Signature string `json:"signature"`
}

func SignValidatorMessagesFromInMemDb(ctx context.Context, signReqs EthereumBLSKeySignatureRequests) (EthereumBLSKeySignatureResponses, error) {
	resp := make(map[string]EthereumBLSKeySignatureResponse)
	batchResp := EthereumBLSKeySignatureResponses{
		Map: resp,
	}
	if len(signReqs.Map) == 0 {
		return batchResp, nil
	}
	txn := ValidatorInMemDB.Txn(false)
	defer txn.Abort()
	it, err := txn.Get("validators", "id")
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("SignValidatorMessagesFromInMemDb")
		return batchResp, err
	}

	tmp := make(map[string]Validator)
	for obj := it.Next(); obj != nil; obj = it.Next() {
		inMemDB := obj.(inMemValidator)
		v := bls_signer.NewEthSignerBLSFromExistingKeyBytes(inMemDB.SecretKey)
		pubkey := bls_signer.ConvertBytesToString(v.PublicKey().Marshal())
		tmp[pubkey] = NewValidator(inMemDB.Index, v)
	}
	txn.Commit()
	for _, v := range tmp {
		pubkey := v.PublicKeyString()
		msg, ok := signReqs.Map[pubkey]
		if ok {
			sig := v.Sign([]byte(msg.Message)).Marshal()
			batchResp.Map[pubkey] = EthereumBLSKeySignatureResponse{bls_signer.ConvertBytesToString(sig)}
		}
	}
	if len(batchResp.Map) != len(signReqs.Map) {
		log.Ctx(ctx).Warn().Msg("SignValidatorMessagesFromInMemDb, did not contain all expected validator signatures")
	}
	return batchResp, nil
}
