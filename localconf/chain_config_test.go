package localconf

import (
	"testing"
	"fmt"
)

func TestReadAndWrite(t *testing.T) {
	var config ChainConfig
	err := config.ReadFile("../config/yml/bc1.yml")
	if err != nil{
		t.Error(err)
	}
	fmt.Println(config)

	err = config.WriteFile("../config/yml/test-bc1.yaml",0664)
	if err != nil{
		t.Error(err)
	}
}
