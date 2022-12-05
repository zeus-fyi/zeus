package apollo_resp_types

import "time"

type ValidatorBalanceYieldIndex struct {
	ValidatorIndex                int `json:"validatorIndex"`
	GweiYieldOverEpochs           int `json:"gweiYieldOverEpochs"`
	GweiTotalYieldAtHigherEpoch   int `json:"gweiTotalYieldAtHigherEpoch"`
	GweiTotalBalanceAtLowerEpoch  int `json:"gweiTotalBalanceAtLowerEpoch"`
	GweiTotalBalanceAtHigherEpoch int `json:"gweiTotalBalanceAtHigherEpoch"`
}

type ValidatorBalancesSum struct {
	LowerEpoch          int                          `json:"lowerEpoch"`
	HigherEpoch         int                          `json:"higherEpoch"`
	ValidatorGweiYields []ValidatorBalanceYieldIndex `json:"validatorGweiYields"`
}

type ValidatorBalancesEpoch struct {
	ValidatorBalances []ValidatorBalanceEpoch `json:"validatorBalances"`
}

type ValidatorBalanceEpoch struct {
	Epoch                 int `json:"validatorIndex"`
	TotalBalanceGwei      int `json:"totalBalanceGwei"`
	CurrentEpochYieldGwei int `json:"currentEpochYieldGwei"`
	YieldToDateGwei       int `json:"yieldToDateGwei"`
}

type Validators []Validator

type Validator struct {
	Index                      int    `json:"index"`
	Pubkey                     string `json:"pubkey"`
	Balance                    int    `json:"balance"`
	EffectiveBalance           int    `json:"effectiveBalance"`
	ActivationEligibilityEpoch int    `json:"activationEligibilityEpoch"`
	ActivationEpoch            int    `json:"activationEpoch"`
	ExitEpoch                  int    `json:"exitEpoch"`
	WithdrawableEpoch          int    `json:"withdrawableEpoch"`
	Slashed                    bool   `json:"slashed"`
	Status                     string `json:"status"`
	WithdrawalCredentials      string `json:"withdrawalCredentials"`

	SubStatus string `json:"subStatus"`
	Network   string `json:"network,omitempty"`

	UpdatedAt time.Time `json:"updatedAt"`
}
