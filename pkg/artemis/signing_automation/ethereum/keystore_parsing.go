package signing_automation_ethereum

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
)

type ValidatorDepositParams struct {
	Pubkey                string `json:"pubkey"`
	WithdrawalCredentials string `json:"withdrawal_credentials"`
	Signature             string `json:"signature"`
	DepositDataRoot       string `json:"deposit_data_root"`
}

type Keystore struct {
	ValidatorDepositParams
	Amount             int64  `json:"amount"`
	DepositMessageRoot string `json:"deposit_message_root"`
	ForkVersion        string `json:"fork_version"`
	NetworkName        string `json:"network_name"`
	DepositCliVersion  string `json:"deposit_cli_version"`
}

func ParseKeystoreJSON(ctx context.Context, p filepaths.Path) ([]Keystore, error) {
	b := p.ReadFileInPath()
	var keystoreSlice []Keystore
	err := json.Unmarshal(b, &keystoreSlice)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("ParseKeystoreJSON")
		return keystoreSlice, err
	}
	return keystoreSlice, err
}
