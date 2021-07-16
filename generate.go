package main

import(
	"cryptogen"
	//"localconf"
)



// reading configFile then generate certs in outputDir
func generate_certs(outputDir string,configFile string) error{
	if err := cryptogen.LoadCryptoGenConfig(configFile); err != nil{
		return err
	}
	cryptogen.OutputDir = outputDir
	if err := cryptogen.Generate(); err != nil{
		return err
	}

	return nil
}

