package localconf

import(
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

	var setting CMConfig
	config,err := ioutil.ReadFile("./chainmaker.yml")
	if err != nil{
		fmt.Print(err)
	}
	yaml.Unmarshal(config,&setting)
	formatStruct(setting)
	if err != nil{
		t.Fatalf("%v",err)
	}

	data2,err := yaml.Marshal(setting)
	if err != nil{
		t.Fatalf("%v",err)
	}

	ioutil.WriteFile("./test.yaml",data2,0664)
}


