package signing_automation_ethereum

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	spec "github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	e2types "github.com/wealdtech/go-eth2-types/v2"
	util "github.com/wealdtech/go-eth2-util"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

type GenesisData struct {
	Data struct {
		GenesisTime           string `json:"genesis_time"`
		GenesisValidatorsRoot string `json:"genesis_validators_root"`
		GenesisForkVersion    string `json:"genesis_fork_version"`
	} `json:"data"`
}

type ForkData struct {
	ExecutionOptimistic bool `json:"execution_optimistic"`
	Finalized           bool `json:"finalized"`
	Data                struct {
		PreviousVersion string `json:"previous_version"`
		CurrentVersion  string `json:"current_version"`
		Epoch           string `json:"epoch"`
	} `json:"data"`
}

type ValidatorInfo struct {
	ExecutionOptimistic bool `json:"execution_optimistic"`
	Finalized           bool `json:"finalized"`
	Data                struct {
		Index     string `json:"index"`
		Balance   string `json:"balance"`
		Status    string `json:"status"`
		Validator struct {
			Pubkey                     string `json:"pubkey"`
			WithdrawalCredentials      string `json:"withdrawal_credentials"`
			EffectiveBalance           string `json:"effective_balance"`
			Slashed                    bool   `json:"slashed"`
			ActivationEligibilityEpoch string `json:"activation_eligibility_epoch"`
			ActivationEpoch            string `json:"activation_epoch"`
			ExitEpoch                  string `json:"exit_epoch"`
			WithdrawableEpoch          string `json:"withdrawable_epoch"`
		} `json:"validator"`
	} `json:"data"`
}

func GetEphemeralForkVersion(ctx context.Context) (*spec.Version, error) {
	return GetForkVersion(ctx, EphemeralBeacon)
}

func GetForkVersion(ctx context.Context, beacon string) (*spec.Version, error) {
	ver := &spec.Version{}
	resp := GenesisData{}
	r := resty.New()
	r.SetBaseURL(beacon)
	_, err := r.R().
		SetResult(&resp).
		Get(BeaconGenesisPath)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return ver, err
	}
	forkVersion, err := hex.DecodeString(strings.TrimPrefix(resp.Data.GenesisForkVersion, "0x"))
	if err != nil {
		log.Ctx(ctx).Err(err)
		return ver, err
	}

	ver[0] = forkVersion[0]
	ver[1] = forkVersion[1]
	ver[2] = forkVersion[2]
	ver[3] = forkVersion[3]
	return ver, err
}

func GetCurrentForkVersion(ctx context.Context, beacon string) (*spec.Version, error) {
	ver := &spec.Version{}
	resp := ForkData{}
	r := resty.New()
	r.SetBaseURL(beacon)
	_, err := r.R().
		SetResult(&resp).
		Get(BeaconForkPath)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return ver, err
	}
	forkVersion, err := hex.DecodeString(strings.TrimPrefix(resp.Data.CurrentVersion, "0x"))
	if err != nil {
		log.Ctx(ctx).Err(err)
		return ver, err
	}

	ver[0] = forkVersion[0]
	ver[1] = forkVersion[1]
	ver[2] = forkVersion[2]
	ver[3] = forkVersion[3]
	return ver, err
}

func GetValidatorIndexFromPubkey(ctx context.Context, beacon, pubkey string) (string, error) {
	resp := ValidatorInfo{}
	pubkey = strings_filter.AddHexPrefix(pubkey)
	urlPath := fmt.Sprintf("/eth/v1/beacon/states/head/validators/%s", pubkey)
	r := resty.New()
	r.SetBaseURL(beacon)
	_, err := r.R().
		SetResult(&resp).
		Get(urlPath)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return resp.Data.Index, err
	}
	return resp.Data.Index, err
}

func SubmitVoluntaryExit(ctx context.Context, beacon string, ve *spec.SignedVoluntaryExit) error {
	r := resty.New()
	r.SetBaseURL(beacon)
	resp, err := r.R().
		SetBody(ve).
		Post(BeaconVoluntaryExitPath)
	if err != nil {
		log.Ctx(ctx).Err(err)
		return err
	}
	if resp.StatusCode() == 400 {
		log.Ctx(ctx).Err(err)
		return errors.New("invalid voluntary exit format")
	}
	if resp.StatusCode() != 200 {
		log.Ctx(ctx).Err(err)
		return errors.New("failed to submit voluntary exit")
	}
	return err
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
