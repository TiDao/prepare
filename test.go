package main

import(
	"io/ioutil"
	"fmt"
)

func main(){
	path := "output/chainmaker/wx-org1.chainmaker.org/config/bc1.yml"
	data,err := ioutil.ReadFile(path)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Printf("%+q\n",string(data))
}

