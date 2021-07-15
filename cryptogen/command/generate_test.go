package command 

import(
	"testing"
	"chainmaker.org/chainmaker-cryptogen/config"
	
)

func TestGenerate(t *testing.T){
	config.LoadCryptoGenConfig("../../config/crypto_config_template.yml")
	err := Generate()
	if err != nil{
		t.Error(err)
	}
}
