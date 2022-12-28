package geth_cmds

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func (g *GethCmdConfig) BuildCliCmd() string {
	apiCmd := g.APIConfig.BuildCliCmd()
	ethOptCmd := g.EthereumOptions.BuildCliCmd()
	lightOptCmd := g.LightClientOptions.BuildCliCmd()
	accountOptCmd := g.AccountOptions.BuildCliCmd()
	txPoolOptCmd := g.TransactionPoolOptions.BuildCliCmd()
	perfTuningOptCmd := g.PerformanceTuningOptions.BuildCliCmd()
	networkOptCmd := g.NetworkingOptions.BuildCliCmd()
	gasPriceOptCmd := g.GasPriceOracleOptions.BuildCliCmd()
	loggingOptCmd := g.LoggingOptions.BuildCliCmd()
	metricsOptCmd := g.MetricsOptions.BuildCliCmd()
	optCmd := g.Options.BuildCliCmd()
	cmd := "geth" + apiCmd + ethOptCmd + lightOptCmd + accountOptCmd +
		txPoolOptCmd + perfTuningOptCmd + networkOptCmd +
		gasPriceOptCmd + loggingOptCmd + metricsOptCmd +
		optCmd
	return cmd
}

func (a *APIConfig) BuildCliCmd() string {
	var b strings.Builder

	if a.IPCDisable {
		b.WriteString(" --ipcdisable")
	}
	if a.IPCPath != "" {
		b.WriteString(" --ipcpath=" + a.IPCPath)
	}
	if a.HTTP {
		b.WriteString(" --http")
	}
	if a.HTTPAddr != "" {
		b.WriteString(" --http.addr=" + a.HTTPAddr)
	}
	if a.HTTPPort != 0 {
		b.WriteString(" --http.port=" + strconv.Itoa(a.HTTPPort))
	}
	if a.HTTPApi != "" {
		b.WriteString(" --http.api=" + a.HTTPApi)
	}
	if a.HTTPRPCPrefix != "" {
		b.WriteString(" --http.rpcprefix=" + a.HTTPRPCPrefix)
	}
	if len(a.HTTPCorsDomain) > 0 {
		b.WriteString(" --http.corsdomain=" + strings.Join(a.HTTPCorsDomain, ","))
	}
	if len(a.HTTPVHosts) > 0 {
		b.WriteString(" --http.vhosts=" + strings.Join(a.HTTPVHosts, ","))
	}
	if a.WS {
		b.WriteString(" --ws")
	}
	if a.WSAddr != "" {
		b.WriteString(" --ws.addr=" + a.WSAddr)
	}
	if a.WSPort != 0 {
		b.WriteString(" --ws.port=" + strconv.Itoa(a.WSPort))
	}
	if a.WSApi != "" {
		b.WriteString(" --ws.api=" + a.WSApi)
	}
	if a.WSRPCPrefix != "" {
		b.WriteString(" --ws.rpcprefix=" + a.WSRPCPrefix)
	}
	if len(a.WSOrigins) > 0 {
		b.WriteString(" --ws.origins=" + strings.Join(a.WSOrigins, ","))
	}
	if a.AuthRPCJWTSecret != "" {
		b.WriteString(" --authrpc.jwtsecret=" + a.AuthRPCJWTSecret)
	}
	if a.AuthRPCAddr != "" {
		b.WriteString(" --authrpc.addr=" + a.AuthRPCAddr)
	}
	if a.AuthRPCPort != 0 {
		b.WriteString(" --authrpc.port=" + strconv.Itoa(a.AuthRPCPort))
	}
	if a.GraphQL {
		b.WriteString(" --graphql")
	}
	if len(a.GraphQLCorsDomain) > 0 {
		b.WriteString(" --graphql.corsdomain=" + strings.Join(a.GraphQLCorsDomain, ","))
	}
	if len(a.GraphQLVHosts) > 0 {
		b.WriteString(" --graphql.vhosts=" + strings.Join(a.GraphQLVHosts, ","))
	}
	if a.RPCGasCap != 0 {
		b.WriteString(" --rpc.gascap=" + strconv.FormatUint(a.RPCGasCap, 10))
	}
	if a.RPCEVMTimeout != 0 {
		b.WriteString(" --rpc.evmtimeout=" + a.RPCEVMTimeout.String())
	}
	if a.RPCTxFeeCap != 0 {
		b.WriteString(" --rpc.txfeecap=" + strconv.FormatFloat(a.RPCTxFeeCap, 'f', -1, 64))
	}
	if a.RPCAllowUnprotectedTX {
		b.WriteString(" --rpc.allow-unprotected-txs")
	}
	if a.JSPath != "" {
		b.WriteString(" --jspath=" + a.JSPath)
	}
	if a.Exec != "" {
		b.WriteString(" --exec=" + a.Exec)
	}
	if len(a.Preload) > 0 {
		b.WriteString(" --preload=" + strings.Join(a.Preload, ","))
	}
	return b.String()
}
func (eo *EthereumOptions) BuildCliCmd() string {
	var b strings.Builder

	if eo.Mainnet {
		b.WriteString("--mainnet ")
	} else if eo.Goerli {
		b.WriteString("--goerli ")
	} else if eo.Sepolia {
		b.WriteString("--sepolia ")
	}
	if eo.Config != "" {
		b.WriteString("--config ")
		b.WriteString(eo.Config)
	}
	if eo.DatadirMinFreeDisk > 0 {
		b.WriteString("--datadir.minfreedisk ")
		b.WriteString(strconv.Itoa(eo.DatadirMinFreeDisk))
	}
	if eo.Keystore != "" {
		b.WriteString("--keystore ")
		b.WriteString(eo.Keystore)
	}
	if eo.USB {
		b.WriteString("--usb ")
	}
	if eo.PCSCDPath != "" {
		b.WriteString(" --pcscdpath")
		b.WriteString(eo.PCSCDPath)
	}
	if eo.NetworkID > 0 {
		b.WriteString(fmt.Sprintf(" --networkid=%d", eo.NetworkID))
	}
	if eo.SyncMode != "" {
		b.WriteString(" --syncmode")
		b.WriteString(eo.SyncMode)
	}
	if eo.ExitWhenSynced {
		b.WriteString(" --exitwhensynced")
	}
	if eo.GCMode != "" {
		b.WriteString(" --gcmode ")
		b.WriteString(eo.GCMode)
	}
	if eo.TxLookupLimit > 0 {
		b.WriteString(" --txlookuplimit")
		b.WriteString(strconv.Itoa(eo.TxLookupLimit))
	}
	if eo.Ethstats != "" {
		b.WriteString(" --ethstats")
		b.WriteString(eo.Ethstats)
	}
	if eo.Identity != "" {
		b.WriteString(" --identity")
		b.WriteString(eo.Identity)
	}
	if eo.LightKDF {
		b.WriteString(" --lightkdf")
	}
	if len(eo.EthRequiredBlocks) > 0 {
		b.WriteString(" --eth.requiredblocks ")
		for i, block := range eo.EthRequiredBlocks {
			if i > 0 {
				b.WriteString(",")
			}
			b.WriteString(block)
		}
	}
	if eo.Mainnet {
		b.WriteString(" --mainnet")
	} else if eo.Goerli {
		b.WriteString(" --goerli")
	} else if eo.Sepolia {
		b.WriteString(" --sepolia")
	}

	if eo.Datadir != "" {
		b.WriteString(fmt.Sprintf(" --datadir=%s", eo.Datadir))
	}

	if eo.DatadirAncient != "" {
		b.WriteString(fmt.Sprintf(" --datadir.ancient=%s", eo.DatadirAncient))
	}

	if eo.RemoteDB != "" {
		b.WriteString(fmt.Sprintf(" --remotedb=%s", eo.RemoteDB))
	}
	return b.String()
}
func (lco *LightClientOptions) BuildCliCmd() string {
	b := strings.Builder{}

	if lco.Serve > 0 {
		b.WriteString(" --light.serve=")
		b.WriteString(strconv.Itoa(lco.Serve))
	}
	if lco.Ingress > 0 {
		b.WriteString(" --light.ingress=")
		b.WriteString(strconv.Itoa(lco.Ingress))
	}

	if lco.Egress > 0 {
		b.WriteString(" --light.egress=")
		b.WriteString(strconv.Itoa(lco.Egress))
	}

	if lco.MaxPeers > 0 {
		b.WriteString(" --light.maxpeers=")
		b.WriteString(strconv.Itoa(lco.MaxPeers))
	}
	if len(lco.ULCServers) > 0 {
		b.WriteString(" --ulc.servers=" + strings.Join(lco.ULCServers, ","))
	}
	if lco.ULCFraction > 0 {
		b.WriteString(" --ulc.fraction=")
		b.WriteString(strconv.Itoa(lco.ULCFraction))
		b.WriteString(" ")
	}
	if lco.ULCOnlyAnnounce {
		b.WriteString(" --ulc.onlyannouce=")
		b.WriteString(strconv.FormatBool(lco.ULCOnlyAnnounce))
	}
	if lco.NoPruning {
		b.WriteString(" --light.nopruning=")
		b.WriteString(strconv.FormatBool(lco.NoPruning))
	}
	if lco.NoSyncServe {
		b.WriteString(" --light.nosyncserve=")
		b.WriteString(strconv.FormatBool(lco.NoSyncServe))
	}
	return b.String()
}
func (ao *AccountOptions) BuildCliCmd() string {
	b := &strings.Builder{}
	if len(ao.Unlock) > 0 {
		b.WriteString(`--unlock "`)
		for _, acc := range ao.Unlock {
			b.WriteString(acc)
			b.WriteString(`,`)
		}
		b.WriteString(`" `)
	}
	if ao.Password != "" {
		b.WriteString(`--password "`)
		b.WriteString(ao.Password)
		b.WriteString(`" `)
	}
	if ao.Signer != "" {
		b.WriteString(`--signer "`)
		b.WriteString(ao.Signer)
		b.WriteString(`" `)
	}
	if ao.AllowUnlock {
		b.WriteString(`--allow-insecure-unlock `)
	}
	return b.String()
}
func (tpo *TransactionPoolOptions) BuildCliCmd() string {
	b := strings.Builder{}
	if len(tpo.Locals) > 0 {
		b.WriteString(" --txpool.locals=" + strings.Join(tpo.Locals, ","))
	}
	if tpo.NoLocals {
		b.WriteString(" --txpool.nolocals=" + strconv.FormatBool(tpo.NoLocals))
	}
	if tpo.Journal != "" {
		b.WriteString(" --txpool.journal=" + tpo.Journal)
	}
	if tpo.Rejournal != 0 {
		b.WriteString(" --txpool.rejournal=" + tpo.Rejournal.String())
	}
	if tpo.PriceLimit != 0 {
		b.WriteString(" --txpool.pricelimit=" + strconv.Itoa(tpo.PriceLimit))
	}
	if tpo.PriceBump != 0 {
		b.WriteString(" --txpool.pricebump=" + strconv.Itoa(tpo.PriceBump))
	}
	if tpo.AccountSlots != 0 {
		b.WriteString(" --txpool.accountslots=" + strconv.Itoa(tpo.AccountSlots))
	}
	if tpo.GlobalSlots != 0 {
		b.WriteString(" --txpool.globalslots=" + strconv.Itoa(tpo.GlobalSlots))
	}
	if tpo.AccountQueue != 0 {
		b.WriteString(" --txpool.accountqueue=" + strconv.Itoa(tpo.AccountQueue))
	}
	if tpo.GlobalQueue != 0 {
		b.WriteString(" --txpool.globalqueue=" + strconv.Itoa(tpo.GlobalQueue))
	}
	if tpo.Lifetime != 0 {
		b.WriteString(" --txpool.lifetime=" + tpo.Lifetime.String())
	}
	return b.String()
}
func (pto *PerformanceTuningOptions) BuildCliCmd() string {
	var sb strings.Builder
	if pto.Cache != 0 {
		sb.WriteString("--cache=")
		sb.WriteString(strconv.Itoa(pto.Cache))
	}
	if pto.Database != 0 {
		sb.WriteString(" --cache.database=")
		sb.WriteString(strconv.Itoa(pto.Database))
	}
	if pto.Trie != 0 {
		sb.WriteString(" --cache.trie=")
		sb.WriteString(strconv.Itoa(pto.Trie))
	}
	if pto.TrieJournal != "" {
		sb.WriteString(" --cache.trie.journal=")
		sb.WriteString(pto.TrieJournal)
	}
	if pto.TrieRejournal != nil {
		sb.WriteString(" --cache.trie.rejournal=")
		sb.WriteString(pto.TrieRejournal.String())
	}
	if pto.GC != 0 {
		sb.WriteString(" --cache.gc=")
		sb.WriteString(strconv.Itoa(pto.GC))
	}
	if pto.Snapshot != 0 {
		sb.WriteString(" --cache.snapshot=")
		sb.WriteString(strconv.Itoa(pto.Snapshot))
	}
	if pto.NoPrefetch {
		sb.WriteString(" --cache.noprefetch=")
		sb.WriteString(strconv.FormatBool(pto.NoPrefetch))
	}
	if pto.Preimages {
		sb.WriteString(" --cache.preimages=")
		sb.WriteString(strconv.FormatBool(pto.Preimages))
	}
	if pto.FDLimit != 0 {
		sb.WriteString(" --fdlimit=")
		sb.WriteString(strconv.Itoa(pto.FDLimit))
	}
	return sb.String()
}
func (no *NetworkingOptions) BuildCliCmd() string {
	b := strings.Builder{}
	if no.Port > 0 {
		b.WriteString(fmt.Sprintf(" --port=%d", no.Port))
	}
	if len(no.BootNodes) > 0 {
		b.WriteString(" --bootnodes=")
	}
	for i, bootnode := range no.BootNodes {
		if i == 0 {
			//if this is the first bootnode in the list
			b.WriteString(bootnode)
			continue
		}
		b.WriteString("," + bootnode)
	}
	if len(no.DNS) > 0 {
		b.WriteString(" --discovery.dns=")
	}
	for i, dns := range no.DNS {
		if i == 0 {
			//if this is the first DNS value in the list
			b.WriteString(dns)
			continue
		}
		b.WriteString("," + dns)
	}

	if no.MaxPeers > 0 {
		b.WriteString(fmt.Sprintf(" --maxpeers=%d", no.MaxPeers))
	}
	if no.MaxPendPeers > 0 {
		b.WriteString(fmt.Sprintf(" --maxpendpeers=%d", no.MaxPendPeers))
	}
	if len(no.NAT) > 0 {
		b.WriteString(fmt.Sprintf(" --nat=%s", no.NAT))
	}
	if no.NoDiscover {
		b.WriteString(" --nodiscover")
	}
	if no.V5Disc {
		b.WriteString(" --v5disc")
	}
	if len(no.NetRestrict) > 0 {
		b.WriteString(" --netrestrict=")
	}
	for i, restrict := range no.NetRestrict {
		if i == 0 {
			//if this is the first net restrict value in the list
			b.WriteString(restrict)
			continue
		}
		b.WriteString("," + restrict)
	}
	if no.NodeKey != "" {
		b.WriteString(" --nodekey=" + no.NodeKey)
	}
	if no.NodeKeyHex != "" {
		b.WriteString(" --nodekeyhex=" + no.NodeKeyHex)
	}
	return b.String()
}
func (gpo *GasPriceOracleOptions) BuildCliCmd() string {
	b := strings.Builder{}
	if gpo.Blocks != 0 {
		b.WriteString("--gpo.blocks=")
		b.WriteString(strconv.Itoa(gpo.Blocks))
	}
	if gpo.Percentile != 0 {
		b.WriteString(" --gpo.percentile=")
		b.WriteString(strconv.Itoa(gpo.Percentile))
	}
	if gpo.MaxPrice != 0 {
		b.WriteString(" --gpo.maxprice=")
		b.WriteString(strconv.Itoa(gpo.MaxPrice))
	}
	if gpo.IgnorePrice != 0 {
		b.WriteString(" --gpo.ignoreprice=")
		b.WriteString(strconv.Itoa(gpo.IgnorePrice))
	}
	return b.String()
}
func (lo *LoggingOptions) BuildCliCmd() string {
	cmd := bytes.Buffer{}
	if lo.FakePow {
		cmd.WriteString(" --fakepow")
	}
	if lo.NoCompaction {
		cmd.WriteString(" --nocompaction")
	}
	if lo.VerbosityEnabled {
		cmd.WriteString(fmt.Sprintf(" --verbosity=%d", lo.Verbosity))
	}
	if lo.VModule != "" {
		cmd.WriteString(fmt.Sprintf(" --vmodule=%s", lo.VModule))
	}
	if lo.JSON {
		cmd.WriteString(" --log.json")
	}
	if lo.Backtrace != "" {
		cmd.WriteString(fmt.Sprintf(" --log.backtrace=%s", lo.Backtrace))
	}
	if lo.Debug {
		cmd.WriteString(" --log.debug")
	}
	if lo.Pprof {
		cmd.WriteString(" --pprof")
		cmd.WriteString(fmt.Sprintf(" --pprof.addr=%s", lo.PprofAddr))
		cmd.WriteString(fmt.Sprintf(" --pprof.port=%d", lo.PprofPort))
		cmd.WriteString(fmt.Sprintf(" --pprof.memprofilerate=%d", lo.MemProfileRate))
		cmd.WriteString(fmt.Sprintf(" --pprof.blockprofilerate=%d", lo.BlockProfileRate))
		if lo.CPUProfile != "" {
			cmd.WriteString(fmt.Sprintf(" --pprof.cpuprofile=%s", lo.CPUProfile))
		}
	}
	if lo.Trace != "" {
		cmd.WriteString(fmt.Sprintf(" --trace=%s", lo.Trace))
	}
	return cmd.String()
}
func (mo *MetricsOptions) BuildCliCmd() string {
	var buffer bytes.Buffer
	if mo.Enabled {
		buffer.WriteString(" --metrics")
		buffer.WriteString(" --metrics.expensive")
		if mo.Expensive {
			buffer.WriteString(fmt.Sprintf(" --metrics.addr %s", mo.Addr))
			buffer.WriteString(fmt.Sprintf(" --metrics.port %d", mo.Port))

			if mo.InfluxDB {
				buffer.WriteString(fmt.Sprintf(" --metrics.influxdb.endpoint %s", mo.InfluxDBEndpoint))
				buffer.WriteString(fmt.Sprintf(" --metrics.influxdb.database %s", mo.InfluxDBDatabase))
				buffer.WriteString(fmt.Sprintf(" --metrics.influxdb.username %s", mo.InfluxDBUsername))
				buffer.WriteString(fmt.Sprintf(" --metrics.influxdb.password %s", mo.InfluxDBPassword))
				buffer.WriteString(fmt.Sprintf(" --metrics.influxdb.tags %s ", mo.InfluxDBTags))
			}

			if mo.InfluxDBv2 {
				buffer.WriteString(fmt.Sprintf(" --metrics.influxdbv2 --metrics.influxdb.token %s", mo.InfluxDBv2Token))
				buffer.WriteString(fmt.Sprintf(" --metrics.influxdb.bucket %s", mo.Bucket))
				buffer.WriteString(fmt.Sprintf(" --metrics.influxdb.organization %s", mo.Organization))
			}
		}
	}
	return buffer.String()
}
func (o *Options) BuildCliCmd() string {
	var b strings.Builder
	if o.Snapshot {
		b.WriteString(" --snapshot")
	}
	if o.BloomFilterSize > 0 {
		b.WriteString(" --bloomfilter.size=")
		b.WriteString(strconv.Itoa(o.BloomFilterSize))
	}
	if o.IgnoreLegacyReceipt {
		b.WriteString(" --ignore-legacy-receipts")
	}
	return b.String()
}
