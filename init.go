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

func getInfo() {
	var initInfo = &InitInfo{
		LogLevel:      "INFO",
		ConsensusType: 1,
		MonitorPort:   14320,
		PprofPort:     24330,
		TrustedPort:   13300,
		P2Port:        11300,
		RpcPort:       12300,
		NodeCNT:       4,
		ChainCNT:      1,
	}

GetLogLevel:
	for {
		var loglevel string
		fmt.Printf("input log level you want[DEBUG|INFO(default)|WARN|ERROR]: ")
		fmt.Scanln(&loglevel)
		switch loglevel {
		case "DEBUG":
			//fmt.Println(loglevel)
			initInfo.LogLevel = loglevel
			break GetLogLevel
		case "INFO":
			//fmt.Println(loglevel)
			initInfo.LogLevel = loglevel
			break GetLogLevel
		case "WARN":
			//fmt.Println(loglevel)
			initInfo.LogLevel = loglevel
			break GetLogLevel
		case "ERROR":
			//fmt.Println(loglevel)
			initInfo.LogLevel = loglevel
			break GetLogLevel
		case "":
			//fmt.Println(loglevel)
			initInfo.LogLevel = "INFO"
			break GetLogLevel
		default:
			//fmt.Println(loglevel)
			fmt.Printf("%s not in [DEBUG|INFO(default)|WARN|ERROR],please input again.\n", loglevel)
			continue
		}
	}

getConsensus:
	for {
		var consensusType int
		fmt.Printf("input consensus type (0-SOLO(default),1-TBFT,3-HOTSTUFF,4-RAFT,5-DPOS): ")
		fmt.Scanln(&consensusType)
		switch consensusType {
		case 0:
			//fmt.Println(consensusType)
			initInfo.ConsensusType = consensusType
			break getConsensus
		case 1:
			//fmt.Println(consensusType)
			initInfo.ConsensusType = consensusType
			break getConsensus
		case 3:
			//fmt.Println(consensusType)
			initInfo.ConsensusType = consensusType
			break getConsensus
		case 4:
			//fmt.Println(consensusType)
			initInfo.ConsensusType = consensusType
			break getConsensus
		case 5:
			//fmt.Println(consensusType)
			initInfo.ConsensusType = consensusType
			break getConsensus
		default:
			//fmt.Println(consensusType)
			fmt.Printf("%s not in [DEBUG|INFO(default)|WARN|ERROR],please input again.\n", 
			consensusType)
			continue
		}
	}

}
