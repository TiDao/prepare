package k8s
import(
	"testing"
	"fmt"
)

func TestList(t *testing.T) {
		chain,_ := NewChainMakerType("/home/magatron/.kube/config","wx-org1-chainmaker-org","test","10","../output/chainmaker/ca","../output/chainmaker/wx-org1-chainmaker-org/config","../output/chainmaker/wx-org1-chainmaker-org/node","../output/chainmaker/wx-org1-chainmaker-org/user")
		list,err := chain.List()
		if err != nil{
			t.Error(t)
		}

		fmt.Println(list)

}
