package main

import(
	"fmt"
	"log"
	"cryptogen"
)





func main(){
	cryptoConfig := cryptogen.GetCryptoGenConfig()
	fmt.Println(cryptoConfig)
	if err := generate_certs("./test_output","./config/crypto_config_template.yml");err != nil{
		log.Fatal(err)
	}
	var initInfo *InitInfo
	initInfo = getInfo()
	fmt.Println(initInfo)
}
