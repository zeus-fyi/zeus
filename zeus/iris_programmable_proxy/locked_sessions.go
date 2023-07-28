package iris_programmable_proxy

import zeus_client "github.com/zeus-fyi/zeus/zeus/z_client"

const (
	SessionLockID  = "Session-Lock-ID"
	SessionLockTTL = "Session-Lock-TTL"
)

type IrisProxyRulesConfigs struct {
	zeus_client.ZeusClient
	Routes  []string          `json:"routes"`
	Headers map[string]string `json:"headers"`
}

func (i *IrisProxyRulesConfigs) AddSessionLockIDHeaderToHeaders(sessionID string) {
	i.Headers[SessionLockID] = sessionID
}

func (i *IrisProxyRulesConfigs) AddSessionLockTTLHeaderToHeaders(ttl string) {
	i.Headers[SessionLockTTL] = ttl
}
