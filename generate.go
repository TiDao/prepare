package main

import (
	"cryptogen"
	"io/ioutil"
	"localconf"
	"os"
	"path/filepath"
	"strconv"
)

func generate_certs(initInfo *InitInfo, outputDir string) error {

	cryptogen.OutputDir = outputDir

	//if consensus type is solo change config Object
	if initInfo.ConsensusType == 0 {
		for i := 0; i < len(cryptogen.CryptoConfig.Item); i++ {
			cryptogen.CryptoConfig.Item[i].Count = 1
		}
	}

	if err := cryptogen.Generate(); err != nil {
		return err
	}

	return nil
}

func generate_config(initInfo *InitInfo, node int) error {

	//read config files and get CMConfig object
	config := &localconf.CMConfig{}
	if err := config.ReadFile("./config/chainmaker.yml"); err != nil {
		return err
	}
	if err := config.ReadFile("./config/log.yml"); err != nil {
		return err
	}
	certsPath := "./test_output"

	//config log
	config.LogConfig.SystemLog.LogLevelDefault = initInfo.LogLevel
	config.LogConfig.SystemLog.LogLevels["core"] = initInfo.LogLevel
	config.LogConfig.SystemLog.LogLevels["net"] = initInfo.LogLevel
	config.LogConfig.BriefLog.LogLevelDefault = initInfo.LogLevel
	config.LogConfig.EventLog.LogLevelDefault = initInfo.LogLevel

	//config chain Block list
	for i := 1; i <= initInfo.ChainCNT; i++ {
		blockChainItem := localconf.BlockchainConfig{
			ChainId: "chain" + strconv.Itoa(i),
			Genesis: "/home/heyue/config" + "bc" + strconv.Itoa(i) + ".yml",
		}

		blockChains := config.GetBlockChains()
		blockChains = append(blockChains, blockChainItem)
	}

	//config p2p net
	for i, orgId := range initInfo.OrgIDs {
		path := filepath.Join(certsPath, orgId, "node/common1.nodeid")
		nodeHash, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		seed := filepath.Join("/dns4", "chainmaker-"+strconv.Itoa(i), "tcp", strconv.Itoa(initInfo.P2Port), "p2p", string(nodeHash))

		config.NetConfig.Seeds = append(config.NetConfig.Seeds, seed)
	}

	//config rpc net
	config.RpcConfig.Port = initInfo.RpcPort

	//config monitor
	config.MonitorConfig.Port = initInfo.MonitorPort

	//config pprof
	config.PProfConfig.Port = initInfo.PProfPort

	fileDir := filepath.Join(certsPath, initInfo.OrgIDs[node], "config")
	err := os.MkdirAll(fileDir, 0777)
	if err != nil {
		return err
	}
	filePath := filepath.Join(fileDir, "chainmaker.yml")
	err = config.WriteFile(filePath, 0664)
	if err != nil {
		return err
	}
	return nil
}

func generate_genesis(initInfo *InitInfo, node int) error {
	//certsDir := "./test_output"
	var fileTemplate string
	switch initInfo.NodeCNT {
	case 1:
		fileTemplate = filepath.Join("./config", "bc_solo.yml")
	case 4 | 7:
		fileTemplate = filepath.Join("./config", "bc_4_7.yml")
	case 16:
		fileTemplate = filepath.Join("./config", "bc_16.yml")
	default:
		fileTemplate = filepath.Join("./config", "bc_10_13.yml")
	}

	for i := 1; i <= initInfo.ChainCNT; i++ {
		config := &localconf.ChainConfig{}
		if err := config.ReadFile(fileTemplate); err != nil {
			return err
		}

		config.ChainId = "chain" + strconv.Itoa(i)

		switch initInfo.ConsensusType {
		case 0:
			//ConsensusType_SOLO
			config.Consensus.Type = 0
		case 1:
			//ConsensusType_TBFT
			config.Consensus.Type = 1
		case 2:
			//ConsensusType_MBFT
			config.Consensus.Type = 2
		case 3:
			//ConsensusType_HOTSTUFF
			config.Consensus.Type = 3
		case 4:
			//ConsensusType_RAFT
			config.Consensus.Type = 4
		case 5:
			//ConsensusType_DPOS
			config.Consensus.Type = 5
		case 10:
			//ConsensusType_POW
			config.Consensus.Type = 10
		default:
			break
		}
	}

	for i := 0; i < initInfo.NodeCNT; i++ {
		//NodeHashPath := filepath.Join("./test_output",initInfo.OrgIDs[i],"node","common1.nodeid")

	}

	return nil
}
