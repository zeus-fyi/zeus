package iris_programmable_proxy_v1_beta

const (
	LoadBalancingStrategy    = "X-Load-Balancing-Strategy"
	Adaptive                 = "Adaptive"
	AdaptiveLoadBalancingKey = "X-Adaptive-Metrics-Key"
	EthereumJsonRPC          = "Ethereum"
)

func (i *IrisProxyRulesConfigs) AddLoadBalancingStrategyHeaderToHeaders(lbs string) {
	i.Headers[LoadBalancingStrategy] = lbs
}

func (i *IrisProxyRulesConfigs) AddEthereumJsonRpcAdaptiveLoadBalancingStrategyHeaderToHeaders() {
	i.Headers[LoadBalancingStrategy] = Adaptive
	i.Headers[AdaptiveLoadBalancingKey] = EthereumJsonRPC
}

/*

EthereumJsonRPC will build stats on the payload using the method value in the JSON RPC request,
and is a reserved key-value pair for the Adaptive load balancing strategy designed for QuickNode's Marketplace.

It uses generated table stats to optimize latency and throughput for Ethereum JSON RPC requests once enough samples
have been taken.

{
	"jsonrpc": "2.0",
	"method": "eth_blockNumber", // builds stats on this method value
	"params": [],
	"id": 1
}
*/
