package iris_proxy_rules_configs

const (
	DurableExecutionID = "Durable-Execution-ID"
	RetryCount         = "Retry-Count"
	MaxRetries         = "Max-Retries"
)

func (i *IrisProxyRulesConfigs) AddDurableExecutionIDHeaderToHeaders(dexecID string) {
	i.Headers[DurableExecutionID] = dexecID
}

func (i *IrisProxyRulesConfigs) AddRetryCountHeaderToHeaders(retryID string) {
	i.Headers[RetryCount] = retryID
}

func (i *IrisProxyRulesConfigs) AddMaxRetriesHeaderToHeaders(maxRetries string) {
	i.Headers[MaxRetries] = maxRetries
}
