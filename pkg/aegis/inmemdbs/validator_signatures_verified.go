package aegis_inmemdbs

import (
	"context"
	"encoding/hex"
	"errors"

	"github.com/rs/zerolog/log"
	types "github.com/wealdtech/go-eth2-types/v2"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

// VerifySignatures returns a slice of pubkeys that have been verified with the given signed message, it returns the pubkeys with a 0x prefix string
func (sr *EthereumBLSKeySignatureResponses) VerifySignatures(ctx context.Context, sigRequests EthereumBLSKeySignatureRequests) ([]string, error) {
	if len(sr.Map) <= 0 {
		return []string{}, nil
	}
	verifiedKeys := make([]string, len(sr.Map))
	i := 0
	for k, sigResp := range sr.Map {
		msgStr, ok := sigRequests.Map[k]
		if !ok {
			err := errors.New("pubkey not in signature request message")
			log.Ctx(ctx).Err(err)
			return []string{}, err
		}
		sigHexStr, err := hex.DecodeString(sigResp.Signature)
		if err != nil {
			log.Ctx(ctx).Err(err)
			return nil, err
		}
		sig, err := types.BLSSignatureFromBytes(sigHexStr)
		if err != nil {
			log.Ctx(ctx).Err(err)
			return nil, err
		}
		pubkeyHexStr, err := hex.DecodeString(k)
		if err != nil {
			log.Ctx(ctx).Err(err)
			return nil, err
		}
		pubkey, err := types.BLSPublicKeyFromBytes(pubkeyHexStr)
		if err != nil {
			log.Ctx(ctx).Err(err)
			return nil, err
		}
		if !sig.Verify([]byte(msgStr.Message), pubkey) {
			err = errors.New("signature does not map to expected pubkey")
			log.Ctx(ctx).Err(err)
			return []string{}, err
		}
		verifiedKeys[i] = strings_filter.AddHexPrefix(k)
		i++
	}
	return verifiedKeys, nil
}
