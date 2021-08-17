package k8s

import(
	corev1 "k8s.io/api/core/v1"
	//b64 "encoding/base64"
	"encoding/json"
	"strings"
)

const secretTemplate = `{
    "apiVersion": "v1",
    "kind": "Secret",
    "metadata": {
        "name": "",
		"namespace": ""
    },
    "type": "kubernetes.io/tls",
    "data": {
    }
}`

func secretInit(nodeName,namespace string,fileName string,fileContent []byte) (*corev1.Secret,error) {
	secret := &corev1.Secret{}
	err := json.Unmarshal([]byte(secretTemplate),secret)
	if err != nil{
		return nil,err
	}

	fileSplits := strings.Split(fileName,".")
	var file string
	if len(fileSplits) == 2{
		file = fileSplits[0] + "-" + fileSplits[1]
	}else{
		file = fileSplits[0] + "-" + fileSplits[1] + "-" + fileSplits[2]
	}

	secret.ObjectMeta.Name = nodeName + "-" + file
	secret.ObjectMeta.Namespace = namespace
	secret.Data[fileName] = fileContent

	return secret,nil
}
