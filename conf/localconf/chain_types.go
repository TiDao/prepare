package localconf

import(
	"io/ioutil"
	"io/fs"
	"gopkg.in/yaml.v2"
)
// ChainConfig
type ChainConfig struct {
	// blockchain identifier
	ChainId string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// blockchain version
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// authentication type
	AuthType string `protobuf:"bytes,3,opt,name=auth_type,json=authType,proto3" json:"auth_type,omitempty"`
	// config sequence
	Sequence uint64 `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// encryption algorithm related configuration
	Crypto *CryptoConfig `protobuf:"bytes,5,opt,name=crypto,proto3" json:"crypto,omitempty"`
	// block related configuration
	Block *BlockConfig `protobuf:"bytes,6,opt,name=block,proto3" json:"block,omitempty"`
	// core module related configuration
	Core *CoreConfig `protobuf:"bytes,7,opt,name=core,proto3" json:"core,omitempty"`
	// consensus related configuration
	Consensus *ConsensusConfig `protobuf:"bytes,8,opt,name=consensus,proto3" json:"consensus,omitempty"`
	// trusted root related configuration
	// for alliance members, the initial member's root info of the consortium; for public chain, there is no need to configure
	// Key: node_id; value: address, node public key / CA certificate
	TrustRoots []*TrustRootConfig `protobuf:"bytes,9,rep,name=trust_roots,json=trustRoots,proto3" json:"trust_roots,omitempty"`
	// permission related configuration
	ResourcePolicies []*ResourcePolicy `protobuf:"bytes,10,rep,name=resource_policies,json=resourcePolicies,proto3" json:"resource_policies,omitempty"`
	Contract         *ContractConfig   `protobuf:"bytes,11,opt,name=contract,proto3" json:"contract,omitempty"`
	// snapshot module related configuration
	Snapshot *SnapshotConfig `protobuf:"bytes,12,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	// scheduler module related configuration
	Scheduler *SchedulerConfig `protobuf:"bytes,13,opt,name=scheduler,proto3" json:"scheduler,omitempty"`
	// tx sim context module related configuration
	Context *ContextConfig `protobuf:"bytes,14,opt,name=context,proto3" json:"context,omitempty"`
}


type ResourcePolicy struct {
	// resource name
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// policy(permission)
	Policy *accesscontrol.Policy `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
}


type CryptoConfig struct {
	// enable Transaction timestamp verification or Not
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}


type BlockConfig struct {
	// enable transaction timestamp verification or Not
	TxTimestampVerify bool `protobuf:"varint,1,opt,name=tx_timestamp_verify,json=txTimestampVerify,proto3" json:"tx_timestamp_verify,omitempty"`
	// expiration time of transaction timestamp (seconds)
	TxTimeout uint32 `protobuf:"varint,2,opt,name=tx_timeout,json=txTimeout,proto3" json:"tx_timeout,omitempty"`
	// maximum number of transactions in a block
	BlockTxCapacity uint32 `protobuf:"varint,3,opt,name=block_tx_capacity,json=blockTxCapacity,proto3" json:"block_tx_capacity,omitempty"`
	// maximum block size, in MB
	BlockSize uint32 `protobuf:"varint,4,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	// block proposing interval, in ms
	BlockInterval uint32 `protobuf:"varint,5,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval,omitempty"`
}

type SchedulerConfig struct {
	// for evidence constract
	EnableEvidence bool `protobuf:"varint,1,opt,name=enable_evidence,json=enableEvidence,proto3" json:"enable_evidence,omitempty"`
}

type SnapshotConfig struct {
	// for the evidence contract
	EnableEvidence bool `protobuf:"varint,1,opt,name=enable_evidence,json=enableEvidence,proto3" json:"enable_evidence,omitempty"`
}

type ContextConfig struct {
	// for the evidence contract
	EnableEvidence bool `protobuf:"varint,1,opt,name=enable_evidence,json=enableEvidence,proto3" json:"enable_evidence,omitempty"`
}

type CoreConfig struct {
	// [0, 60], the time when the transaction scheduler gets the transaction from the transaction pool to schedule
	TxSchedulerTimeout uint64 `protobuf:"varint,1,opt,name=tx_scheduler_timeout,json=txSchedulerTimeout,proto3" json:"tx_scheduler_timeout,omitempty"`
	// [0, 60], the time-out for verification after the transaction scheduler obtains the transaction from the block
	TxSchedulerValidateTimeout uint64 `protobuf:"varint,2,opt,name=tx_scheduler_validate_timeout,json=txSchedulerValidateTimeout,proto3" json:"tx_scheduler_validate_timeout,omitempty"`
}

type ConsensusConfig struct {
	// consensus type
	Type consensus.ConsensusType `protobuf:"varint,1,opt,name=type,proto3,enum=consensus.ConsensusType" json:"type,omitempty"`
	// organization list of nodes
	Nodes []*OrgConfig `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// expand the field, record the difficulty, reward and other consensus algorithm configuration
	ExtConfig []*common.KeyValuePair `protobuf:"bytes,3,rep,name=ext_config,json=extConfig,proto3" json:"ext_config,omitempty"`
	// Initialize the configuration of DPOS
	DposConfig []*common.KeyValuePair `protobuf:"bytes,4,rep,name=dpos_config,json=dposConfig,proto3" json:"dpos_config,omitempty"`
}

// organization related configuration
type OrgConfig struct {
	// organization identifier
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// address list owned by the organization
	// Deprecated , replace by node_id
	Address []string `protobuf:"bytes,2,rep,name=address,proto3" json:"address,omitempty"`
	// node id list owned by the organization
	NodeId []string `protobuf:"bytes,3,rep,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

// trusted root related configuration
type TrustRootConfig struct {
	// oranization ideftifier
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// root certificate / public key
	Root string `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
}

type ContractConfig struct {
	EnableSqlSupport bool `protobuf:"varint,1,opt,name=enable_sql_support,json=enableSqlSupport,proto3" json:"enable_sql_support,omitempty"`
}

func (config *ChainConfig) ReadFile(fileName string) error {
	data,err := ioutil.ReadFile(fileName)
	if err != nil{
		return err
	}
}
