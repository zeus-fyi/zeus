package iris_proxy_rules_configs

const (
	LoadBalancingStrategy = "Load-Balancing-Strategy"
)

func (i *IrisProxyRulesConfigs) AddLoadBalancingStrategyHeaderToHeaders(lbs string) {
	i.Headers[LoadBalancingStrategy] = lbs
}
