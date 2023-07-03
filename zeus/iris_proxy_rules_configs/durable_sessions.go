package iris_proxy_rules_configs

const (
	DurableExecutionID = "Durable-Execution-ID"
	RetryCount         = "Retry-Count"
	MaxRetries         = "Max-Retries"
)

func (i *IrisProxyRulesConfigs) AddDurableExecutionIDHeaderToHeaders(dexecID string) {
	i.Headers[DurableExecutionID] = dexecID
}

func (i *IrisProxyRulesConfigs) AddRetryCountHeaderToHeaders(dexecID string) {
	i.Headers[RetryCount] = dexecID
}

func (i *IrisProxyRulesConfigs) AddMaxRetriesHeaderToHeaders(dexecID string) {
	i.Headers[MaxRetries] = dexecID
}
