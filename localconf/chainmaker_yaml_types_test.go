package localconf

import(
	"testing"
	"fmt"
	"encoding/json"
	"bytes"
)

func formatStruct(t interface{}){
	bs,_ := json.Marshal(t)
	var out bytes.Buffer
	json.Indent(&out,bs,"","\t")
	fmt.Printf("%v\n",out.String())
}

func TestTypes(t *testing.T) {
	var setting = &CMConfig{}
	if err := setting.ReadFile("../config/chainmaker.yml"); err != nil {
		t.Errorf("read chainmaker.yml error: %v",err)
	}

	formatStruct(setting)

	if err := setting.ReadFile("../config/log.yml");err != nil{
		t.Errorf("read log.yml error: %v",err)
	}

	if err := setting.WriteFile("../config/test-1.yml",0664); err != nil{
		t.Errorf("write test-1.yml error: %v",err)
	}

}


