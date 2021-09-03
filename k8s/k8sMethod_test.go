package k8s

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	chain, _ := NewChainMakerType("/home/magatron/.kube/config", "wx-org1-chainmaker-org", "test", "10", "../output/chainmaker/ca", "../output/chainmaker/wx-org1-chainmaker-org/config", "../output/chainmaker/wx-org1-chainmaker-org/node", "../output/chainmaker/wx-org1-chainmaker-org/user")
	list, err := chain.serviceList()
	if err != nil {
		t.Error(t)
	}

	fmt.Println(list)

}

func TestServiceGet(t *testing.T) {
	chain, _ := NewChainMakerType("/home/magatron/.kube/config", "wx-org1-chainmaker-org", "test", "10", "../output/chainmaker/ca", "../output/chainmaker/wx-org1-chainmaker-org/config", "../output/chainmaker/wx-org1-chainmaker-org/node", "../output/chainmaker/wx-org1-chainmaker-org/user")
	service,err := chain.serviceGet()
	if err != nil{
		t.Error(t)
	}

	fmt.Println(service)
}
