package k8s

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"path"
	"path/filepath"
	"io/ioutil"
	//"log"
	//b64 "encoding/base64"
)

type ChainMakerType struct {
	clientset             *kubernetes.Clientset
	NodeName              string
	NameSpace             string
	Storage               string
	Service               *corev1.Service
	Deployment            *appsv1.Deployment
	PersistentVolumeClaim *corev1.PersistentVolumeClaim
	ConfigMaps            []*corev1.ConfigMap
	Secrets               []*corev1.Secret
}


func NewChainMakerType(clientConfigPath,nodeName,nameSpace string,storageSize string,caPath,configPath,nodePath,userPath string) (*ChainMakerType,error){
	chainMaker := &ChainMakerType{}
	if clientConfigPath == "" {
		err := chainMaker.clientsetRest()
		if err != nil {
			return nil,err
		}
	}else{
		err := chainMaker.clientsetCMD(clientConfigPath)
		if err != nil{
			return nil,err
		}
	}

	chainMaker.NodeName = nodeName
	chainMaker.NameSpace = nameSpace

	chainMaker.configMapsInit(configPath)
	chainMaker.secretsInit(caPath)
	chainMaker.secretsInit(nodePath)
	chainMaker.secretsInit(userPath)

	chainMaker.serviceInit()

	chainMaker.persistentVolumeClaimInit(storageSize + "Gi")
	chainMaker.deploymentInit(caPath,configPath,nodePath,userPath)

	return chainMaker,nil
}


func (chain *ChainMakerType) NodeCreate() error{

	err := &k8sError{}
	err.secret = chain.secretCreate()
	err.configMap = chain.configMapCreate()
	err.persistentVolumeClaim = chain.persistentVolumeClaimCreate()
	err.deployment = chain.deploymentCreate()
	err.service = chain.serviceCreate()

	if err.deployment != nil{
		return err
	}else if err.service != nil{
		return err
	}else if err.secret != nil{
		return err
	}else if err.configMap != nil{
		return err
	}else if err.persistentVolumeClaim != nil{
		return err
	}
	return nil
}

func (chain *ChainMakerType) NodeDelete() error{

	err := &k8sError{}
	err.deployment = chain.deploymentDelete()
	err.service = chain.serviceDelete()

	err.secret = chain.secretDelete()
	err.configMap = chain.configMapDelete()
	err.persistentVolumeClaim = chain.persistentVolumeClaimDelete()

	if err.deployment != nil{
		return err
	}else if err.service != nil{
		return err
	}else if err.secret != nil{
		return err
	}else if err.configMap != nil{
		return err
	}else if err.persistentVolumeClaim != nil{
		return err
	}
	return nil
}

func (chain *ChainMakerType) NodeGet() (*corev1.Service,error) {
	service,err := chain.serviceGet()
	if err != nil{
		return nil,err
	}

	return service,nil
}

func (chain *ChainMakerType) NodeList() ([]corev1.Service,error) {
	list,err := chain.serviceList()
	if err != nil {
		return nil,err
	}

	return list,nil
}

func (chain *ChainMakerType) configMapsInit(configMapPath string) error {
	filePaths,err := filepath.Glob(filepath.Join(configMapPath,"*"))
	if err != nil{
		return err
	}

	for _,filePath := range filePaths{
		data,err := ioutil.ReadFile(filePath)
		if err != nil{
			return err
		}
		fileName := path.Base(filePath)
		configMap,err := configMapInit(chain.NodeName,chain.NameSpace,fileName,data)
		if err != nil {
			return err
		}

		chain.ConfigMaps = append(chain.ConfigMaps,configMap)
	}

	return nil
}

func (chain *ChainMakerType) secretsInit(secretPath string) error{
	filePaths,err := filepath.Glob(filepath.Join(secretPath,"*"))
	if err != nil{
		return err
	}

	for _,filePath := range filePaths{
		data,err := ioutil.ReadFile(filePath)
		//dataBase64 :=  []byte(b64.StdEncoding.EncodeToString(data))
		dataBase64 :=  data
		if err != nil{
			return err
		}
		fileName := path.Base(filePath)
		secret,err := secretInit(chain.NodeName,chain.NameSpace,fileName,dataBase64)
		if err != nil {
			return err
		}

		chain.Secrets = append(chain.Secrets,secret)
	}

	return nil
}
