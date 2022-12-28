package geth_cmds

import (
	"fmt"
	"strconv"
	"strings"
)

// API config
func gethCLI(config GethCmdConfig) string {
	cli := "geth"
	var b strings.Builder

	if config.IPCDisable {
		b.WriteString(" --ipcdisable")
	}
	if config.IPCPath != "" {
		b.WriteString(" --ipcpath=" + config.IPCPath)
	}
	if config.HTTP {
		b.WriteString(" --http")
	}
	if config.HTTPAddr != "" {
		b.WriteString(" --http.addr=" + config.HTTPAddr)
	}
	if config.HTTPPort != 0 {
		b.WriteString(" --http.port=" + strconv.Itoa(config.HTTPPort))
	}
	if config.HTTPApi != "" {
		b.WriteString(" --http.api=" + config.HTTPApi)
	}
	if config.HTTPRPCPrefix != "" {
		b.WriteString(" --http.rpcprefix=" + config.HTTPRPCPrefix)
	}
	if len(config.HTTPCorsDomain) > 0 {
		b.WriteString(" --http.corsdomain=" + strings.Join(config.HTTPCorsDomain, ","))
	}
	if len(config.HTTPVHosts) > 0 {
		b.WriteString(" --http.vhosts=" + strings.Join(config.HTTPVHosts, ","))
	}
	if config.WS {
		b.WriteString(" --ws")
	}
	if config.WSAddr != "" {
		b.WriteString(" --ws.addr=" + config.WSAddr)
	}
	if config.WSPort != 0 {
		b.WriteString(" --ws.port=" + strconv.Itoa(config.WSPort))
	}
	if config.WSApi != "" {
		b.WriteString(" --ws.api=" + config.WSApi)
	}
	if config.WSRPCPrefix != "" {
		b.WriteString(" --ws.rpcprefix=" + config.WSRPCPrefix)
	}
	if len(config.WSOrigins) > 0 {
		b.WriteString(" --ws.origins=" + strings.Join(config.WSOrigins, ","))
	}
	if config.AuthRPCJWTSecret != "" {
		b.WriteString(" --authrpc.jwtsecret=" + config.AuthRPCJWTSecret)
	}
	if config.AuthRPCAddr != "" {
		b.WriteString(" --authrpc.addr=" + config.AuthRPCAddr)
	}
	if config.AuthRPCPort != 0 {
		b.WriteString(" --authrpc.port=" + strconv.Itoa(config.AuthRPCPort))
	}
	if config.GraphQL {
		b.WriteString(" --graphql")
	}
	if len(config.GraphQLCorsDomain) > 0 {
		b.WriteString(" --graphql.corsdomain=" + strings.Join(config.GraphQLCorsDomain, ","))
	}
	if len(config.GraphQLVHosts) > 0 {
		b.WriteString(" --graphql.vhosts=" + strings.Join(config.GraphQLVHosts, ","))
	}
	if config.RPCGasCap != 0 {
		b.WriteString(" --rpc.gascap=" + strconv.FormatUint(config.RPCGasCap, 10))
	}
	if config.RPCEVMTimeout != 0 {
		b.WriteString(" --rpc.evmtimeout=" + config.RPCEVMTimeout.String())
	}
	if config.RPCTxFeeCap != 0 {
		b.WriteString(" --rpc.txfeecap=" + strconv.FormatFloat(config.RPCTxFeeCap, 'f', -1, 64))
	}
	if config.RPCAllowUnprotectedTX {
		b.WriteString(" --rpc.allow-unprotected-txs")
	}
	if config.JSPath != "" {
		b.WriteString(" --jspath=" + config.JSPath)
	}
	if config.Exec != "" {
		b.WriteString(" --exec=" + config.Exec)
	}
	if len(config.Preload) > 0 {
		b.WriteString(" --preload=" + strings.Join(config.Preload, ","))
	}

	return cli
}

// EthereumOptions config
func gethCommand(config GethCmdConfig) string {
	var b strings.Builder

	b.WriteString("geth ")
	if config.Mainnet {
		b.WriteString("--mainnet ")
	} else if config.Goerli {
		b.WriteString("--goerli ")
	} else if config.Sepolia {
		b.WriteString("--sepolia ")
	}
	if config.Config != "" {
		b.WriteString("--config ")
		b.WriteString(config.Config)
		b.WriteString(" ")
	}
	if config.DatadirMinFreeDisk > 0 {
		b.WriteString("--datadir.minfreedisk ")
		b.WriteString(strconv.Itoa(config.DatadirMinFreeDisk))
		b.WriteString(" ")
	}
	if config.Keystore != "" {
		b.WriteString("--keystore ")
		b.WriteString(config.Keystore)
		b.WriteString(" ")
	}
	if config.USB {
		b.WriteString("--usb ")
	}
	if config.PCSCDPath != "" {
		b.WriteString("--pcscdpath ")
		b.WriteString(config.PCSCDPath)
		b.WriteString(" ")
	}
	if config.NetworkID > 0 {
		b.WriteString("--networkid ")
		b.WriteString(strconv.Itoa(config.NetworkID))
		b.WriteString(" ")
	}
	if config.SyncMode != "" {
		b.WriteString("--syncmode ")
		b.WriteString(config.SyncMode)
		b.WriteString(" ")
	}
	if config.ExitWhenSynced {
		b.WriteString("--exitwhensynced ")
	}
	if config.GCMode != "" {
		b.WriteString("--gcmode ")
		b.WriteString(config.GCMode)
		b.WriteString(" ")
	}
	if config.TxLookupLimit > 0 {
		b.WriteString("--txlookuplimit ")
		b.WriteString(strconv.Itoa(config.TxLookupLimit))
		b.WriteString(" ")
	}
	if config.Ethstats != "" {
		b.WriteString("--ethstats ")
		b.WriteString(config.Ethstats)
		b.WriteString(" ")
	}
	if config.Identity != "" {
		b.WriteString("--identity ")
		b.WriteString(config.Identity)
		b.WriteString(" ")
	}
	if config.LightKDF {
		b.WriteString(" --lightkdf")
	}

	if len(config.EthRequiredBlocks) > 0 {
		b.WriteString(" --eth.requiredblocks ")
		for i, block := range config.EthRequiredBlocks {
			if i > 0 {
				b.WriteString(",")
			}
			b.WriteString(block)
		}
	}

	if config.Mainnet {
		b.WriteString(" --mainnet")
	} else if config.Goerli {
		b.WriteString(" --goerli")
	} else if config.Sepolia {
		b.WriteString(" --sepolia")
	}

	if config.Datadir != "" {
		b.WriteString(fmt.Sprintf(" --datadir=%s", config.Datadir))
	}

	if config.DatadirAncient != "" {
		b.WriteString(fmt.Sprintf(" --datadir.ancient=%s", config.DatadirAncient))
	}

	if config.RemoteDB != "" {
		b.WriteString(fmt.Sprintf(" --remotedb=%s", config.RemoteDB))
	}

	return b.String()
}
