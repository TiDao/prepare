package main

import(
	//"fmt"
	"log"
	"cryptogen"
	"command"
	"flag"
	"net/http"
	//"localconf"
)

var outputDir string
var commandFlag bool



func init(){
	flag.BoolVar(&commandFlag,"-command",false,"使用命令行生成创建")
	flag.StringVar(&outputDir,"-output","./output/chainmaker","配置和证书输出路径")
}


func main(){
	flag.Parse()
	if commandFlag {
		commandModle()
	}else{
		serverModle()
	}
}

func serverModle(){
	http.HandleFunc("/create",CreateChain)
	http.HandleFunc("/delete",DeleteChain)
	//http.HandleFunc("/list",ListChain)


	log.Println("start server and listen 0.0.0.0:10001")
	err := http.ListenAndServe("0.0.0.0:10001",nil)
	if err != nil{
		log.Fatal(err)
	}
}

func commandModle() {
	err := cryptogen.LoadCryptoGenConfig("./config/crypto_config_template.yml")
	if err != nil{
		log.Fatal(err)
	}
	//get init info
	initInfo := command.GetInfo()

	initInfo.DomainName = "test.svc.cluster.local"
	//log.Println(initInfo)

	if err := generate_certs(initInfo);err != nil{
		log.Fatal(err)
	}


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
