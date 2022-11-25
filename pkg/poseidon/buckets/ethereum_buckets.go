package buckets

import "github.com/zeus-fyi/zeus/pkg/poseidon"

var GethMainnetBucket = poseidon.BucketRequest{
	BucketName: "zeus-fyi-ethereum",
	Protocol:   "ethereum",
	Network:    "mainnet",
	ClientType: "exec.client.standard",
	ClientName: "geth",
}

var LighthouseMainnetBucket = poseidon.BucketRequest{
	BucketName: "zeus-fyi-ethereum",
	Protocol:   "ethereum",
	Network:    "mainnet",
	ClientType: "consensus.client.standard",
	ClientName: "lighthouse",
}
