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



//the function be used where outside of kubernetes cluster
func (chain *ChainMakerType) clientsetCMD(path string) error {

	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return err
	}

	chain.clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	return nil
}

//use function be used where inside of kubernetes cluster
func (chain *ChainMakerType) clientsetRest() error {

	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	chain.clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	return nil
}



func (chain *ChainMakerType)secretCreate(secrets []*corev1.Secret) error {
	for _,secret := range secrets {
		_,err := chain.clientset.CoreV1().Secrets(chain.NameSpace).Create(context.TODO(),secret,metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (chain *ChainMakerType)configMapCreate(configMaps []*corev1.ConfigMap) error {
	for _,configMap := range configMaps{
		_,err := chain.clientset.CoreV1().ConfigMaps(chain.NameSpace).Create(context.TODO(),configMap,metav1.CreateOptions{})
		if err != nil{
			return err
		}
	}

	return nil
}

func (chain *ChainMakerType)persistentVolumeClaimCreate(persistentVolumeClaim *corev1.PersistentVolumeClaim) error {
	_,err := chain.clientset.CoreV1().PersistentVolumeClaims(chain.NameSpace).Create(context.TODO(),persistentVolumeClaim,metav1.CreateOptions{})
	if err != nil{
		return err
	}

	return nil
}

func (chain *ChainMakerType)deploymentCreate(deployment *appsv1.Deployment) error{
	_,err := chain.clientset.AppsV1().Deployments(chain.NameSpace).Create(context.TODO(),deployment,metav1.CreateOptions{})
	if err != nil{
		return err
	}

	return nil
}

func (chain *ChainMakerType)serviceCreate(service *corev1.Service) error{
	_,err := chain.clientset.CoreV1().Services(chain.NameSpace).Create(context.TODO(),service,metav1.CreateOptions{})
	if err != nil{
		return err
	}

	return nil
}


func (chain *ChainMakerType)secretDelete(secrets []*corev1.Secret) error {
	for _,secret := range secrets {
		err := chain.clientset.CoreV1().Secrets(chain.NameSpace).Delete(context.TODO(),secret.ObjectMeta.Name,metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (chain *ChainMakerType)configMapDelete(configMaps []*corev1.ConfigMap) error {
	for _,configMap := range configMaps{
		err := chain.clientset.CoreV1().ConfigMaps(chain.NameSpace).Delete(context.TODO(),configMap.ObjectMeta.Name,metav1.DeleteOptions{})
		if err != nil{
			return err
		}
	}

	return nil
}

func (chain *ChainMakerType)persistentVolumeClaimDelete(name string) error {
	err := chain.clientset.CoreV1().PersistentVolumeClaims(chain.NameSpace).Delete(context.TODO(),name,metav1.DeleteOptions{})
	if err != nil{
		return err
	}

	return nil
}

func (chain *ChainMakerType)deploymentDelete(name string) error{
	err := chain.clientset.AppsV1().Deployments(chain.NameSpace).Delete(context.TODO(),name,metav1.DeleteOptions{})
	if err != nil{
		return err
	}

	return nil
}

func (chain *ChainMakerType)serviceDelete(name string) error{
	err := chain.clientset.CoreV1().Services(chain.NameSpace).Delete(context.TODO(),name,metav1.DeleteOptions{})
	if err != nil{
		return err
	}

	return nil
}
