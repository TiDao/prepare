package k8s

import(
	"testing"
	"fmt"
	"encoding/json"
	"bytes"
)

func showJson(v interface{}){
	data,_ := json.Marshal(v)
	var out bytes.Buffer
	json.Indent(&out,data,"","\t")
	fmt.Printf("%v\n",out.String())
}

func TestDeploymentInit(t *testing.T){
	chain := &ChainMakerType{
		NodeName: "test",
		NameSpace: "test",
	}
	err := chain.deploymentInit("../output/chainmaker/ca","../output/chainmaker/wx-org1.chainmaker.org/config","../output/chainmaker/wx-org1.chainmaker.org/node","../output.chainmaker/wx-org1.chainmaker.org/user")
	if err != nil{
		t.Error(err)
	}

	showJson(chain.Deployment)
}
