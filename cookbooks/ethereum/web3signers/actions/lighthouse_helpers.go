package ethereum_web3signer_actions

import (
	"context"
	"fmt"

	"github.com/ghodss/yaml"
	signing_automation_ethereum "github.com/zeus-fyi/zeus/pkg/artemis/signing_automation/ethereum"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
)

const Web3SignerType = "web3signer"

// LighthouseWeb3SignerRequest uses 0x prefixed addresses
type LighthouseWeb3SignerRequest struct {
	Type                  string `json:"type" yaml:"type"`
	Enabled               bool   `json:"enabled" yaml:"enable"`
	Description           string `json:"description" yaml:"description"`
	SuggestedFeeRecipient string `json:"suggested_fee_recipient" yaml:"suggested_fee_recipient"`
	VotingPublicKey       string `json:"voting_public_key" yaml:"voting_public_key"`
	Graffiti              string `json:"graffiti,omitempty" yaml:"graffiti,omitempty"`
	Url                   string `json:"url,omitempty" yaml:"url,omitempty"`
	RootCertificatePath   string `json:"root_certificate_path,omitempty" yaml:"root_certificate_path,omitempty"`
	RequestTimeoutMs      int    `json:"request_timeout_ms,omitempty" yaml:"request_timeout_ms,omitempty"`
}

type LighthouseWeb3SignerRequests struct {
	Enabled       bool
	Web3SignerURL string
	FeeAddr       string
	Slice         []LighthouseWeb3SignerRequest
}

func (l *LighthouseWeb3SignerRequests) ReadDepositParamsAndExtractToEnableKeysOnWeb3Signer(ctx context.Context, dpSlice signing_automation_ethereum.ValidatorDepositSlice) {
	l.Slice = make([]LighthouseWeb3SignerRequest, len(dpSlice))
	for i, param := range dpSlice {
		l.Slice[i] = LighthouseWeb3SignerRequest{
			Enabled:               l.Enabled,
			Type:                  Web3SignerType,
			Description:           fmt.Sprintf("network: %s, fork: %s, pubkey: %s", param.NetworkName, param.ForkVersion, param.Pubkey),
			SuggestedFeeRecipient: strings_filter.AddHexPrefix(l.FeeAddr),
			VotingPublicKey:       strings_filter.AddHexPrefix(param.Pubkey),
			Url:                   l.Web3SignerURL,
		}
	}
}

func (l *LighthouseWeb3SignerRequests) WriteYamlConfig(p filepaths.Path) error {
	ymlBytes, err := yaml.Marshal(&l.Slice)
	if err != nil {
		return err
	}
	p.FnOut = "validator_definitions.yml"
	err = p.WriteToFileOutPath(ymlBytes)
	if err != nil {
		return err
	}
	return err
}
