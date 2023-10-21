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

type ExtendedDepositParams struct {
	ValidatorDepositParams
	Amount             int    `json:"amount"`
	DepositMessageRoot string `json:"deposit_message_root"`
	ForkVersion        string `json:"fork_version"`
	NetworkName        string `json:"network_name,omitempty"`
	DepositCliVersion  string `json:"deposit_cli_version,omitempty"`
}

type ValidatorDepositSlice []ExtendedDepositParams

func ParseValidatorDepositSliceJSON(ctx context.Context, p filepaths.Path) (ValidatorDepositSlice, error) {
	var dpSlice ValidatorDepositSlice
	b, err := p.ReadFirstFileInPathWithFilter()
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("ParseValidatorDepositSliceJSON")
		return dpSlice, err
	}
	err = json.Unmarshal(b, &dpSlice)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("ParseValidatorDepositSliceJSON")
		return dpSlice, err
	}
	return dpSlice, err
}
