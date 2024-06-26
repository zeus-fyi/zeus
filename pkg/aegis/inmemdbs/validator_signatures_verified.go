package aegis_inmemdbs

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/rs/zerolog/log"
	types "github.com/wealdtech/go-eth2-types/v2"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

// VerifySignatures returns a slice of pubkeys that have been verified with the given signed message, it returns the pubkeys with a 0x prefix string
// yous should only use this to verify hex payloads, normal strings likely won't work
func (sr *EthereumBLSKeySignatureResponses) VerifySignatures(ctx context.Context, sigRequests EthereumBLSKeySignatureRequests, isHexPayload bool) ([]string, error) {
	if len(sr.Map) <= 0 {
		return []string{}, nil
	}
	verifiedKeys := make([]string, len(sr.Map))
	i := 0
	for pubkeySigner, sigResp := range sr.Map {
		msgStr, ok := sigRequests.Map[pubkeySigner]
		if !ok {
			err := errors.New("pubkey not in signature request message")
			log.Ctx(ctx).Err(err)
			return []string{}, err
		}
		sigHexDecode, err := hex.DecodeString(strings_filter.Trim0xPrefix(sigResp.Signature))
		if err != nil {
			log.Ctx(ctx).Err(err)
			return nil, err
		}
		sig, err := types.BLSSignatureFromBytes(sigHexDecode)
		if err != nil {
			log.Ctx(ctx).Err(err)
			return nil, err
		}
		pubkeyHexStr, err := hex.DecodeString(strings_filter.Trim0xPrefix(pubkeySigner))
		if err != nil {
			log.Ctx(ctx).Err(err)
			return nil, err
		}
		pubkey, err := types.BLSPublicKeyFromBytes(pubkeyHexStr)
		if err != nil {
			log.Ctx(ctx).Err(err)
			return nil, err
		}

		if isHexPayload {
			b, berr := hex.DecodeString(strings_filter.Trim0xPrefix(msgStr.Message))
			if berr != nil {
				log.Ctx(ctx).Err(berr).Msg("failed to decode hex payload")
				return nil, berr
			}
			if !sig.Verify(b, pubkey) {
				err = errors.New("signature does not map to expected pubkey")
				log.Ctx(ctx).Err(err)
				return []string{}, err
			}
		} else {
			if !sig.Verify([]byte(msgStr.Message), pubkey) {
				err = errors.New("signature does not map to expected pubkey")
				log.Ctx(ctx).Err(err)
				return []string{}, err
			}
		}
		verifiedKeys[i] = pubkeySigner
		i++
	}
	return verifiedKeys, nil
}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
