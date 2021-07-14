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
	if err := setting.ReadFile("./chainmaker.yml"); err != nil {
		t.Error(err)
	}

	if err := setting.WriteFile("./test-1.yml",0664); err != nil{
		t.Error(err)
	}

}


