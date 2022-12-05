package apollo_req_types

type ValidatorBalancesRequest struct {
	ValidatorIndexes []int `json:"validatorIndexes"`
	LowerEpoch       int   `json:"lowerEpoch"`
	HigherEpoch      int   `json:"higherEpoch"`
}

type ValidatorsRequest struct {
	ValidatorIndexes []int `json:"validatorIndexes"`
}
