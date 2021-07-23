package main

import(
	//"fmt"
	"log"
	"cryptogen"
	//"localconf"
)

const outputDir = "./output/chainmaker"



func main(){
	err := cryptogen.LoadCryptoGenConfig("./config/crypto_config_template.yml")
	if err != nil{
		log.Fatal(err)
	}


	//get init info
	initInfo := getInfo()
	initInfo.DomainName = "test.svc.cluster.local"

	if err := generate_certs(initInfo);err != nil{
		log.Fatal(err)
	}


	//log.Println(initInfo)
	for i := 0;i < initInfo.NodeCNT; i++ {
		err := generate_config(initInfo,i)
		if err != nil{
			log.Fatalf("generate config error: %v",err)
		}
	}

	for i := 0;i < initInfo.NodeCNT; i++ {
		err := generate_genesis(initInfo,i)
		if err != nil{
			log.Fatalf("generate genesis config error: %v",err)
		}
	}

	//fmt.Println(initInfo)
}
