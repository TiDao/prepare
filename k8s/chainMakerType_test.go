package k8s

import(
	"testing"
	//"log"
)

func TestNewChainMakerType(t *testing.T){
	chain,err := NewChainMakerType("/home/magatron/.kube/config","chainnode-1","test","10","../output/chainmaker/ca","../output/chainmaker/wx-org1.chainmaker.org/config","../output/chainmaker/wx-org1.chainmaker.org/node","../output/chainmaker/wx-org1.chainmaker.org/user")
	if err != nil{
		t.Error(err)
	}
	//for _,data := range chain.Secrets{
	//	log.Println(data.Data)
	//}

	err = chain.NodeCreate()
	if err != nil{
		t.Error(err)
	}
}
