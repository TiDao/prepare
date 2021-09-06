package main

import (
	"fmt"
)

type InitError struct {
	LogLevel       string
	ConsensusType  string
	NodeCNT        string
	ChainCNT       string
	MonitorPort    string
	PProfPort      string
	TrustedPort    string
	P2Port         string
	RpcPort        string
	DomainName     string
	NodeNamePrefix string
}

func (e *InitError) checkError() bool {
	if e.LogLevel != "" {
		return false
	} else if e.ConsensusType != "" {
		return false
	} else if e.NodeCNT != "" {
		return false
	} else if e.ChainCNT != "" {
		return false
	} else if e.MonitorPort != "" {
		return false
	} else if e.PProfPort != "" {
		return false
	} else if e.TrustedPort != "" {
		return false
	} else if e.P2Port != "" {
		return false
	} else if e.RpcPort != "" {
		return false
	} else if e.DomainName != "" {
		return false
	} else if e.NodeNamePrefix != "" {
		return false
	}
	return true
}

func (e *InitError) Error() string {
	return fmt.Sprintf("\nLogLevel: %s\nConsensusType: %s\nNodeCNT: %s\nChainCNT: %s\nMonitorPort: %s\nPProfPort: %s\nTrustedPort: %s\nP2Port: %s\nRpcPort: %s\nDomainName: %s\n",
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
