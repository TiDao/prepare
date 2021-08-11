package main

import (
	"cryptogen"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func checkPort(port int, min int, max int) bool {
	if port >= min && port <= max {
		return true
	}
	return false
}

type InitError struct {
	LogLevel      string
	ConsensusType string
	NodeCNT       string
	ChainCNT      string
	MonitorPort   string
	PProfPort     string
	TrustedPort   string
	P2Port        string
	RpcPort       string
	DomainName    string
}

func (e *InitError) Error() string {
	fmt.Sprintf("\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s",
	e.LogLevel,
	e.ConsensusType,
	e.NodeCNT,
	e.ChainCNT,
	e.MonitorPort,
	e.PProfPort,
	e.TrustedPort,
	e.P2Port,
	e.RpcPort,
	e.DomainName)
}


func httpCheckInfo(initInfo *InitInfo) error {

	var e = &InitError{}


	//check log level
	switch initInfo.LogLevel{
	case "DEBUG":
		break
	case "INFO":
		break
	case "WARN":
		break
	case "ERROR":
		break
	case "":
		initInfo.LogLevel = "INFO"
	default:
		e.LogLevel = fmt.Sprintf("%s not in [DEBUG|INFO(default)|WARN|ERROR],please input again.\n",initInfo.LogLevel)
	}


	//check consensus type
	switch initInfo.ConsensusType {
	case 0:
		break
	case 1:
		break
	case 3:
		break
	case 4:
		break
	case 5:
		break
	default:
		e.ConsensusType = fmt.sprintf("%d not in [0-SOLO,1-TBFT,3-HOTSTUFF,4-RAFT,5-DPOS],please input and check again\n",initInfo.ConsensusType)
	}

	//check MonitorPort 
	if !checkPort(initInfo.MonitorPort,10000,60000) {
		e.MonitorPort = fmt.Sprintf("%d not in 10000-60000,please input again\n",
		initInfo.MonitorPort)
	}else if  initInfo.MonitorPort == 0{
		initInfo.MonitorPort = 14320
	}

	//check PPorfPort 
	if !checkPort(initInfo.PPorfPort,10000,60000) {
		e.MonitorPort = fmt.Sprintf("%d not in 10000-60000,please input again\n",
		initInfo.PPorfPort)
	}else if initInfo.PPorfPort == 0{
		initInfo.PPorfPort = 24330
	}

	//check TrustedPort
	if !checkPort(initInfo.TrustedPort,10000,60000) {
		e.MonitorPort = fmt.Sprintf("%d not in 10000-60000,please input again\n",
		initInfo.TrustedPort)
	}else if initInfo.TrustedPort == 0 {
		initInfo.PPorfPort = 13300
	}

	//check P2Port
	if !checkPort(initInfo.P2Port,10000,60000) {
		e.MonitorPort = fmt.Sprintf("%d not in 10000-60000,please input again\n",
		initInfo.P2Port)
	}else if init
	return initInfo
}
