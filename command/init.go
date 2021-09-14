package command

import (
	"cryptogen"
	"fmt"
	"strconv"
)

type InitInfo struct {
	LogLevel       string   `json:"logLevel,omitempty"`
	ConsensusType  int      `json:"consensusType,omitempty"`
	NodeCNT        int      `json:"nodeCount,omitempty"`
	ChainCNT       int      `json:"chainCount,omitempty"`
	MonitorPort    int      `json:"monitorPort,omitempty"`
	PProfPort      int      `json:"pprofPort,omitempty"`
	TrustedPort    int      `json:"trustedPort,omitempty"`
	P2Port         int      `json:"p2pPort,omitempty"`
	RpcPort        int      `json:"rpcPort,omitempty"`
	StorageSize    string   `json:"storageSize"`
	DomainName     string   `json:"domainName"`
	NodeNamePrefix string   `json:"nodeNamePrefix"`
	OrgIDs         []string
}

func GetInfo() *InitInfo {
	var initInfo = &InitInfo{
		LogLevel:      "INFO",
		ConsensusType: 1,
		NodeCNT:       4,
		ChainCNT:      1,
		MonitorPort:   14320,
		PProfPort:     24330,
		TrustedPort:   13300,
		P2Port:        11300,
		RpcPort:       12300,
		DomainName:    "test.svc.cluster.local",
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
		fmt.Printf("input consensus type [0-SOLO,1-TBFT,3-HOTSTUFF,4-RAFT,5-DPOS,default: 1]: ")
		fmt.Scanln(&consensusType)
		switch consensusType {
		case "0": //fmt.Printf(consensusType)
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
		case "5":
			//fmt.Printf(consensusType)
			initInfo.ConsensusType = 5
			break getConsensus
		case "":
			initInfo.ConsensusType = 1
			break getConsensus
		default:
			//fmt.Printf(consensusType)
			fmt.Printf("%s not in [0-SOLO,1-TBFT,3-HOTSTUFF,4-RAFT,5-DPOS],please input again.\n", consensusType)
			continue
		}
	}

getMonitorPort:
	for {
		var port int
		fmt.Printf("input Monitor Port[10000-60000,default:14320]:")
		fmt.Scanln(&port)
		if CheckPort(port, 10000, 60000) {
			initInfo.MonitorPort = port
			break getMonitorPort
		} else if port == 0 {
			initInfo.MonitorPort = 14320
			break getMonitorPort
		}
	}

getPProfPort:
	for {
		var port int
		fmt.Printf("input pprof Port[10000-60000,default:24330]:")
		fmt.Scanln(&port)
		if CheckPort(port, 10000, 60000) {
			initInfo.PProfPort = port
			break getPProfPort
		} else if port == 0 {
			initInfo.PProfPort = 24330
			break getPProfPort
		}
	}

getTrustedPort:
	for {
		var port int
		fmt.Printf("input trusted Port[10000-60000,default:13300]:")
		fmt.Scanln(&port)
		if CheckPort(port, 10000, 60000) {
			initInfo.TrustedPort = port
			break getTrustedPort
		} else if port == 0 {
			initInfo.TrustedPort = 13300
			break getTrustedPort
		}
	}

getP2Port:
	for {
		var port int
		fmt.Printf("input P2P Port[10000-60000,default:11300]:")
		fmt.Scanln(&port)
		if CheckPort(port, 10000, 60000) {
			initInfo.P2Port = port
			break getP2Port
		} else if port == 0 {
			initInfo.P2Port = 11300
			break getP2Port
		}
	}

getRpcPort:
	for {
		var port int
		fmt.Printf("input RPC Port[10000-60000,default:12300]:")
		fmt.Scanln(&port)
		if  CheckPort(port, 10000, 60000) {
			initInfo.RpcPort = port
			break getRpcPort
		} else if port == 0 {
			initInfo.RpcPort = 12300
			break getRpcPort
		}
	}

getNodeCNT:
	for {
		var CNT string
		fmt.Printf("input node count number[1,4,7,10,13,defautl:4]: ")
		fmt.Scanln(&CNT)
		switch CNT {
		case "1":
			initInfo.NodeCNT = 1
			for i := 0; i < len(cryptogen.CryptoConfig.Item); i++ {
				cryptogen.CryptoConfig.Item[i].Count = 1
			}
			break getNodeCNT
		case "4":
			initInfo.NodeCNT = 4
			for i := 0; i < len(cryptogen.CryptoConfig.Item); i++ {
				cryptogen.CryptoConfig.Item[i].Count = 4
			}
			break getNodeCNT
		case "7":
			initInfo.NodeCNT = 7
			for i := 0; i < len(cryptogen.CryptoConfig.Item); i++ {
				cryptogen.CryptoConfig.Item[i].Count = 7
			}
			break getNodeCNT
		case "10":
			initInfo.NodeCNT = 10
			for i := 0; i < len(cryptogen.CryptoConfig.Item); i++ {
				cryptogen.CryptoConfig.Item[i].Count = 10
			}
			break getNodeCNT
		case "13":
			initInfo.NodeCNT = 13
			for i := 0; i < len(cryptogen.CryptoConfig.Item); i++ {
				cryptogen.CryptoConfig.Item[i].Count = 13
			}
			break getNodeCNT
		case "":
			initInfo.NodeCNT = 4
			for i := 0; i < len(cryptogen.CryptoConfig.Item); i++ {
				cryptogen.CryptoConfig.Item[i].Count = 4
			}
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
			initInfo.ChainCNT = 1
			break getChainCNT
		default:
			fmt.Printf("chain count should be 1 - 4")
			continue
		}
	}

getNodeNamePrefix:
	for {
		var name string
		fmt.Printf("input Node name prefix[for example: wx-org(default)]: ")
		fmt.Scanln(&name)
		if name == "" {
			initInfo.NodeNamePrefix = "wx-org"
			break getNodeNamePrefix
		} else {
			initInfo.NodeNamePrefix = name
			break getNodeNamePrefix
		}
	}

	for i := 1; i <= initInfo.NodeCNT; i++ {
		orgId := initInfo.NodeNamePrefix + strconv.Itoa(i) + "-chainmaker-org"
		initInfo.OrgIDs = append(initInfo.OrgIDs, orgId)
	}

	if initInfo.ConsensusType == 0 {
		initInfo.ChainCNT = 1
		initInfo.NodeCNT = 1
		initInfo.OrgIDs = []string{"wx-org-chainmaker-org"}
		return initInfo
	}

	return initInfo
}
