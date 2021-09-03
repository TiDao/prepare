package httpnet

import (
	"fmt"
)


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
	return fmt.Sprintf("\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s",
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
