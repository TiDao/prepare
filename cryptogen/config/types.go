package config

type userConfig struct {
	Type       string         `yaml:"type"`
	Count      int32          `yaml:"count"`
	Location   locationConfig `yaml:"location"`
	ExpireYear int32          `yaml:"expire_year"`
}

type nodeConfig struct {
	Type     string         `yaml:"type"`
	Count    int32          `yaml:"count"`
	Location locationConfig `yaml:"location"`
	Specs    specsConfig    `yaml:"specs"`
}

type specsConfig struct {
	ExpireYear int32    `yaml:"expire_year"`
	SANS       []string `yaml:"sans"`
}

type caConfig struct {
	Location locationConfig `yaml:"location"`
	Specs    specsConfig    `yaml:"specs"`
}

type itemConfig struct {
	Domain   string         `yaml:"domain"`
	HostName string         `yaml:"host_name"`
	PKAlgo   string         `yaml:"pk_algo"`
	SKIHash  string         `yaml:"ski_hash"`
	Specs    specsConfig    `yaml:"specs"`
	Location locationConfig `yaml:"location"`
	Count    int32          `yaml:"count"`
	CA       caConfig       `yaml:"ca"`
	Node     []nodeConfig   `yaml:"node"`
	User     []userConfig   `yaml:"user"`
}

type locationConfig struct {
	Country  string `yaml:"country"`
	Locality string `yaml:"locality"`
	Province string `yaml:"province"`
}

type CryptoGenConfig struct {
	Item []itemConfig `yaml:"crypto_config"`
}
