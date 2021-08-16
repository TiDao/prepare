package k8s

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"context"
)

type ChainMakerType struct {
	clientset             *kubernetes.Clientset
	NodeName              string
	NameSpace             string
	Storage               int
	Service               *corev1.Service
	Deployment            *appsv1.Deployment
	PersistentVolumeClaim *corev1.PersistentVolumeClaim
	ConfigMaps            []*corev1.ConfigMap
	Secrets               []*corev1.Secret
}

func NewChainMakerType(NodeName,NameSpace,path string) (*ChainMakerType,error){
	chainMaker := &ChainMakerType{}
	if path == "" {
		chainMaker.clientset,err := clientsetResr()
		if err != nil {
			return nil,err
		}
	}else{
		chainMaker.clientset,err := clientsetCMD(path)
		if err != nil{
			return nil,err
		}
	}
}


func (chain *ChainMakerType) NodeCreate() error{

	err := &k8sError{}
	err.secret = secretCreate(chain.Secrets)
	err.configMap = configMapCreate(chain.ConfigMaps)
	err.persistentVolumeClaim = persistVolumeClaimCreate(chain.PersistenVolumeClaim)
	err.deployment = deploymentCreate(chain.Deployment)
	err.service = serviceCreate(chain.Service)

	return err
}

func (chain *ChainMakerType) NodeDelete() error{

	err := &k8sError{}
	err.secret = secretDelete(chain.Secrets)
	err.configMap = configMapDelete(chain.ConfigMaps)
	err.persistentVolumeClaim = persistVolumeClaimDelete(chain.PersistenVolumeClaim)
	err.deployment = deploymentDelete(chain.Deployment)
	err.service = serviceDelete(chain.Service)

	return err
}
