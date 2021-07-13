package localconf

import(
	Viper "github.com/spf13/viper"
	"testing"
	"fmt"
	"encoding/json"
	"bytes"
)

func formatStruct(t interface{}){
	bs,_ := json.Marshal(t)
	var out bytes.Buffer
	json.Indent(&out,bs,"","\t")
	fmt.Printf("%v\n",out.String())
}

func TestTypes(t *testing.T) {
	viper := Viper.New()
	viper.SetConfigName("chainmaker")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil{
		t.Fatalf("%v",err)
	}
	fmt.Println(viper.config)

	var c  CMConfig
	err := viper.Unmarshal(&c)
	formatStruct(c)
	if err != nil{
		t.Fatalf("%v",err)
	}

}


