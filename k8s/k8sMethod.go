package k8s

import (
	//appsv1 "k8s.io/api/apps/v1"
	//corev1 "k8s.io/api/core/v1"
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



func (chain *ChainMakerType)secretCreate() error {
	for _,secret := range chain.Secrets {
		_,err := chain.clientset.CoreV1().Secrets(chain.NameSpace).Create(context.TODO(),secret,metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (chain *ChainMakerType)configMapCreate() error {
	for _,configMap := range chain.ConfigMaps{
		_,err := chain.clientset.CoreV1().ConfigMaps(chain.NameSpace).Create(context.TODO(),configMap,metav1.CreateOptions{})
		if err != nil{
			return err
		}
	}

	return nil
}

func (chain *ChainMakerType)persistentVolumeClaimCreate() error {
	_,err := chain.clientset.CoreV1().PersistentVolumeClaims(chain.NameSpace).Create(context.TODO(),chain.PersistentVolumeClaim,metav1.CreateOptions{})
	if err != nil{
		return err
	}

	return nil
}

func (chain *ChainMakerType)deploymentCreate() error{
	_,err := chain.clientset.AppsV1().Deployments(chain.NameSpace).Create(context.TODO(),chain.Deployment,metav1.CreateOptions{})
	if err != nil{
		return err
	}

	return nil
}

func (chain *ChainMakerType)serviceCreate() error{
	_,err := chain.clientset.CoreV1().Services(chain.NameSpace).Create(context.TODO(),chain.Service,metav1.CreateOptions{})
	if err != nil{
		return err
	}

	return nil
}


func (chain *ChainMakerType)secretDelete() error {
	for _,secret := range chain.Secrets {
		err := chain.clientset.CoreV1().Secrets(chain.NameSpace).Delete(context.TODO(),secret.ObjectMeta.Name,metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (chain *ChainMakerType)configMapDelete() error {
	for _,configMap := range chain.ConfigMaps{
		err := chain.clientset.CoreV1().ConfigMaps(chain.NameSpace).Delete(context.TODO(),configMap.ObjectMeta.Name,metav1.DeleteOptions{})
		if err != nil{
			return err
		}
	}

	return nil
}

func (chain *ChainMakerType)persistentVolumeClaimDelete() error {
	err := chain.clientset.CoreV1().PersistentVolumeClaims(chain.NameSpace).Delete(context.TODO(),chain.PersistentVolumeClaim.ObjectMeta.Name,metav1.DeleteOptions{})
	if err != nil{
		return err
	}

	return nil
}

func (chain *ChainMakerType)deploymentDelete() error{
	err := chain.clientset.AppsV1().Deployments(chain.NameSpace).Delete(context.TODO(),chain.Deployment.ObjectMeta.Name,metav1.DeleteOptions{})
	if err != nil{
		return err
	}

	return nil
}

func (chain *ChainMakerType)serviceDelete() error{
	err := chain.clientset.CoreV1().Services(chain.NameSpace).Delete(context.TODO(),chain.Service.ObjectMeta.Name,metav1.DeleteOptions{})
	if err != nil{
		return err
	}

	return nil
}

func (chain *ChainMakerType)List() ([]string,error){
	var serviceNames []string
	services,err := chain.clientset.CoreV1().Services(chain.NameSpace).List(context.TODO(),metav1.ListOptions{})
	if err != nil{
		return nil,err
	}

	for _,service := range services.Items {
		serviceNames = append(serviceNames,service.ObjectMeta.Name)
	}

	return serviceNames,nil
}
