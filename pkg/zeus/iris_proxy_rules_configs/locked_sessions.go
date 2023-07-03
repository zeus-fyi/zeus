package iris_proxy_rules_configs

const (
	SessionLockID = "Session-Lock-ID"
)

type IrisProxyRulesConfigs struct {
	Routes  []string
	Headers map[string]string
}

func (i *IrisProxyRulesConfigs) AddSessionLockIDHeaderToHeaders(sessionID string) {
	i.Headers[SessionLockID] = sessionID
}
