package main

import (
	common "chainmaker.org/chainmaker-go/pb/protogo/common"
	"cryptogen"
	"fmt"
	"io/ioutil"
	"localconf"
	"os"
	"path/filepath"
	"strconv"
)

func generate_certs(initInfo *InitInfo) error {

	cryptogen.OutputDir = outputDir

	//if consensus type is solo change config Object
	if initInfo.ConsensusType == 0 {
		for i := 0; i < len(cryptogen.CryptoConfig.Item); i++ {
			cryptogen.CryptoConfig.Item[i].Count = 1
		}
	}
	for i,_ := range cryptogen.CryptoConfig.Item{
		cryptogen.CryptoConfig.Item[i].HostName = initInfo.NodeNamePrefix
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
	certsPath := outputDir

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
			Genesis: "/home/heyue/config" + "/bc" + strconv.Itoa(i) + ".yml",
		}

		config.BlockChainConfig = append(config.BlockChainConfig, blockChainItem)
	}

	//config node
	config.NodeConfig.OrgId = initInfo.OrgIDs[node]

	//config p2p net
	config.NetConfig.ListenAddr = "/ip4/0.0.0.0/tcp/"+strconv.Itoa(initInfo.P2Port)
	for i, orgId := range initInfo.OrgIDs {
		path := filepath.Join(certsPath, orgId, "node/consensus1.nodeid")
		nodeHash, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		seed := filepath.Join("/dns4", initInfo.OrgIDs[i] + "."+initInfo.DomainName ,"tcp", strconv.Itoa(initInfo.P2Port), "p2p", string(nodeHash))

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
	var fileTemplate string
	fileTemplate = filepath.Join("./config", "bc_template.yml")
	//switch initInfo.NodeCNT {
	//case 1:
	//	fileTemplate = filepath.Join("./config", "bc_solo.yml")
	//case 4 | 7:
	//	fileTemplate = filepath.Join("./config", "bc_4_7.yml")
	//case 16:
	//	fileTemplate = filepath.Join("./config", "bc_16.yml")
	//default:
	//	fileTemplate = filepath.Join("./config", "bc_10_13.yml")
	//}

	for j := 1; j <= initInfo.ChainCNT; j++ {
		config := &localconf.ChainConfig{}
		if err := config.ReadFile(fileTemplate); err != nil {
			return err
		}

		config.ChainId = "chain" + strconv.Itoa(j)

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
			totalValue := strconv.Itoa(initInfo.NodeCNT * 2500000)
			nodeHashPath := filepath.Join(outputDir, initInfo.OrgIDs[0], "node", "consensus1.nodeid")
			nodeHash, err := ioutil.ReadFile(nodeHashPath)
			if err != nil {
				return err
			}
			totalKeyValuePair := &common.KeyValuePair{
				Key:   fmt.Sprintf("erc20.total"),
				Value: totalValue,
			}
			ownerKeyValuePair := &common.KeyValuePair{
				Key:   fmt.Sprintf("erc20.owner"),
				Value: string(nodeHash),
			}
			accountKeyValuePair := &common.KeyValuePair{
				Key:   fmt.Sprintf("erc20.account:SYSTEM_CONTRACT_DPOS_STAKE"),
				Value: totalValue,
			}
			epochValidatorNumKeyValuePair := &common.KeyValuePair{
				Key:   fmt.Sprintf("stake.epochValidatorNum"),
				Value: strconv.Itoa(initInfo.NodeCNT),
			}

			config.Consensus.DposConfig = append(config.Consensus.DposConfig, totalKeyValuePair)
			config.Consensus.DposConfig = append(config.Consensus.DposConfig, ownerKeyValuePair)
			config.Consensus.DposConfig = append(config.Consensus.DposConfig, accountKeyValuePair)
			config.Consensus.DposConfig = append(config.Consensus.DposConfig, epochValidatorNumKeyValuePair)
		case 10:
			//ConsensusType_POW
			config.Consensus.Type = 10
		default:
			break
		}

		for i := 0; i < initInfo.NodeCNT; i++ {
			nodeHashPath := filepath.Join(outputDir, initInfo.OrgIDs[i], "node", "common1.nodeid")
			nodeHash, err := ioutil.ReadFile(nodeHashPath)
			if err != nil {
				return err
			}

			orgConfig := &localconf.OrgConfig{
				OrgId:  initInfo.OrgIDs[i],
				NodeId: []string{string(nodeHash)},
			}
			config.Consensus.Nodes = append(config.Consensus.Nodes, orgConfig)

			candidateKeyValuePair := &common.KeyValuePair{
				Key:   fmt.Sprintf("stake.candidate:\"{org%d_peeraddr}\"", i+1),
				Value: "2500000",
			}
			config.Consensus.DposConfig = append(config.Consensus.DposConfig, candidateKeyValuePair)

			nodeIDKeyValuePair := &common.KeyValuePair{
				Key:   fmt.Sprintf("stake.nodeID:\"{org%d_peeraddr}\"", i+1),
				Value: string(nodeHash),
			}
			config.Consensus.DposConfig = append(config.Consensus.DposConfig, nodeIDKeyValuePair)

			trustRootConfig := &localconf.TrustRootConfig{
				OrgId: initInfo.OrgIDs[i],
				Root:  fmt.Sprintf("/home/heyue/ca/ca%d.crt", i+1),
			}
			config.TrustRoots = append(config.TrustRoots, trustRootConfig)

		}

		outputFile := filepath.Join(outputDir, initInfo.OrgIDs[node], "config", fmt.Sprintf("bc%d.yml", j))
		err := config.WriteFile(outputFile, 0664)
		if err != nil {
			return err
		}
	}

	return nil
}
