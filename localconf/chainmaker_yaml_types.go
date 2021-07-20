/*
Copyright (C) BABEC. All rights reserved.
Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package localconf

import (
	"chainmaker.org/chainmaker-go/logger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"io/fs"
)

type nodeConfig struct {
	Type            string       `yaml:"type"`
	CertFile        string       `yaml:"cert_file"`
	PrivKeyFile     string       `yaml:"priv_key_file"`
	PrivKeyPassword string       `yaml:"priv_key_password"`
	AuthType        string       `yaml:"auth_type"`
	P11Config       pkcs11Config `yaml:"pkcs11"`
	NodeId          string       `yaml:"node_id"`
	OrgId           string       `yaml:"org_id"`
	SignerCacheSize int          `yaml:"signer_cache_size"`
	CertCacheSize   int          `yaml:"cert_cache_size"`
}

type netConfig struct {
	Provider                string            `yaml:"provider"`
	ListenAddr              string            `yaml:"listen_addr"`
	PeerStreamPoolSize      int               `yaml:"peer_stream_pool_size"`
	MaxPeerCountAllow       int               `yaml:"max_peer_count_allow"`
	PeerEliminationStrategy int               `yaml:"peer_elimination_strategy"`
	Seeds                   []string          `yaml:"seeds"`
	TLSConfig               netTlsConfig      `yaml:"tls"`
	BlackList               blackList         `yaml:"blacklist"`
	CustomChainTrustRoots   []chainTrustRoots `yaml:"custom_chain_trust_roots"`
}

type netTlsConfig struct {
	Enabled     bool   `yaml:"enabled"`
	PrivKeyFile string `yaml:"priv_key_file"`
	CertFile    string `yaml:"cert_file"`
}

type pkcs11Config struct {
	Enabled          bool   `yaml:"enabled"`
	Library          string `yaml:"library"`
	Label            string `yaml:"label"`
	Password         string `yaml:"password"`
	SessionCacheSize int    `yaml:"session_cache_size"`
	Hash             string `yaml:"hash"`
}

type blackList struct {
	Addresses []string `yaml:"addresses"`
	NodeIds   []string `yaml:"node_ids"`
}

type chainTrustRoots struct {
	ChainId    string       `yaml:"chain_id"`
	TrustRoots []trustRoots `yaml:"trust_roots"`
}

type trustRoots struct {
	OrgId string `yaml:"org_id"`
	Root  string `yaml:"root"`
}

type rpcConfig struct {
	Provider                               string           `yaml:"provider"`
	Port                                   int              `yaml:"port"`
	TLSConfig                              tlsConfig        `yaml:"tls"`
	RateLimitConfig                        rateLimitConfig  `yaml:"ratelimit"`
	SubscriberConfig                       subscriberConfig `yaml:"subscriber"`
	CheckChainConfTrustRootsChangeInterval int              `yaml:"check_chain_conf_trust_roots_change_interval"`
}

type tlsConfig struct {
	Mode                  string `yaml:"mode"`
	PrivKeyFile           string `yaml:"priv_key_file"`
	CertFile              string `yaml:"cert_file"`
	TestClientPrivKeyFile string `yaml:"test_client_priv_key_file"`
	TestClientCertFile    string `yaml:"test_client_cert_file"`
}

type rateLimitConfig struct {
	TokenPerSecond  int `yaml:"token_per_second"`
	TokenBucketSize int `yaml:"token_bucket_size"`
}

type subscriberConfig struct {
	RateLimitConfig rateLimitConfig `yaml:"ratelimit"`
}

type debugConfig struct {
	IsCliOpen           bool `yaml:"is_cli_open"`
	IsHttpOpen          bool `yaml:"is_http_open"`
	IsProposer          bool `yaml:"is_proposer"`
	IsNotRWSetCheck     bool `yaml:"is_not_rwset_check"`
	IsConcurPropose     bool `yaml:"is_concur_propose"`
	IsConcurVerify      bool `yaml:"is_concur_verify"`
	IsSolo              bool `yaml:"is_solo"`
	IsHaltPropose       bool `yaml:"is_halt_propose"`
	IsSkipAccessControl bool `yaml:"is_skip_access_control"` // true: minimize access control; false: use full access control
	IsTraceMemoryUsage  bool `yaml:"is_trace_memory_usage"`  // true for trace memory usage information periodically

	IsProposeDuplicately          bool `yaml:"is_propose_duplicately"`           // Simulate a node which would propose duplicate after it has proposed Proposal
	IsProposeMultiNodeDuplicately bool `yaml:"is_propose_multinode_duplicately"` // Simulate a malicious node which would propose duplicate proposals
	IsProposalOldHeight           bool `yaml:"is_proposal_old_height"`
	IsPrevoteDuplicately          bool `yaml:"is_prevote_duplicately"`   // Simulate a malicious node which would prevote duplicately
	IsPrevoteOldHeight            bool `yaml:"is_prevote_old_height"`    // Simulate a malicious node which would prevote for oldheight
	IsPrevoteLost                 bool `yaml:"is_prevote_lost"`          //prevote vote lost
	IsPrecommitDuplicately        bool `yaml:"is_precommit_duplicately"` //Simulate a malicious node which would propose duplicate precommits
	IsPrecommitOldHeight          bool `yaml:"is_precommit_old_height"`  // Simulate a malicious node which would Precommit a lower height than current height

	IsProposeLost    bool `yaml:"is_propose_lost"`     //proposal vote lost
	IsProposeDelay   bool `yaml:"is_propose_delay"`    //proposal lost
	IsPrevoteDelay   bool `yaml:"is_prevote_delay"`    //network problem resulting in preovote lost
	IsPrecommitLost  bool `yaml:"is_precommit_lost"`   //precommit vote lost
	IsPrecommitDelay bool `yaml:"is_prevcommit_delay"` //network problem resulting in precommit lost

	IsCommitWithoutPublish bool `yaml:"is_commit_without_publish"` //if the node committing block without publishing, TRUE；else, FALSE
	IsPrevoteInvalid       bool `yaml:"is_prevote_invalid"`        //simulate a node which sends an invalid prevote(hash=nil)
	IsPrecommitInvalid     bool `yaml:"is_precommit_invalid"`      //simulate a node which sends an invalid precommit(hash=nil)

	IsModifyTxPayload    bool `yaml:"is_modify_tx_payload"`
	IsExtreme            bool `yaml:"is_extreme"` //extreme fast mode
	UseNetMsgCompression bool `yaml:"use_net_msg_compression"`
	IsNetInsecurity      bool `yaml:"is_net_insecurity"`
}

type BlockchainConfig struct {
	ChainId string
	Genesis string
}

type StorageConfig struct {
	//默认的Leveldb配置，如果每个DB有不同的设置，可以在自己的DB中进行设置
	StorePath            string `yaml:"store_path"`
	DbPrefix             string `yaml:"db_prefix"`
	WriteBufferSize      int    `yaml:"write_buffer_size"`
	BloomFilterBits      int    `yaml:"bloom_filter_bits"`
	BlockWriteBufferSize int    `yaml:"block_write_buffer_size"`
	//数据库模式：light只存区块头,normal存储区块头和交易以及生成的State,full存储了区块头、交易、状态和交易收据（读写集、日志等）
	//Mode string `yaml:"mode"`
	DisableHistoryDB       bool      `yaml:"disable_historydb"`
	DisableResultDB        bool      `yaml:"disable_resultdb"`
	DisableContractEventDB bool      `yaml:"disable_contract_eventdb"`
	LogDBWriteAsync        bool      `yaml:"logdb_write_async"`
	BlockDbConfig          *DbConfig `yaml:"blockdb_config"`
	StateDbConfig          *DbConfig `yaml:"statedb_config"`
	HistoryDbConfig        *DbConfig `yaml:"historydb_config"`
	ResultDbConfig         *DbConfig `yaml:"resultdb_config"`
	ContractEventDbConfig  *DbConfig `yaml:"contract_eventdb_config"`
	UnArchiveBlockHeight   uint64    `yaml:"unarchive_block_height"`
}

func (config *StorageConfig) setDefault() {
	if config.DbPrefix != "" {
		if config.BlockDbConfig != nil && config.BlockDbConfig.SqlDbConfig != nil && config.BlockDbConfig.SqlDbConfig.DbPrefix == "" {
			config.BlockDbConfig.SqlDbConfig.DbPrefix = config.DbPrefix
		}
		if config.StateDbConfig != nil && config.StateDbConfig.SqlDbConfig != nil && config.StateDbConfig.SqlDbConfig.DbPrefix == "" {
			config.StateDbConfig.SqlDbConfig.DbPrefix = config.DbPrefix
		}
		if config.HistoryDbConfig != nil && config.HistoryDbConfig.SqlDbConfig != nil && config.HistoryDbConfig.SqlDbConfig.DbPrefix == "" {
			config.HistoryDbConfig.SqlDbConfig.DbPrefix = config.DbPrefix
		}
		if config.ResultDbConfig != nil && config.ResultDbConfig.SqlDbConfig != nil && config.ResultDbConfig.SqlDbConfig.DbPrefix == "" {
			config.ResultDbConfig.SqlDbConfig.DbPrefix = config.DbPrefix
		}
		if config.ContractEventDbConfig != nil && config.ContractEventDbConfig.SqlDbConfig != nil && config.ContractEventDbConfig.SqlDbConfig.DbPrefix == "" {
			config.ContractEventDbConfig.SqlDbConfig.DbPrefix = config.DbPrefix
		}
	}
}
func (config *StorageConfig) GetBlockDbConfig() *DbConfig {
	if config.BlockDbConfig == nil {
		return config.GetDefaultDBConfig()
	}
	config.setDefault()
	return config.BlockDbConfig
}
func (config *StorageConfig) GetStateDbConfig() *DbConfig {
	if config.StateDbConfig == nil {
		return config.GetDefaultDBConfig()
	}
	config.setDefault()
	return config.StateDbConfig
}
func (config *StorageConfig) GetHistoryDbConfig() *DbConfig {
	if config.HistoryDbConfig == nil {
		return config.GetDefaultDBConfig()
	}
	config.setDefault()
	return config.HistoryDbConfig
}
func (config *StorageConfig) GetResultDbConfig() *DbConfig {
	if config.ResultDbConfig == nil {
		return config.GetDefaultDBConfig()
	}
	config.setDefault()
	return config.ResultDbConfig
}
func (config *StorageConfig) GetContractEventDbConfig() *DbConfig {
	if config.ContractEventDbConfig == nil {
		return config.GetDefaultDBConfig()
	}
	config.setDefault()
	return config.ContractEventDbConfig
}
func (config *StorageConfig) GetDefaultDBConfig() *DbConfig {
	lconfig := &LevelDbConfig{
		StorePath:            config.StorePath,
		WriteBufferSize:      config.WriteBufferSize,
		BloomFilterBits:      config.BloomFilterBits,
		BlockWriteBufferSize: config.WriteBufferSize,
	}
	return &DbConfig{
		Provider:      "leveldb",
		LevelDbConfig: lconfig,
	}
}

//根据配置的DisableDB的情况，确定当前配置活跃的数据库数量
func (config *StorageConfig) GetActiveDBCount() int {
	count := 5
	if config.DisableContractEventDB {
		count--
	}
	if config.DisableHistoryDB {
		count--
	}
	if config.DisableResultDB {
		count--
	}
	return count
}

type DbConfig struct {
	//leveldb,rocksdb,sql
	Provider      string         `yaml:"provider"`
	LevelDbConfig *LevelDbConfig `yaml:"leveldb_config"`
	SqlDbConfig   *SqlDbConfig   `yaml:"sqldb_config"`
}

const DbConfig_Provider_Sql = "sql"
const DbConfig_Provider_LevelDb = "leveldb"
const DbConfig_Provider_RocksDb = "rocksdb"

func (dbc *DbConfig) IsKVDB() bool {
	return dbc.Provider == DbConfig_Provider_LevelDb || dbc.Provider == DbConfig_Provider_RocksDb
}
func (dbc *DbConfig) IsSqlDB() bool {
	return dbc.Provider == DbConfig_Provider_Sql || dbc.Provider == "mysql" || dbc.Provider == "rdbms" //兼容其他配置情况
}

type LevelDbConfig struct {
	StorePath            string `yaml:"store_path"`
	WriteBufferSize      int    `yaml:"write_buffer_size"`
	BloomFilterBits      int    `yaml:"bloom_filter_bits"`
	BlockWriteBufferSize int    `yaml:"block_write_buffer_size"`
}
type SqlDbConfig struct {
	//mysql, sqlite, postgres, sqlserver
	SqlDbType       string `yaml:"sqldb_type"`
	Dsn             string `yaml:"dsn"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifeTime int    `yaml:"conn_max_lifetime"` //second
	SqlLogMode      string `yaml:"sqllog_mode"`       //Silent,Error,Warn,Info
	SqlVerifier     string `yaml:"sql_verifier"`      //simple,safe
	DbPrefix        string `yaml:"db_prefix"`
}

const SqlDbConfig_SqlDbType_MySQL = "mysql"
const SqlDbConfig_SqlDbType_Sqlite = "sqlite"

type txPoolConfig struct {
	PoolType            string `yaml:"pool_type"`
	MaxTxPoolSize       uint32 `yaml:"max_txpool_size"`
	MaxConfigTxPoolSize uint32 `yaml:"max_config_txpool_size"`
	IsMetrics           bool   `yaml:"is_metrics"`
	Performance         bool   `yaml:"performance"`
	BatchMaxSize        int    `yaml:"batch_max_size"`
	BatchCreateTimeout  int64  `yaml:"batch_create_timeout"`
	CacheFlushTicker    int64  `yaml:"cache_flush_ticker"`
	CacheThresholdCount int64  `yaml:"cache_threshold_count"`
	CacheFlushTimeOut   int64  `yaml:"cache_flush_timeout"`
	AddTxChannelSize    int64  `yaml:"add_tx_channel_size"`
}

type syncConfig struct {
	BroadcastTime             uint32  `yaml:"broadcast_time"`
	BlockPoolSize             uint32  `yaml:"block_pool_size"`
	WaitTimeOfBlockRequestMsg uint32  `yaml:"wait_time_requested"`
	BatchSizeFromOneNode      uint32  `yaml:"batch_Size_from_one_node"`
	ProcessBlockTick          float64 `yaml:"process_block_tick"`
	NodeStatusTick            float64 `yaml:"node_status_tick"`
	LivenessTick              float64 `yaml:"liveness_tick"`
	SchedulerTick             float64 `yaml:"scheduler_tick"`
	ReqTimeThreshold          float64 `yaml:"req_time_threshold"`
	DataDetectionTick         float64 `yaml:"data_detection_tick"`
}

type spvConfig struct {
	RefreshReqCacheMills     int64 `yaml:"refresh_reqcache_mils"`
	MessageCacheSize         int64 `yaml:"message_cahche_size"`
	ReSyncCheckIntervalMills int64 `yaml:"resync_check_interval_mils"`
	SyncTimeoutMills         int64 `yaml:"sync_timeout_mils"`
	ReqSyncBlockNum          int64 `yaml:"reqsync_blocknum"`
	MaxReqSyncBlockNum       int64 `yaml:"max_reqsync_blocknum"`
	PeerActiveTime           int64 `yaml:"peer_active_time"`
}

type monitorConfig struct {
	Enabled bool `yaml:"enabled"`
	Port    int  `yaml:"port"`
}

type pprofConfig struct {
	Enabled bool `yaml:"enabled"`
	Port    int  `yaml:"port"`
}

type redisConfig struct {
	Url          string `yaml:"url"`
	Auth         string `yaml:"auth"`
	DB           int    `yaml:"db"`
	MaxIdle      int    `yaml:"max_idle"`
	MaxActive    int    `yaml:"max_active"`
	IdleTimeout  int    `yaml:"idle_timeout"`
	CacheTimeout int    `yaml:"cache_timeout"`
}

type clientConfig struct {
	OrgId           string `yaml:"org_id"`
	UserKeyFilePath string `yaml:"user_key_file_path"`
	UserCrtFilePath string `yaml:"user_crt_file_path"`
	HashType        string `yaml:"hash_type"`
}

type schedulerConfig struct {
	RWSetLog bool `yaml:"rwset_log"`
}

type coreConfig struct {
	Evidence bool `yaml:"evidence"`
}

// CMConfig - Local config struct
type CMConfig struct {
	LogConfig        logger.LogConfig   `yaml:"log"`
	NetConfig        netConfig          `yaml:"net"`
	NodeConfig       nodeConfig         `yaml:"node"`
	RpcConfig        rpcConfig          `yaml:"rpc"`
	BlockChainConfig []BlockchainConfig `yaml:"blockchain"`
	StorageConfig    StorageConfig      `yaml:"storage"`
	TxPoolConfig     txPoolConfig       `yaml:"txpool"`
	SyncConfig       syncConfig         `yaml:"sync"`
	SpvConfig        spvConfig          `yaml:"spv"`

	// 开发调试使用
	DebugConfig     debugConfig     `yaml:"debug"`
	PProfConfig     pprofConfig     `yaml:"pprof"`
	MonitorConfig   monitorConfig   `yaml:"monitor"`
	CoreConfig      coreConfig      `yaml:"core"`
	SchedulerConfig schedulerConfig `yaml:"scheduler"`
}

// write config into file
func (config  *CMConfig) WriteFile(fileName string,fileMode fs.FileMode) error {

	data,err := yaml.Marshal(config)
	if err != nil{
		return err
	}

	err = ioutil.WriteFile(fileName,data,fileMode)
	if err != nil{
		return err
	}

	return nil
}

//read config from configfile
func (config *CMConfig) ReadFile(fileName string) error {

	data,err := ioutil.ReadFile(fileName)
	if err != nil{
		return err
	}

	err = yaml.Unmarshal(data,config)
	if err != nil{
		return err
	}

	return nil
}

// GetBlockChains - get blockchain config list
func (c *CMConfig) GetBlockChains() []BlockchainConfig {
	return c.BlockChainConfig
}
