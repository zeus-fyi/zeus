package iris_programmable_proxy

const (
	LoadBalancingStrategy = "Load-Balancing-Strategy"
)

func (i *IrisProxyRulesConfigs) AddLoadBalancingStrategyHeaderToHeaders(lbs string) {
	i.Headers[LoadBalancingStrategy] = lbs
}
