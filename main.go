package main

import(
	//"fmt"
	"log"
	"cryptogen"
	//"localconf"
)





func main(){
	err := cryptogen.LoadCryptoGenConfig("./config/crypto_config_template.yml")
	if err != nil{
		log.Fatal(err)
	}
	//cryptoConfig := cryptogen.GetCryptoGenConfig()
	//fmt.Println(cryptoConfig)
	if err := generate_certs("./test_output");err != nil{
		log.Fatal(err)
	}

	//return *InitInfo
	initInfo := getInfo()
	//fmt.Println(initInfo)


	for i := 0;i < initInfo.NodeCNT; i++ {
		err := generate_config(initInfo,i)
		if err != nil{
			log.Fatalf("generate config error: %v",err)
		}
	}

	//fmt.Println(initInfo)
}
