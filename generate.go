package main

import (
	"cryptogen"
	"io/ioutil"
	"localconf"
	"path/filepath"
	"strconv"
	"os"
)

func generate_certs(outputDir string) error {
	cryptogen.OutputDir = outputDir
	if err := cryptogen.Generate(); err != nil {
		return err
	}
	return nil
}

func generate_config(initInfo *InitInfo, node int) error {

	//read config files and get CMConfig object
	config := &localconf.CMConfig{}
	if err := config.ReadFile("./config/chainmaker.yml"); err != nil{
		return err
	}
	if err := config.ReadFile("./config/log.yml"); err != nil{
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
			ChainId: string(i),
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
	err := os.MkdirAll(fileDir,0777)
	if err != nil {
		return err
	}
	filePath := filepath.Join(fileDir,"chainmaker.yml")
	err = config.WriteFile(filePath, 0664)
	if err != nil {
		return err
	}
	return nil
}
