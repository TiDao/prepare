package k8s

import(
	corev1 "k8s.io/api/core/v1"
	"encoding/json"
	"strings"
)

const configMapTemplate = `{
    "apiVersion": "v1",
    "kind": "ConfigMap",
    "metadata": {
        "name": "",
        "namespace": ""
    },
    "data": {
    }
}`

func configMapInit(nodeName string,namespace string,fileName string,fileContent []byte) (*corev1.ConfigMap,error) {
	configMap := &corev1.ConfigMap{}
	err := json.Unmarshal([]byte(configMapTemplate),configMap)
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

	configMap.ObjectMeta.Name = nodeName + "-" + file
	configMap.ObjectMeta.Namespace = namespace
	configMap.Data[fileName] = string(fileContent)

	return configMap,nil
}
