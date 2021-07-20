package cryptogen

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	LoadCryptoGenConfig("../config/crypto_config_template.yml")
	OutputDir = "../test_output"
	err := Generate()
	if err != nil {
		t.Error(err)
	}
}
