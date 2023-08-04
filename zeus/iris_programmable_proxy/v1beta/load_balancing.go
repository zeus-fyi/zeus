package iris_programmable_proxy_v1_beta

const (
	LoadBalancingStrategy = "Load-Balancing-Strategy"
)

func (i *IrisProxyRulesConfigs) AddLoadBalancingStrategyHeaderToHeaders(lbs string) {
	i.Headers[LoadBalancingStrategy] = lbs
}
