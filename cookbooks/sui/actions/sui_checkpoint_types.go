package sui_actions

type SuiCheckpointResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Epoch                      string `json:"epoch"`
		SequenceNumber             string `json:"sequenceNumber"`
		Digest                     string `json:"digest"`
		NetworkTotalTransactions   string `json:"networkTotalTransactions"`
		PreviousDigest             string `json:"previousDigest"`
		EpochRollingGasCostSummary struct {
			ComputationCost         string `json:"computationCost"`
			StorageCost             string `json:"storageCost"`
			StorageRebate           string `json:"storageRebate"`
			NonRefundableStorageFee string `json:"nonRefundableStorageFee"`
		} `json:"epochRollingGasCostSummary"`
		TimestampMs           string        `json:"timestampMs"`
		Transactions          []string      `json:"transactions"`
		CheckpointCommitments []interface{} `json:"checkpointCommitments"`
		ValidatorSignature    string        `json:"validatorSignature"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}
