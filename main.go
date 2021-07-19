package main

import(
	"fmt"
	"log"
	"cryptogen"
)





func main(){
	err := cryptogen.LoadCryptoGenConfig("./config/crypto_config_template.yml")
	if err != nil{
		log.Fatal(err)
	}
	cryptoConfig := cryptogen.GetCryptoGenConfig()
	fmt.Println(cryptoConfig)
	if err := generate_certs("./test_output");err != nil{
		log.Fatal(err)
	}
	var initInfo *InitInfo
	initInfo = getInfo()
	fmt.Println(initInfo)
}
