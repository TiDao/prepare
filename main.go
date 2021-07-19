package main

import(
	"flag"
	"fmt"
	//"os"
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



func main(){
	flag.Parse()

	var initInfo *InitInfo
	initInfo = getInfo()
	fmt.Println(initInfo)
}
