package k8s

import(
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/rest"
)



//the function be used where outside of kubernetes cluster
func CMDClientset(path string) (*kubernetes.Clientset error){

	config, err := clientcmd.BuildConfigFromFlags("",path)
	if err != nil {
		return nil,err
	}

	clientset,err := kubernets.NewForConfig(config)
	if err != nil{
		return nil,err
	}

	return clientset,nil
}

//use function be used where inside of kubernetes cluster
func RestClientset() (*kubernetes.Clientset,error){

	config,err := rest.InClusterConfig()
	if err != nil{
		return nil,err
	}

	client,err := kubernetes.NewForConfig(config)
	if err != nil{
		return nil,err
	}

	return clientset,nil
}
