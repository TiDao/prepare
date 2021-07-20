package main

import(
	"cryptogen"
	"localconf"
)

func generate_certs(outputDir string) error{
	cryptogen.OutputDir = outputDir
	if err := cryptogen.Generate();err != nil{
		return err
	}
	return nil
}

func generate_config(initInfo *InitInfo,config *localconf.CMConfig,workDir string) error{

	if err := config.ReadFile("./config/chainmaker.yml"); err != nil{
		return err
	}

	if err := config.ReadFile("./config/log.yml"); err != nil{
		return err
	}

	config.LogConfig.SystemLog.LogLevelDefault = initInfo.LogLevel
	config.LogConfig.SystemLog.LogLevels["core"] = initInfo.LogLevel
	config.LogConfig.SystemLog.LogLevels["net"] = initInfo.LogLevel
	config.LogConfig.BriefLog.LogLevelDefault = initInfo.LogLevel
	config.LOgconfig.EventLog.LogLevelDefault = initInfo.LogLevel

	for i := 0; i < initInfo.ChainCNT; i++ {
		blockChianItem := localconf.BlockchainConfig{
			ChainId: string(i+1),
			Genesis: "/home/heyue/config" + string(i+1),
		}
		config.GetBlockChains().append(config.GetBlockChains(),
		localconf.BlockchainConfig{ChainId:initInfo.}
	}
}
