package main

import(
	"testing"
)

var initInfoTest = &InitInfo{
		LogLevel:      "INFO",
		ConsensusType: 1,
		NodeCNT:       4,
		ChainCNT:      1,
		MonitorPort:   14320,
		PProfPort:     24330,
		TrustedPort:   13300,
		P2Port:        11300,
		RpcPort:       12300,
		OrgIDs:        []string{"wx-org1.chainmaker.org","wx-org2.chainmaker.org","wx-org3.chainmaker.org","wx-org4.chainmaker.org"},
}

func TestGenerate_genesis(t *testing.T) {
	for i:=0; i < initInfoTest.NodeCNT; i++ {
		if err := generate_genesis(initInfoTest,i);err !=nil{
			t.Errorf("error run generate_genessis() in %d round,the error: %v ",i,err)
		}
	}
}
