package signing_automation_ethereum

import (
	"encoding/hex"
	"fmt"
	"strings"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	e2types "github.com/wealdtech/go-eth2-types/v2"
	util "github.com/wealdtech/go-eth2-util"
)

type GenesisData struct {
	Data struct {
		GenesisTime           string `json:"genesis_time"`
		GenesisValidatorsRoot string `json:"genesis_validators_root"`
		GenesisForkVersion    string `json:"genesis_fork_version"`
	} `json:"data"`
}

func GetEphemeralForkVersion() (*spec.Version, error) {
	return GetForkVersion(EphemeralBeacon)
}

func GetForkVersion(beacon string) (*spec.Version, error) {
	ver := &spec.Version{}
	resp := GenesisData{}
	r := resty.New()
	r.SetBaseURL(beacon)
	_, err := r.R().
		SetResult(&resp).
		Get(BeaconGenesisPath)
	if err != nil {
		log.Err(err)
		return ver, err
	}
	forkVersion, err := hex.DecodeString(strings.TrimPrefix(resp.Data.GenesisForkVersion, "0x"))
	if err != nil {
		log.Err(err)
		return ver, err
	}

	ver[0] = forkVersion[0]
	ver[1] = forkVersion[1]
	ver[2] = forkVersion[2]
	ver[3] = forkVersion[3]
	return ver, err
}

// ValidateAndReturnEcdsaPubkeyBytes is borrowed from https://github.com/wealdtech/ethdo
func ValidateAndReturnEcdsaPubkeyBytes(ecdsaWithdrawalKey string) ([]byte, error) {
	withdrawalAddressBytes, err := hex.DecodeString(strings.TrimPrefix(ecdsaWithdrawalKey, "0x"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode withdrawal address")
	}
	if len(withdrawalAddressBytes) != 20 {
		return nil, errors.New("withdrawal address must be exactly 20 bytes in length")
	}
	// Ensure the address is properly checksummed.
	checksummedAddress := addressBytesToEIP55(withdrawalAddressBytes)
	if checksummedAddress != ecdsaWithdrawalKey {
		return nil, fmt.Errorf("withdrawal address checksum does not match (expected %s)", checksummedAddress)
	}
	withdrawalCredentials := make([]byte, 32)
	copy(withdrawalCredentials[12:32], withdrawalAddressBytes)
	// This is hard-coded, to allow deposit data to be generated without a connection to the beacon node.
	withdrawalCredentials[0] = byte(1) // ETH1_ADDRESS_WITHDRAWAL_PREFIX
	return withdrawalCredentials, err
}

// ValidateAndReturnBLSPubkeyBytes is borrowed from https://github.com/wealdtech/ethdo
func ValidateAndReturnBLSPubkeyBytes(blsPubKey string) ([]byte, error) {
	withdrawalPubKeyBytes, err := hex.DecodeString(strings.TrimPrefix(blsPubKey, "0x"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode withdrawal public key")
	}
	if len(withdrawalPubKeyBytes) != 48 {
		return nil, errors.New("withdrawal public key must be exactly 48 bytes in length")
	}
	pubKey, err := e2types.BLSPublicKeyFromBytes(withdrawalPubKeyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "withdrawal public key is not valid")
	}
	withdrawalCredentials := util.SHA256(pubKey.Marshal())
	// This is hard-coded, to allow deposit data to be generated without a connection to the beacon node.
	withdrawalCredentials[0] = byte(0) // BLS_WITHDRAWAL_PREFIX
	return withdrawalCredentials, err
}

// addressBytesToEIP55 converts a byte array in to an EIP-55 string format.
// and is borrowed from https://github.com/wealdtech/ethdo
func addressBytesToEIP55(address []byte) string {
	bytes := []byte(fmt.Sprintf("%x", address))
	hash := util.Keccak256(bytes)
	for i := 0; i < len(bytes); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte >>= 4
		} else {
			hashByte &= 0xf
		}
		if bytes[i] > '9' && hashByte > 7 {
			bytes[i] -= 32
		}
	}

	return fmt.Sprintf("0x%s", string(bytes))
}
