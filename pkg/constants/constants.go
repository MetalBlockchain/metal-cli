// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package constants

import (
	"time"
)

const (
	DefaultPerms755 = 0o755

	BaseDirName = ".metal-cli"
	LogDir      = "logs"

	ServerRunFile      = "gRPCserver.run"
	AvalancheCliBinDir = "bin"
	RunDir             = "runs"

	SuffixSeparator = "_"
	SidecarFileName = "sidecar.json"
	GenesisFileName = "genesis.json"
	SidecarSuffix   = SuffixSeparator + SidecarFileName
	GenesisSuffix   = SuffixSeparator + GenesisFileName

	SidecarVersion = "1.2.1"

	MaxLogFileSize   = 4
	MaxNumOfLogFiles = 5
	RetainOldFiles   = 0 // retain all old log files

	RequestTimeout = 3 * time.Minute

	SimulatePublicNetwork = "SIMULATE_PUBLIC_NETWORK"
	FujiAPIEndpoint       = "https://tahoe.metalblockchain.org"
	MainnetAPIEndpoint    = "https://api.metalblockchain.org"

	// this depends on bootstrap snapshot
	LocalAPIEndpoint = "http://127.0.0.1:9650"
	LocalNetworkID   = 1337

	DefaultTokenName = "TEST"

	HealthCheckInterval = 100 * time.Millisecond

	// it's unlikely anyone would want to name a snapshot `default`
	// but let's add some more entropy
	SnapshotsDirName             = "snapshots"
	DefaultSnapshotName          = "default-1654102509"
	BootstrapSnapshotArchiveName = "bootstrapSnapshot.tar.gz"
	BootstrapSnapshotLocalPath   = "assets/" + BootstrapSnapshotArchiveName
	BootstrapSnapshotURL         = "https://github.com/MetalBlockchain/metal-cli/raw/main/" + BootstrapSnapshotLocalPath
	BootstrapSnapshotSHA256URL   = "https://github.com/MetalBlockchain/metal-cli/raw/main/assets/sha256sum.txt"

	KeyDir     = "key"
	KeySuffix  = ".pk"
	YAMLSuffix = ".yml"

	TimeParseLayout    = "2006-01-02 15:04:05"
	MinStakeDuration   = 24 * 14 * time.Hour
	MaxStakeDuration   = 24 * 365 * time.Hour
	MaxStakeWeight     = 100
	MinStakeWeight     = 1
	DefaultStakeWeight = 20

	// The absolute minimum is 25 seconds, but set to 1 minute to allow for
	// time to go through the command
	StakingStartLeadTime   = 1 * time.Minute
	StakingMinimumLeadTime = 25 * time.Second

	DefaultConfigFileName = ".metal-cli"
	DefaultConfigFileType = "json"

	CustomVMDir = "vms"

	AvaLabsOrg          = "MetalBlockchain"
	AvalancheGoRepoName = "metalgo"
	SubnetEVMRepoName   = "subnet-evm"
	SpacesVMRepoName    = "spacesvm"

	AvalancheGoInstallDir = "metalgo"
	SubnetEVMInstallDir   = "subnet-evm"
	SpacesVMInstallDir    = "spacesvm"

	SubnetEVMBin = "subnet-evm"
	SpacesVMBin  = "spacesvm"

	DefaultNodeRunURL = "http://127.0.0.1:9650"

	APMDir                = ".apm"
	APMLogName            = "apm.log"
	DefaultAvaLabsPackage = "MetalBlockchain/metal-plugins-core"
	APMPluginDir          = "apm_plugins"

	// #nosec G101
	GithubAPITokenEnvVarName = "METAL_CLI_GITHUB_TOKEN"

	ReposDir       = "repos"
	SubnetDir      = "subnets"
	VMDir          = "vms"
	ChainConfigDir = "chains"

	SubnetConfigFileName       = "subnet.json"
	ChainConfigFileName        = "chain.json"
	PerNodeChainConfigFileName = "per-node-chain.json"

	GitRepoCommitName  = "Metal-CLI"
	GitRepoCommitEmail = "info@avax.network"

	AvaLabsMaintainers = "MetalBlockchain"

	UpgradeBytesFileName      = "upgrade.json"
	UpgradeBytesLockExtension = ".lock"
	NotAvailableLabel         = "Not available"
	BackendCmd                = "metal-cli-backend"

	AvalancheGoCompatibilityVersionAdded = "v1.9.2"
	AvalancheGoCompatibilityURL          = "https://raw.githubusercontent.com/MetalBlockchain/metalgo/master/version/compatibility.json"
	SubnetEVMRPCCompatibilityURL         = "https://raw.githubusercontent.com/MetalBlockchain/subnet-evm/master/compatibility.json"
	SpacesVMRPCCompatibilityURL          = "https://raw.githubusercontent.com/MetalBlockchain/spacesvm/master/compatibility.json"

	YesLabel = "Yes"
	NoLabel  = "No"

	SubnetIDLabel     = "SubnetID: "
	BlockchainIDLabel = "BlockchainID: "

	PluginDir = "plugins"
)
