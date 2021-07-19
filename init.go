package main

import (
	"fmt"
)

type InitInfo struct {
	LogLevel      string
	ConsensusType int
	MonitorPort   int
	PprofPort     int
	TrustedPort   int
	NodeCNT       int
	ChainCNT      int
	P2Port        int
	RpcPort       int
}

func checkPort(port int, min int, max int) bool {
	if port >= min && port <= max {
		return true
	}
	return false
}

func getInfo() *InitInfo{

	var initInfo = &InitInfo{
		LogLevel:      "INFO",
		ConsensusType: 1,
		NodeCNT:       4,
		ChainCNT:      1,
		MonitorPort:   14320,
		PprofPort:     24330,
		TrustedPort:   13300,
		P2Port:        11300,
		RpcPort:       12300,
	}

GetLogLevel:
	for {
		var loglevel string
		fmt.Printf("input log level you want[DEBUG|INFO(default)|WARN|ERROR]: ")
		fmt.Scanln(&loglevel)
		switch loglevel {
		case "DEBUG":
			//fmt.Printf(loglevel)
			initInfo.LogLevel = loglevel
			break GetLogLevel
		case "INFO":
			//fmt.Printf(loglevel)
			initInfo.LogLevel = loglevel
			break GetLogLevel
		case "WARN":
			//fmt.Printf(loglevel)
			initInfo.LogLevel = loglevel
			break GetLogLevel
		case "ERROR":
			//fmt.Printf(loglevel)
			initInfo.LogLevel = loglevel
			break GetLogLevel
		case "":
			//fmt.Printf(loglevel)
			initInfo.LogLevel = "INFO"
			break GetLogLevel
		default:
			//fmt.Printf(loglevel)
			fmt.Printf("%s not in [DEBUG|INFO(default)|WARN|ERROR],please input again.\n", loglevel)
			continue
		}
	}

getConsensus:
	for {
		var consensusType string
		fmt.Printf("input consensus type (0-SOLO,1-TBFT(default),3-HOTSTUFF,4-RAFT,5-DPOS): ")
		fmt.Scanln(&consensusType)
		switch consensusType {
		case "0":
			//fmt.Printf(consensusType)
			initInfo.ConsensusType = 0
			break getConsensus
		case "1":
			//fmt.Printf(consensusType)
			initInfo.ConsensusType = 1
			break getConsensus
		case "3":
			//fmt.Printf(consensusType)
			initInfo.ConsensusType = 3
			break getConsensus
		case "4":
			//fmt.Printf(consensusType)
			initInfo.ConsensusType = 4
			break getConsensus
		case "5":
			//fmt.Printf(consensusType)
			initInfo.ConsensusType = 5
			break getConsensus
		case "" :
			break getConsensus
		default:
			//fmt.Printf(consensusType)
			fmt.Printf("%s not in (0-SOLO(default),1-TBFT,3-HOTSTUFF,4-RAFT,5-DPOS),please input again.\n", consensusType)
			continue
		}
	}

getMonitorPort:
	for {
		var port int
		fmt.Printf("input Monitor Port[10000-60000,default:14320]:")
		fmt.Scanln(&port)
		if checkPort(port, 10000, 60000) {
			initInfo.MonitorPort = port
			break getMonitorPort
		} else if port == 0 {
			break getMonitorPort
		}
	}

getPprofPort:
	for {
		var port int
		fmt.Printf("input pprof Port[10000-60000,default:24330]:")
		fmt.Scanln(&port)
		if checkPort(port, 10000, 60000) {
			initInfo.PprofPort = port
			break getPprofPort
		} else if port == 0 {
			break getPprofPort
		}
	}

getTrustedPort:
	for {
		var port int
		fmt.Printf("input trusted Port[10000-60000,default:13300]:")
		fmt.Scanln(&port)
		if checkPort(port, 10000, 60000) {
			initInfo.TrustedPort = port
			break getTrustedPort
		} else if port == 0 {
			break getTrustedPort
		}
	}

getP2Port:
	for {
		var port int
		fmt.Printf("input P2P Port[10000-60000,default:11300]:")
		fmt.Scanln(&port)
		if checkPort(port, 10000, 60000) {
			initInfo.P2Port = port
			break getP2Port
		} else if port == 0 {
			break getP2Port
		}
	}
getRpcPort:
	for {
		var port int
		fmt.Printf("input RPC Port[10000-60000,default:112300]:")
		fmt.Scanln(&port)
		if checkPort(port, 10000, 60000) {
			initInfo.RpcPort = port
			break getRpcPort
		} else if port == 0 {
			break getRpcPort
		}
	}
getNodeCNT:
	for {
		var CNT int
		fmt.Printf("input node count number[1,4,7,10,13,defautl:4],: ")
		fmt.Scanln(&CNT)
		switch CNT {
		case 1:
			initInfo.NodeCNT = CNT
			break getNodeCNT
		case 4:
			initInfo.NodeCNT = CNT
			break getNodeCNT
		case 7:
			initInfo.NodeCNT = CNT
			break getNodeCNT
		case 10:
			initInfo.NodeCNT = CNT
			break getNodeCNT
		case 13:
			initInfo.NodeCNT = CNT
			break getNodeCNT
		case 0:
			break getNodeCNT
		default:
			fmt.Printf("node count should be 1 or 4 or 7 or 10 or 13")
			continue
		}
	}
getChainCNT:
	for {
		var CNT int
		fmt.Printf("input chain count number[1 - 4,default:1]: ")
		fmt.Scanln(&CNT)
		switch CNT {
		case 1:
			initInfo.ChainCNT = CNT
			break getChainCNT
		case 2:
			initInfo.ChainCNT = CNT
			break getChainCNT
		case 3:
			initInfo.ChainCNT = CNT
			break getChainCNT
		case 4:
			initInfo.ChainCNT = CNT
			break getChainCNT
		case 0:
			break getChainCNT
		default:
			fmt.Printf("chain count should be 1 - 4")
			continue
		}
	}
	return initInfo
}
