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
	Service               *corev1.Service
	Deployment            *appsv1.Deployment
	PersistentVolumeClaim *corev1.PersistentVolumeClaim
	ConfigMaps            []*corev1.ConfigMap
	Secrets               []*corev1.Secret
}

//the function be used where outside of kubernetes cluster
func (chain *ChainMakerType) CMDClientset(path string) error {

	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return err
	}

	chain.clientset, err = kubernets.NewForConfig(config)
	if err != nil {
		return err
	}

	return nil
}

//use function be used where inside of kubernetes cluster
func (chain *ChainMakerType) RestClientset() error {

	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	return nil
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


func secretCreate(secrets []*corev1.Secret) error {
	for _,secret := range secrets {
		err := chain.clientset.CoreV1().Secret(chain.NameSpace).Create(context.TODO(),secret,metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func configMapCreate(configMaps []*corev1.ConfigMap) error {
	for _,configMap := range configMaps{
		err := chain.clientset.CoreV1().ConfigMap(chain.Namespace).Create(context.TODO(),configMap,metav1.CreateOptions{})
		if err != nil{
			return err
		}
	}

	return nil
}

func persistentVolumeClaimCreate(persistentVolumeClaim *corev1.PersistentVolumeClaim) error {
	err := chain.clientset.CoreV1().PersistentVolumeClaim(chain.NameSpace).Create(context.TODO(),persistentVolumeClaim,metav1.CreateOption{})
	if err != nil{
		return err
	}

	return nil
}

func deploymentCreate(deployment *appsv1.Deployment) error{
	err := chain.client.AppsV1().Deployment(chain.NameSpace).Create(context.TODO(),deployment,metav1.CreateOption{})
	if err != nil{
		return err
	}

	return nil
}

func serviceCreate(service *corev1.Service) error{
	err := chain.client.CoreV1().Service(chain.NameSpace).Create(context.TODO(),service,matev1.CreateOption{})
	if err != nil{
		return err
	}

	return nil
}


func secretDelete(secrets []*corev1.Secret) error {
	for _,secret := range secrets {
		err := chain.clientset.CoreV1().Secret(chain.NameSpace).Delete(context.TODO(),secret.ObjectMeta.Name,metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func configMapDelete(configMaps []*corev1.ConfigMap) error {
	for _,configMap := configMaps{
		err := chain.clientset.CoreV1().ConfigMap(chain.Namespace).Delete(context.TODO(),configMap.ObjectMeta.Name,metav1.DeleteOptions{})
		if err != nil{
			return err
		}
	}

	return nil
}

func persistentVolumeClaimDelete(name string) error {
	err := chain.clientset.CoreV1().PersistentVolumeClaim(chain.NameSpace).Delete(context.TODO(),name,metav1.DeleteOption{})
	if err != nil{
		return err
	}

	return nil
}

func deploymentDelete(name string) error{
	err := chain.client.AppsV1().Deployment(chain.NameSpace).Delete(context.TODO(),name,metav1.DeleteOption{})
	if err != nil{
		return err
	}

	return nil
}

func serviceDelete(name string) error{
	err := chain.client.CoreV1().Service(chain.NameSpace).Delete(context.TODO(),name,matev1.DeleteOption{})
	if err != nil{
		return err
	}

	return nil
}
