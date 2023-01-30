package web3signer_cmds_ai_generated

import (
	"fmt"
	"strconv"
	"strings"
)

type Web3SignerCmdArgs struct {
	Config
	TLSOpts
	AWSCfg
	AzureConfig

	// subcommand configs
	SlashingConfig
	KeyManagerConfig
}

const (
	startDashAC            = "-ac"
	web3signerStartCommand = "/opt/web3signer/bin/web3signer"
)

// CreateFieldsForCLI platformSubcommand should be eth1, eth2, or filecoin
func (w *Web3SignerCmdArgs) CreateFieldsForCLI(platformSubcommand string) ([]string, error) {
	args := []string{web3signerStartCommand}
	cargs, err := w.Config.SetArgs()
	if err != nil {
		return nil, err
	}
	args = append(args, cargs...)
	args = append(args, w.TLSOpts.SetArgs()...)
	args = append(args, w.AWSCfg.SetArgs()...)
	args = append(args, w.AzureConfig.SetArgs()...)
	args = append(args, platformSubcommand)

	args = append(args, w.SlashingConfig.SetArgs()...)
	args = append(args, w.KeyManagerConfig.SetArgs()...)

	cliFormatted := []string{startDashAC, strings.Join(args, " ")}
	return cliFormatted, nil
}

type TLSOpts struct {
	TLSKeyStoreFile     string
	TLSPasswordFile     string
	TLSAllowCAClients   bool
	TLSAllowAnyClient   bool
	TLSKnownClientsFile string
}

type Config struct {
	ConfigFile     string
	DataPath       string
	KeyStorePath   string
	LoggingLevel   string
	HTTPCors       string
	HTTPListen     string
	HTTPListenPort int
	HTTPAllowlist  string
	MetricsEnabled bool
	MetricsPort    int
	MetricsPrefix  string
}

type AWSCfg struct {
	SecretsEnabled      bool
	ConnectionCacheSize int64
	AuthMode            string
	AccessKeyID         string
	SecretAccessKey     string
	Region              string
	PrefixesFilter      []string
	TagNamesFilter      []string
	TagValuesFilter     []string
}

func (c *AWSCfg) SetArgs() []string {
	var args []string

	if c.SecretsEnabled {
		args = append(args, "--aws-secrets-enabled")
	}
	if c.ConnectionCacheSize > 0 {
		args = append(args, fmt.Sprintf("--aws-connection-cache-size=%d", c.ConnectionCacheSize))
	}
	if c.AuthMode != "" {
		args = append(args, fmt.Sprintf("--aws-secrets-auth-mode=%s", c.AuthMode))
	}
	if c.AccessKeyID != "" {
		args = append(args, fmt.Sprintf("--aws-secrets-access-key-id=%s", c.AccessKeyID))
	}
	if c.SecretAccessKey != "" {
		args = append(args, fmt.Sprintf("--aws-secrets-secret-access-key=%s", c.SecretAccessKey))
	}
	if c.Region != "" {
		args = append(args, fmt.Sprintf("--aws-secrets-region=%s", c.Region))
	}
	if len(c.PrefixesFilter) > 0 {
		var s string
		for _, prefix := range c.PrefixesFilter {
			s += prefix + ", "
		}
		s = strings.TrimSuffix(s, ", ")
		args = append(args, fmt.Sprintf("--aws-secrets-prefixes-filter=%s", s))
	}
	if len(c.TagNamesFilter) > 0 {
		var s string
		for _, prefix := range c.TagNamesFilter {
			s += prefix + ", "
		}
		s = strings.TrimSuffix(s, ", ")
		args = append(args, fmt.Sprintf("--aws-secrets-tag-names-filter=%s", s))
	}
	if len(c.TagValuesFilter) > 0 {
		var s string
		for _, prefix := range c.TagValuesFilter {
			s += prefix + ", "
		}
		s = strings.TrimSuffix(s, ", ")
		args = append(args, fmt.Sprintf("--aws-secrets-tag-values-filter=%s", s))
	}

	return args
}
func (c *Config) SetArgs() ([]string, error) {
	args := []string{}

	if c.ConfigFile != "" {
		args = append(args, "--config-file="+c.ConfigFile)
	}
	if c.DataPath != "" {
		args = append(args, "--data-path="+c.DataPath)
	}
	if c.KeyStorePath != "" {
		args = append(args, "--key-store-path="+c.KeyStorePath)
	}
	if c.LoggingLevel != "" {
		err := c.VerifyLogLevel()
		if err != nil {
			return nil, err
		}
		args = append(args, "--logging="+c.LoggingLevel)
	}
	if c.HTTPCors != "" {
		args = append(args, "--http-cors-origins="+c.HTTPCors)
	}
	if c.HTTPListen != "" {
		args = append(args, "--http-listen-host="+c.HTTPListen)
	}
	if c.HTTPListenPort != 0 {
		args = append(args, fmt.Sprintf("--http-listen-port=%d", c.HTTPListenPort))
	}
	if c.HTTPAllowlist != "" {
		args = append(args, "--http-host-allowlist="+c.HTTPAllowlist)
	}
	if c.MetricsEnabled {
		args = append(args, "--metrics-enabled=true")
	}
	if c.MetricsPort != 0 {
		args = append(args, fmt.Sprintf("--metrics-port=%d", c.MetricsPort))
	}
	if c.MetricsPrefix != "" {
		args = append(args, "--metrics-prefix="+c.MetricsPrefix)
	}
	return args, nil
}

// VerifyLogLevel Log Level Verification Function
func (c *Config) VerifyLogLevel() error {
	validLevels := []string{"OFF", "FATAL", "WARN", "INFO", "DEBUG", "TRACE", "ALL"}
	for _, level := range validLevels {
		if c.LoggingLevel == level {
			return nil
		}
	}
	return fmt.Errorf("invalid log level: %s", c.LoggingLevel)
}

func (opts *TLSOpts) SetArgs() []string {
	args := []string{}

	if opts.TLSKeyStoreFile != "" {
		args = append(args, "--tls-keystore-file="+opts.TLSKeyStoreFile)
	}

	if opts.TLSPasswordFile != "" {
		args = append(args, "--tls-keystore-password-file="+opts.TLSPasswordFile)
	}

	if opts.TLSAllowCAClients {
		args = append(args, "--tls-allow-ca-clients")
	}

	if opts.TLSAllowAnyClient {
		args = append(args, "--tls-allow-any-client=true")
	}

	if opts.TLSKnownClientsFile != "" {
		args = append(args, "--tls-known-clients-file="+opts.TLSKnownClientsFile)
	}

	return args
}

type AzureConfig struct {
	VaultEnabled  bool
	ClientId      string
	ClientSecret  string
	TenantId      string
	VaultAuthMode string
	VaultName     string
}

func (config *AzureConfig) SetArgs() []string {
	var args []string

	if config.VaultEnabled {
		args = append(args, fmt.Sprintf("--azure-vault-enabled=%v", config.VaultEnabled))
	}

	if config.ClientId != "" {
		args = append(args, fmt.Sprintf("--azure-client-id=%s", config.ClientId))
	}

	if config.ClientSecret != "" {
		args = append(args, fmt.Sprintf("--azure-client-secret=%s", config.ClientSecret))
	}

	if config.TenantId != "" {
		args = append(args, fmt.Sprintf("--azure-tenant-id=%s", config.TenantId))
	}

	if config.VaultAuthMode != "" {
		args = append(args, fmt.Sprintf("--azure-vault-auth-mode=%s", config.VaultAuthMode))
	}

	if config.VaultName != "" {
		args = append(args, fmt.Sprintf("--azure-vault-name=%s", config.VaultName))
	}

	return args
}

type KeyManagerConfig struct {
	KeyManagerAPIEenabled  bool
	KeystoresPasswordFile  string
	KeystoresPasswordsPath string
	KeystoresPath          string
	Network                string
}

func (kmc *KeyManagerConfig) SetArgs() []string {
	var retArgs []string

	if kmc.KeyManagerAPIEenabled {
		retArgs = append(retArgs, "--key-manager-api-enabled="+strconv.FormatBool(kmc.KeyManagerAPIEenabled))
	}

	if kmc.KeystoresPasswordFile != "" {
		retArgs = append(retArgs, "--keystores-password-file="+kmc.KeystoresPasswordFile)
	}

	if kmc.KeystoresPath != "" {
		retArgs = append(retArgs, "--keystores-passwords-path="+kmc.KeystoresPasswordsPath)
	}

	if kmc.KeystoresPath != "" {
		retArgs = append(retArgs, "--keystores-path="+kmc.KeystoresPath)
	}

	if kmc.Network != "" {
		retArgs = append(retArgs, "--network="+kmc.Network)
	}

	return retArgs
}

type SlashingConfig struct {
	ProtectionDbPassword           string
	ProtectionDbUrl                string
	ProtectionDbUsername           string
	ProtectionDisabled             bool
	ProtectionEnabled              bool
	ProtectionPruningEnabled       bool
	ProtectionPruningEpochsToKeep  int
	ProtectionPruningInterval      int
	ProtectionPruningSlotsPerEpoch int
}

func (sc *SlashingConfig) SetArgs() []string {
	if sc.ProtectionDisabled {
		sc.ProtectionEnabled = false
	} else {
		sc.ProtectionEnabled = true
	}
	args := []string{
		"--slashing-protection-enabled=" + strconv.FormatBool(sc.ProtectionEnabled),
	}
	if len(sc.ProtectionDbPassword) > 0 {
		args = append(args, "--slashing-protection-db-password="+sc.ProtectionDbPassword)
	}
	if len(sc.ProtectionDbUrl) > 0 {
		args = append(args, "--slashing-protection-db-url="+sc.ProtectionDbUrl)
	}
	if len(sc.ProtectionDbUsername) > 0 {
		args = append(args, "--slashing-protection-db-username="+sc.ProtectionDbUsername)
	}
	if sc.ProtectionPruningEnabled {
		args = append(args, "--slashing-protection-pruning-enabled="+strconv.FormatBool(sc.ProtectionPruningEnabled))
	}
	if sc.ProtectionPruningEpochsToKeep > 0 {
		args = append(args, fmt.Sprintf("--slashing-protection-pruning-epochs-to-keep=%d", sc.ProtectionPruningEpochsToKeep))
	}
	if sc.ProtectionPruningInterval > 0 {
		args = append(args, fmt.Sprintf("--slashing-protection-pruning-interval=%d", sc.ProtectionPruningInterval))
	}
	if sc.ProtectionPruningSlotsPerEpoch > 0 {
		args = append(args, fmt.Sprintf("--slashing-protection-pruning-slots-per-epoch=%d", sc.ProtectionPruningSlotsPerEpoch))
	}
	return args
}
