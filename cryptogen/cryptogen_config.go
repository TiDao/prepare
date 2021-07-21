package cryptogen

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	//"log"
)

var CryptoConfig *CryptoGenConfig

const (
	defaultCryptoConfigPath = "../config/crypto_config_template.yml"
)

func LoadCryptoGenConfig(path string) error {
	CryptoConfig = &CryptoGenConfig{}

	if err := CryptoConfig.loadConfig(path); err != nil {
	//	log.Fatalf("load crypto config [%s] failed, %s",
	//		path, err)
		return err
	}

	return nil
	//cryptoGenConfig.printLog()
}

func GetCryptoGenConfig() *CryptoGenConfig {
	return CryptoConfig
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
