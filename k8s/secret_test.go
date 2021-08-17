package k8s

import(
	"testing"
	"io/ioutil"
	"fmt"
)

func TestInitSecret(t *testing.T) {
	fileContent,err := ioutil.ReadFile("./secret.go")
	if err != nil{
		t.Error(err)
	}

	secret,err := secretInit("name","namespace","secret.go",fileContent)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(secret)
}
