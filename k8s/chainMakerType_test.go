package k8s

import(
	"testing"
	//"log"
)

func TestNodeCreate(t *testing.T){
	chain1,_ := NewChainMakerType("/home/magatron/.kube/config","chainmaker-1","test","10","../output/chainmaker/ca","../output/chainmaker/wx-org1.chainmaker.org/config","../output/chainmaker/wx-org1.chainmaker.org/node","../output/chainmaker/wx-org1.chainmaker.org/user")
	chain2,_ := NewChainMakerType("/home/magatron/.kube/config","chainmaker-2","test","10","../output/chainmaker/ca","../output/chainmaker/wx-org2.chainmaker.org/config","../output/chainmaker/wx-org2.chainmaker.org/node","../output/chainmaker/wx-org2.chainmaker.org/user")
	chain3,_ := NewChainMakerType("/home/magatron/.kube/config","chainmaker-3","test","10","../output/chainmaker/ca","../output/chainmaker/wx-org3.chainmaker.org/config","../output/chainmaker/wx-org3.chainmaker.org/node","../output/chainmaker/wx-org3.chainmaker.org/user")
	chain4,_ := NewChainMakerType("/home/magatron/.kube/config","chainmaker-4","test","10","../output/chainmaker/ca","../output/chainmaker/wx-org4.chainmaker.org/config","../output/chainmaker/wx-org4.chainmaker.org/node","../output/chainmaker/wx-org4.chainmaker.org/user")
	err := chain1.NodeCreate()
	if err != nil{
		t.Error(err)
	}

	err = chain2.NodeCreate()
	if err != nil{
		t.Error(err)
	}

	err = chain3.NodeCreate()
	if err != nil{
		t.Error(err)
	}

	err = chain4.NodeCreate()
	if err != nil{
		t.Error(err)
	}
}

func TestNodeDelete(t *testing.T){
	chain1,_ := NewChainMakerType("/home/magatron/.kube/config","chainmaker-1","test","10","../output/chainmaker/ca","../output/chainmaker/wx-org1.chainmaker.org/config","../output/chainmaker/wx-org1.chainmaker.org/node","../output/chainmaker/wx-org1.chainmaker.org/user")
	chain2,_ := NewChainMakerType("/home/magatron/.kube/config","chainmaker-2","test","10","../output/chainmaker/ca","../output/chainmaker/wx-org2.chainmaker.org/config","../output/chainmaker/wx-org2.chainmaker.org/node","../output/chainmaker/wx-org2.chainmaker.org/user")
	chain3,_ := NewChainMakerType("/home/magatron/.kube/config","chainmaker-3","test","10","../output/chainmaker/ca","../output/chainmaker/wx-org3.chainmaker.org/config","../output/chainmaker/wx-org3.chainmaker.org/node","../output/chainmaker/wx-org3.chainmaker.org/user")
	chain4,_ := NewChainMakerType("/home/magatron/.kube/config","chainmaker-4","test","10","../output/chainmaker/ca","../output/chainmaker/wx-org4.chainmaker.org/config","../output/chainmaker/wx-org4.chainmaker.org/node","../output/chainmaker/wx-org4.chainmaker.org/user")
	err := chain1.NodeDelete()
	if err != nil{
		t.Error(err)
	}

	err = chain2.NodeDelete()
	if err != nil{
		t.Error(err) }

	err = chain3.NodeDelete()
	if err != nil{
		t.Error(err)
	}

	err = chain4.NodeDelete()
	if err != nil{
		t.Error(err)
	}
}
