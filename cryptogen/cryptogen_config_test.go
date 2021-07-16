package cryptogen

import "testing"

func TestLoadCryptoGenConfig(t *testing.T) {
	LoadCryptoGenConfig("../config/crypto_config_template.yml")
}
