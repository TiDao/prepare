package main

import(
	"flag"
	"fmt"
	"os"
)

var (
	nodeCNT uint
	chainCNT uint
	p2pPort uint
	rpcPort uint
)

func init(){
	nodeCNT = *flag.Uint("nodeCnt",4,"the number of nodes")
	chainCNT = *flag.Uint("chainCnt",1,"the number of chains")
	p2pPort = *flag.Uint("p2pPort",11300,"the port of p2p")
	rpcPort = *flag.Uint("rpcPort",12300,"the port of rpc")
}

func check_nodeNumber(nodeCNT uint) bool{
	var nodeNumber = []uint{1,4,5,10,13}
	for _,b := range nodeNumber{
		if b == nodeCNT{
			return true
		}
	}
	return false
}

func checkParams(nodeCNT,chainCNT,p2pPort,rpcPort uint){
	if !check_nodeNumber(nodeCNT) {
		fmt.Println("nodeCnt should be 1 or 4 or 7 or 10 or 13")
		fmt.Println(nodeCNT)
		os.Exit(1)
	}

	if chainCNT <1 || chainCNT > 4 {
		fmt.Println("chainCnt should be 1 - 4")
		os.Exit(1)
	}

	if p2pPort <= 10000 || p2pPort >= 60000 {
		fmt.Println("p2pPort should >= 10000 and  <=60000 ")
		os.Exit(1)
	}

	if rpcPort <= 10000 || p2pPort >= 60000 {
		fmt.Println("rpcPort should >= 10000 and  <=60000 ")
		os.Exit(1)
	}
}


func main(){
	flag.Parse()
	checkParams(nodeCNT,chainCNT,p2pPort,rpcPort)
	getInfo()
}
