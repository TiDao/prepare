package cryptogen

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	//"log"
)

var cryptoGenConfig *CryptoGenConfig

const (
	defaultCryptoConfigPath = "../config/crypto_config_template.yml"
)

func LoadCryptoGenConfig(path string) {
	cryptoGenConfig = &CryptoGenConfig{}

	if err := cryptoGenConfig.loadConfig(path); err != nil {
	//	log.Fatalf("load crypto config [%s] failed, %s",
	//		path, err)
		return err
	}

	//cryptoGenConfig.printLog()
}

func GetCryptoGenConfig() *CryptoGenConfig {
	return cryptoGenConfig
}

func (c *CryptoGenConfig) loadConfig(path string) error {
	if path == "" {
		path = defaultCryptoConfigPath
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *CryptoGenConfig) printLog() {
	// fmt.Printf("Load crypto config success!\n")
	// fmt.Printf("%+v\n", cryptoGenConfig)
}
