package main

import(
	"net/http"
	"encoding/json"
	"command"
	"fmt"
	"log"
	"k8s"
	"path/filepath"
)

func CreateChain(w http.ResponseWriter,r *http.Request) {

	initInfo := &command.InitInfo{}

	if err := json.NewDecoder(r.Body).Decode(initInfo);err != nil{
		log.Println(err)
		fmt.Fprintf(w,"json unmarshal error: %s",err.Error())
		return
	}

	generateChain(initInfo,&w)

	chains,err :=  generateChainMakerType("/home/magatron/.kube/config",initInfo,"test","output")
	if err != nil{
		log.Println(err)
		fmt.Fprintf(w,"generate ChainMakerType error: %s",err.Error())
		return
	}

	for _,chain := range chains{
		if err := chain.NodeCreate();err != nil{
			log.Println(err)
			fmt.Fprintf(w,"create node error: %s",err.Error())
		}
	}
}

func DeleteChain(w http.ResponseWriter,r *http.Request){

	initInfo := &command.InitInfo{}

	if err := json.NewDecoder(r.Body).Decode(initInfo);err != nil{
		log.Println(err)
		fmt.Fprintf(w,"json unmarshal error: %s",err.Error())
		return
	}

	chains,err :=  generateChainMakerType("/home/magatron/.kube/config",initInfo,"test","output")
	if err != nil{
		log.Println(err)
		fmt.Fprintf(w,"generate ChainMakerType error: %s",err.Error())
		return
	}

	for _,chain := range chains{
		if err := chain.NodeDelete();err != nil{
			log.Println(err)
			fmt.Fprintf(w,"create node error: %s",err.Error())
		}
	}
}

func generateChain(initInfo *command.InitInfo,w *http.ResponseWriter) {

	if err := httpCheckInfo(initInfo);err != nil {
		log.Println(err)
		fmt.Fprintf(*w,"check request body error: %s",err.Error())
	}

	if err := generate_certs(initInfo); err != nil{
		log.Println(err)
		fmt.Fprintf(*w,"generate certs error: %s",err.Error())
	}

	for i:=0;i<initInfo.NodeCNT;i++{
		if err := generate_config(initInfo,i);err != nil{
			log.Println(err)
			fmt.Fprintf(*w,"generate config error: %s",err.Error())
		}
	}

	for i:=0;i<initInfo.NodeCNT;i++{
		if err := generate_genesis(initInfo,i); err != nil{
			log.Println(err)
			fmt.Fprintf(*w,"generate gensis error: %s",err.Error())
		}
	}
}

func generateChainMakerType(clientConfigPath string,initInfo *command.InitInfo,namespace string,outputDir string) ([]*k8s.ChainMakerType,error){

	var  chains []*k8s.ChainMakerType

	for _,name := range initInfo.OrgIDs {
		caPath := filepath.Join(outputDir,"ca")
		configPath := filepath.Join(outputDir,name,"config")
		nodePath := filepath.Join(outputDir,name,"node")
		userPath := filepath.Join(outputDir,name,"user")
		chain,err := k8s.NewChainMakerType(clientConfigPath,name,namespace,initInfo.StorageSize,caPath,configPath,nodePath,userPath)
		if err != nil{
			return nil,err
		}

		chains = append(chains,chain)
	}

	return chains,nil
}



