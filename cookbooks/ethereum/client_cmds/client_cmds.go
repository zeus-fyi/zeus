package client_cmds

import geth_cmds "github.com/zeus-fyi/zeus/cookbooks/ethereum/client_cmds/ai_generated/geth"

// You should verify commands give the right cli format if using the ai generated cmd structs,
// since they aren't tested with 100% coverages and the ai could have misinterpreted some of them
// you can set your own cli commands and use client_cmds_test case to print the cli it creates to verify

var GethEphemeralConfigTemplate = geth_cmds.GethCmdConfig{
	APIConfig: geth_cmds.APIConfig{
		HTTP:             true,
		HTTPAddr:         "0.0.0.0",
		HTTPPort:         8545,
		HTTPVHosts:       []string{"*"},
		HTTPCorsDomain:   []string{"*"},
		WS:               true,
		WSAddr:           "0.0.0.0",
		WSPort:           8546,
		WSOrigins:        []string{"*"},
		AuthRPCJWTSecret: "/data/jwt.hex",
		AuthRPCAddr:      "0.0.0.0",
		AuthRPCPort:      8551,
		AuthRPCVHosts:    "*",
	},
	EthereumOptions: geth_cmds.EthereumOptions{
		Datadir:   "/data",
		NetworkID: 1337531, // this changes on each reset
	},
	NetworkingOptions: geth_cmds.NetworkingOptions{
		Port:      30303,
		BootNodes: []string{"enode://0f2c301a9a3f9fa2ccfa362b79552c052905d8c2982f707f46cd29ece5a9e1c14ecd06f4ac951b228f059a43c6284a1a14fce709e8976cac93b50345218bf2e9@135.181.140.168:30343"},
	},
}
