package geth_cmds

import "time"

type GethCmdConfig struct {
	APIConfig
	EthereumOptions
	LightClientOptions
	AccountOptions
	TransactionPoolOptions
	PerformanceTuningOptions
	NetworkingOptions
	GasPriceOracleOptions
	LoggingOptions
	MetricsOptions
	Options
}

// APIConfig holds the configuration options for the api CLI commands.
type APIConfig struct {
	IPCDisable            bool          `long:"ipcdisable" description:"Disable the IPC-RPC server"`
	IPCPath               string        `long:"ipcpath" description:"Filename for IPC socket/pipe within the datadir (explicit paths escape it)"`
	HTTP                  bool          `long:"http" description:"Enable the HTTP-RPC server"`
	HTTPAddr              string        `long:"http.addr" default:"localhost" description:"HTTP-RPC server listening interface"`
	HTTPPort              int           `long:"http.port" default:"8545" description:"HTTP-RPC server listening port"`
	HTTPApi               string        `long:"http.api" description:"API's offered over the HTTP-RPC interface"`
	HTTPRPCPrefix         string        `long:"http.rpcprefix" default:"/" description:"HTTP path path prefix on which JSON-RPC is served. Use '/' to serve on all paths."`
	HTTPCorsDomain        []string      `long:"http.corsdomain" description:"Comma separated list of domains from which to accept cross origin requests (browser enforced)"`
	HTTPVHosts            []string      `long:"http.vhosts" default:"localhost" description:"Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts '*' wildcard."`
	WS                    bool          `long:"ws" description:"Enable the WS-RPC server"`
	WSAddr                string        `long:"ws.addr" default:"localhost" description:"WS-RPC server listening interface"`
	WSPort                int           `long:"ws.port" default:"8546" description:"WS-RPC server listening port"`
	WSApi                 string        `long:"ws.api" description:"API's offered over the WS-RPC interface"`
	WSRPCPrefix           string        `long:"ws.rpcprefix" default:"/" description:"HTTP path prefix on which JSON-RPC is served. Use '/' to serve on all paths."`
	WSOrigins             []string      `long:"ws.origins" description:"Origins from which to accept websockets requests"`
	AuthRPCJWTSecret      string        `long:"authrpc.jwtsecret" description:"Path to a JWT secret to use for authenticated RPC endpoints"`
	AuthRPCAddr           string        `long:"authrpc.addr" default:"localhost" description:"Listening address for authenticated APIs"`
	AuthRPCPort           int           `long:"authrpc.port" default:"8551" description:"Listening port for authenticated APIs"`
	AuthRPCVHosts         string        `long:"authrpc.vhosts" default:"localhost" description:"Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts '*' wildcard."`
	GraphQL               bool          `long:"graphql" description:"Enable GraphQL on the HTTP-RPC server. Note that GraphQL can only be started if an HTTP server is started as well."`
	GraphQLCorsDomain     []string      `long:"graphql.corsdomain" description:"Comma separated list of domains from which to accept cross origin requests (browser enforced)"`
	GraphQLVHosts         []string      `long:"graphql.vhosts" default:"localhost" description:"Comma separated list of virtual hostnames from which to accept requests (server enforced). Accepts '*' wildcard."`
	RPCGasCap             uint64        `long:"rpc.gascap" default:"50000000" description:"Sets a cap on gas that can be used in eth_call/estimateGas (0=infinite)"`
	RPCEVMTimeout         time.Duration `long:"rpc.evmtimeout" default:"5s" description:"Sets a timeout used for eth_call (0=infinite)"`
	RPCTxFeeCap           float64       `long:"rpc.txfeecap" default:"1" description:"Sets a cap on transaction fee (in ether) that can be sent via the RPC APIs (0 = no cap)"`
	RPCAllowUnprotectedTX bool          `long:"rpc.allow-unprotected-txs" description:"Allow for unprotected (non EIP155 signed) transactions to be submitted via RPC"`
	JSPath                string        `long:"jspath" default:"." description:"JavaScript root path for loadScript"`
	Exec                  string        `long:"exec" description:"Execute JavaScript statement"`
	Preload               []string      `long:"preload" description:"Comma separated list of JavaScript files to preload into the console"`
}

type EthereumOptions struct {
	Config             string   `long:"config" description:"TOML configuration file"`
	DatadirMinFreeDisk int      `long:"datadir.minfreedisk" description:"Minimum free disk space in MB, once reached triggers auto shut down (default = --cache.gc converted to MB, 0 = disabled)"`
	Keystore           string   `long:"keystore" description:"Directory for the keystore (default = inside the datadir)"`
	USB                bool     `long:"usb" description:"Enable monitoring and management of USB hardware wallets"`
	PCSCDPath          string   `long:"pcscdpath" description:"Path to the smartcard daemon (pcscd) socket file"`
	NetworkID          int      `long:"networkid" description:"Explicitly set network id (integer)(For testnets: use --sepolia, --goerli instead) (default: 1)"`
	SyncMode           string   `long:"syncmode" description:"Blockchain sync mode ('snap', 'full' or 'light') (default: snap)"`
	ExitWhenSynced     bool     `long:"exitwhensynced" description:"Exits after block synchronisation completes"`
	GCMode             string   `long:"gcmode" description:"Blockchain garbage collection mode ('full', 'archive') (default: 'full')"`
	TxLookupLimit      int      `long:"txlookuplimit" description:"Number of recent blocks to maintain transactions index for (default = about one year, 0 = entire chain) (default: 2350000)"`
	Ethstats           string   `long:"ethstats" description:"Reporting URL of a ethstats service (nodename:secret@host:port)"`
	Identity           string   `long:"identity" description:"Custom node name"`
	LightKDF           bool     `long:"lightkdf" description:"Reduce key-derivation RAM & CPU usage at some expense of KDF strength"`
	EthRequiredBlocks  []string `long:"eth.requiredblocks" description:"Comma separated block number-to-hash mappings to require for peering (<number>=<hash>)"`
	Mainnet            bool     `long:"mainnet" description:"Ethereum mainnet"`
	Goerli             bool     `long:"goerli" description:"Görli network: pre-configured proof-of-authority test network"`
	Sepolia            bool     `long:"sepolia" description:"Sepolia network: pre-configured proof-of-work test network"`
	Datadir            string   `long:"datadir" description:"Data directory for the databases and keystore (default: '~/.ethereum')"`
	DatadirAncient     string   `long:"datadir.ancient" description:"Data directory for ancient chain segments (default = inside chaindata)"`
	RemoteDB           string   `long:"remotedb" description:"URL for remote database"`
}

type LightClientOptions struct {
	Serve           int      `long:"light.serve" description:"Maximum percentage of time allowed for serving LES requests (multi-threaded processing allows values over 100) (default: 0)"`
	Ingress         int      `long:"light.ingress" description:"Incoming bandwidth limit for serving light clients (kilobytes/sec, 0 = unlimited) (default: 0)"`
	Egress          int      `long:"light.egress" description:"Outgoing bandwidth limit for serving light clients (kilobytes/sec, 0 = unlimited) (default: 0)"`
	MaxPeers        int      `long:"light.maxpeers" description:"Maximum number of light clients to serve, or light servers to attach to (default: 100)"`
	ULCServers      []string `long:"ulc.servers" description:"List of trusted ultra-light servers"`
	ULCFraction     int      `long:"ulc.fraction" description:"Minimum % of trusted ultra-light servers required to announce a new head (default: 75)"`
	ULCOnlyAnnounce bool     `long:"ulc.onlyannounce" description:"Ultra light server sends announcements only"`
	NoPruning       bool     `long:"light.nopruning" description:"Disable ancient light chain data pruning"`
	NoSyncServe     bool     `long:"light.nosyncserve" description:"Enables serving light clients before syncing"`
}

type AccountOptions struct {
	Unlock      []string `long:"unlock" default:"" description:"Comma separated list of accounts to unlock"`
	Password    string   `long:"password" default:"" description:"Password file to use for non-interactive password input"`
	Signer      string   `long:"signer" default:"" description:"External signer (url or path to ipc file)"`
	AllowUnlock bool     `long:"allow-insecure-unlock" default:"false" description:"Allow insecure account unlocking when account-related RPCs are exposed by http"`
}

type TransactionPoolOptions struct {
	Locals       []string      `long:"txpool.locals" default:"" description:"Comma separated accounts to treat as locals (no flush, priority inclusion)"`
	NoLocals     bool          `long:"txpool.nolocals" default:"false" description:"Disables price exemptions for locally submitted transactions"`
	Journal      string        `long:"txpool.journal" default:"transactions.rlp" description:"Disk journal for local transaction to survive node restarts"`
	Rejournal    time.Duration `long:"txpool.rejournal" default:"1h0m0s" description:"Time interval to regenerate the local transaction journal"`
	PriceLimit   int           `long:"txpool.pricelimit" default:"1" description:"Minimum gas price limit to enforce for acceptance into the pool"`
	PriceBump    int           `long:"txpool.pricebump" default:"10" description:"Price bump percentage to replace an already existing transaction"`
	AccountSlots int           `long:"txpool.accountslots" default:"16" description:"Minimum number of executable transaction slots guaranteed per account"`
	GlobalSlots  int           `long:"txpool.globalslots" default:"5120" description:"Maximum number of executable transaction slots for all accounts"`
	AccountQueue int           `long:"txpool.accountqueue" default:"64" description:"Maximum number of non-executable transaction slots permitted per account"`
	GlobalQueue  int           `long:"txpool.globalqueue" default:"1024" description:"Maximum number of non-executable transaction slots for all accounts"`
	Lifetime     time.Duration `long:"txpool.lifetime" default:"3h0m0s" description:"Maximum amount of time non-executable transaction are queued"`
}

type PerformanceTuningOptions struct {
	Cache         int            `long:"cache" default:"1024" description:"Megabytes of memory allocated to internal caching (default = 4096 mainnet full node, 128 light mode)"`
	Database      int            `long:"cache.database" default:"50" description:"Percentage of cache memory allowance to use for database io"`
	Trie          int            `long:"cache.trie" default:"15" description:"Percentage of cache memory allowance to use for trie caching (default = 15% full mode, 30% archive mode)"`
	TrieJournal   string         `long:"cache.trie.journal" default:"triecache" description:"Disk journal directory for trie cache to survive node restarts"`
	TrieRejournal *time.Duration `long:"cache.trie.rejournal" default:"1h0m0s" description:"Time interval to regenerate the trie cache journal"`
	GC            int            `long:"cache.gc" default:"25" description:"Percentage of cache memory allowance to use for trie pruning (default = 25% full mode, 0% archive mode)"`
	Snapshot      int            `long:"cache.snapshot" default:"10" description:"Percentage of cache memory allowance to use for snapshot caching (default = 10% full mode, 20% archive mode)"`
	NoPrefetch    bool           `long:"cache.noprefetch" default:"false" description:"Disable heuristic state prefetch during block import (less CPU and disk IO, more time waiting for data)"`
	Preimages     bool           `long:"cache.preimages" default:"false" description:"Enable recording the SHA3/keccak preimages of trie keys"`
	FDLimit       int            `long:"fdlimit" default:"0" description:"Raise the open file descriptor resource limit (default = system fd limit)"`
}

type NetworkingOptions struct {
	BootNodes    []string `long:"bootnodes" default:"" description:"Comma separated enode URLs for P2P discovery bootstrap"`
	DNS          []string `long:"discovery.dns" default:"" description:"Sets DNS discovery entry points (use \"\" to disable DNS)"`
	Port         int      `long:"port" default:"30303" description:"Network listening port"`
	MaxPeers     int      `long:"maxpeers" default:"50" description:"Maximum number of network peers (network disabled if set to 0)"`
	MaxPendPeers int      `long:"maxpendpeers" default:"0" description:"Maximum number of pending connection attempts (defaults used if set to 0)"`
	NAT          string   `long:"nat" default:"any" description:"NAT port mapping mechanism (any|none|upnp|pmp|extip:<IP>)"`
	NoDiscover   bool     `long:"nodiscover" default:"false" description:"Disables the peer discovery mechanism (manual peer addition)"`
	V5Disc       bool     `long:"v5disc" default:"false" description:"Enables the experimental RLPx V5 (Topic Discovery) mechanism"`
	NetRestrict  []string `long:"netrestrict" default:"" description:"Restricts network communication to the given IP networks (CIDR masks)"`
	NodeKey      string   `long:"nodekey" default:"" description:"P2P node key file"`
	NodeKeyHex   string   `long:"nodekeyhex" default:"" description:"P2P node key as hex (for testing)"`
}

type GasPriceOracleOptions struct {
	Blocks      int `long:"gpo.blocks" default:"20" description:"Number of recent blocks to check for gas prices"`
	Percentile  int `long:"gpo.percentile" default:"60" description:"Suggested gas price is the given percentile of a set of recent transaction gas prices"`
	MaxPrice    int `long:"gpo.maxprice" default:"500000000000" description:"Maximum transaction priority fee (or gasprice before London fork) to be recommended by gpo"`
	IgnorePrice int `long:"gpo.ignoreprice" default:"2" description:"Gas price below which gpo will ignore transactions"`
}

type LoggingOptions struct {
	FakePow          bool `long:"fakepow" default:"false" description:"Disables proof-of-work verification"`
	NoCompaction     bool `long:"nocompaction" default:"false" description:"Disables db compaction after import"`
	VerbosityEnabled bool
	Verbosity        int    `long:"verbosity" default:"3" description:"Logging verbosity: 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=detail"`
	VModule          string `long:"vmodule" default:"" description:"Per-module verbosity: comma-separated list of <pattern>=<level> (e.g. eth/*=5,p2p=4)"`
	JSON             bool   `long:"log.json" default:"false" description:"Format logs with JSON"`
	Backtrace        string `long:"log.backtrace" default:"" description:"Request a stack trace at a specific logging statement (e.g. \"block.go:271\")"`
	Debug            bool   `long:"log.debug" default:"false" description:"Prepends log messages with call-site location (file and line number)"`
	Pprof            bool   `long:"pprof" default:"false" description:"Enable the pprof HTTP server"`
	PprofAddr        string `long:"pprof.addr" default:"127.0.0.1" description:"pprof HTTP server listening interface"`
	PprofPort        int    `long:"pprof.port" default:"6060" description:"pprof HTTP server listening port"`
	MemProfileRate   int    `long:"pprof.memprofilerate" default:"524288" description:"Turn on memory profiling with the given rate"`
	BlockProfileRate int    `long:"pprof.blockprofilerate" default:"0" description:"Turn on block profiling with the given rate"`
	CPUProfile       string `long:"pprof.cpuprofile" default:"" description:"Write CPU profile to the given file"`
	Trace            string `long:"trace" default:"" description:"Write execution trace to the given file"`
}

type MetricsOptions struct {
	Enabled          bool   `long:"metrics" description:"Enable metrics collection and reporting"`
	Expensive        bool   `long:"metrics.expensive" description:"Enable expensive metrics collection and reporting"`
	Addr             string `long:"metrics.addr" description:"Enable stand-alone metrics HTTP server listening interface" default:"127.0.0.1"`
	Port             int    `long:"metrics.port" description:"Metrics HTTP server listening port" default:"6060"`
	InfluxDB         bool   `long:"metrics.influxdb" description:"Enable metrics export/push to an external InfluxDB database"`
	InfluxDBEndpoint string `long:"metrics.influxdb.endpoint" description:"InfluxDB API endpoint to report metrics to" default:"http://localhost:8086"`
	InfluxDBDatabase string `long:"metrics.influxdb.database" description:"InfluxDB database name to push reported metrics to" default:"geth"`
	InfluxDBUsername string `long:"metrics.influxdb.username" description:"Username to authorize access to the database" default:"test"`
	InfluxDBPassword string `long:"metrics.influxdb.password" description:"Password to authorize access to the database" default:"test"`
	InfluxDBTags     string `long:"metrics.influxdb.tags" description:"Comma-separated InfluxDB tags (key/values) attached to all measurements" default:"host=localhost"`
	InfluxDBv2       bool   `long:"metrics.influxdbv2" description:"Enable metrics export/push to an external InfluxDB v2 database"`
	InfluxDBv2Token  string `long:"metrics.influxdb.token" description:"Token to authorize access to the database (v2 only)" default:"test"`
	Token            string `long:"metrics.influxdb.token" description:"Token to authorize access to the database (v2 only) (default: 'test')"`
	Bucket           string `long:"metrics.influxdb.bucket" description:"InfluxDB bucket name to push reported metrics to (v2 only) (default: 'geth')"`
	Organization     string `long:"metrics.influxdb.organization" description:"InfluxDB organization name (v2 only) (default: 'geth')"`
}

type Options struct {
	Snapshot            bool `long:"snapshot"`
	BloomFilterSize     int  `long:"bloomfilter.size" default:"2048"`
	IgnoreLegacyReceipt bool `long:"ignore-legacy-receipts"`
}
