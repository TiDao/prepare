package main

import (
	"cryptogen"
	//"encoding/json"
	"fmt"
	//"net/http"
	"strconv"
	"command"
)


func cryptogenCount(count int) {
	for i:=0; i< len(cryptogen.CryptoConfig.Item); i++{
		cryptogen.CryptoConfig.Item[i].Count = int32(count)
	}
}

func httpCheckPort(port *int,e *string,name string,min,max,defaultPort int){
	if !command.CheckPort(*port,min,max) {
		*e = fmt.Sprintf("%s %d not in %d-%d,please input again\n",name,*port,min,max)
	}else if  *port == 0{
		*port = defaultPort
	}
}

//check log level
func checkLogLevel(initInfo *command.InitInfo,e *InitError) {
	switch initInfo.LogLevel{
	case "DEBUG":
		return
	case "INFO":
		return
	case "WARN":
		return
	case "ERROR":
		return
	case "":
		initInfo.LogLevel = "INFO"
	default:
		e.LogLevel = fmt.Sprintf("%s not in [DEBUG|INFO(default)|WARN|ERROR],please input again.\n",initInfo.LogLevel)
	}
}

func checkConsensusType(initInfo *command.InitInfo,e *InitError){
	//check consensus type
	switch initInfo.ConsensusType {
	case 0:
		return
	case 1:
		return
	case 3:
		return
	case 4:
		return
	case 5:
		return
	default:
		e.ConsensusType = fmt.Sprintf("consensusType %d not in [0-SOLO,1-TBFT,3-HOTSTUFF,4-RAFT,5-DPOS],please input and check again\n",initInfo.ConsensusType)
	}
}

func checkNodeCNT(initInfo *command.InitInfo,e *InitError) {
	switch initInfo.NodeCNT {
	case 1:
		cryptogenCount(1)
	case 4:
		cryptogenCount(4)
	case 7:
		cryptogenCount(7)
	case 10:
		cryptogenCount(10)
	case 13:
		cryptogenCount(13)
	case 0:
		cryptogenCount(4)
	default:
		e.NodeCNT= fmt.Sprintf("NodeCNT %d not in [1,4,710,13]",initInfo.NodeCNT)
	}
}

func checkChainCNT(initInfo *command.InitInfo,e *InitError){
	switch initInfo.ChainCNT {
	case 1:
		return
	case 2:
		return
	case 3:
		return
	case 4:
		return
	case 0:
		initInfo.ChainCNT = 1
	default:
		e.ChainCNT = fmt.Sprintf("ChainCNT %d not in 1 - 4",initInfo.ChainCNT)
	}
}

func httpCheckNodeNamePrefix(initInfo *command.InitInfo,e *InitError) {
	if initInfo.NodeNamePrefix == "" {
		initInfo.NodeNamePrefix = "wx-org"
	}
}

//check the request initInfo item 
func httpCheckInfo(initInfo *command.InitInfo) error {

	//error clection
	var e = &InitError{}

	//check LogLevel
	checkLogLevel(initInfo,e)

	//check ConsensusType
	checkConsensusType(initInfo,e)

	//check NodeCNT 
	checkNodeCNT(initInfo,e)

	//check ChainCNT
	checkChainCNT(initInfo,e)

	//check NodeNamePrefix
	httpCheckNodeNamePrefix(initInfo,e)

	//check MonitorPort 
	httpCheckPort(&initInfo.MonitorPort,&e.MonitorPort,"MonitorPort",10000,60000,14320)

	//check PPorfPort 
	httpCheckPort(&initInfo.PProfPort,&e.PProfPort,"PPorfProt",10000,60000,24320)

	//check TrustedPort
	httpCheckPort(&initInfo.TrustedPort,&e.TrustedPort,"TrustedPort",10000,60000,24320)

	//check P2Port
	httpCheckPort(&initInfo.P2Port,&e.P2Port,"P2Port",10000,60000,11300)

	//check RpcPort
	httpCheckPort(&initInfo.RpcPort,&e.RpcPort,"RpcPort",10000,60000,12300)

	//create ORG ID (for example: wx-org1.chainmaker.org ...etc) DomainName for each Node
	for i:=1; i<=initInfo.NodeCNT;i++{
		orgId := "wx-org"+ strconv.Itoa(i) + "-chainmaker-org"
		initInfo.OrgIDs = append(initInfo.OrgIDs,orgId)
	}


	//if the ConsensusType == 0,the mod is solo the chain count and Node count must be 1
	//and the ORG ID must be wx-org.chainmaker.org
	if initInfo.ConsensusType == 0{
		initInfo.ChainCNT =1
		initInfo.NodeCNT = 1
		initInfo.OrgIDs = []string{"wx-org.chainmaker.org"}
	}

	if e.checkError() {
		return nil
	}else{
		return e
	}

}

