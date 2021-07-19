package localconf

import(
	"testing"
	"fmt"
	//"encoding/json"
	//"bytes"
)

//func formatStruct(t interface{}){
//	bs,_ := json.Marshal(t)
//	var out bytes.Buffer
//	json.Indent(&out,bs,"","\t")
//	fmt.Printf("%v\n",out.String())
//}

func TestTypes(t *testing.T) {
	var setting = &CMConfig{}
	if err := setting.ReadFile("../config/chainmaker.yml"); err != nil {
		fmt.Println("chainmaker.yml")
		t.Error("unmarsh chainamker.yml",err)
	}

	//formatStruct(setting)

	if err := setting.ReadFile("../config/log.yml");err != nil{
		fmt.Println("log.yml")
		t.Error("unmarshal log.yml: ",err)
	}

	if err := setting.WriteFile("../config/test-1.yml",0664); err != nil{
		t.Error("write", err)
	}

}


