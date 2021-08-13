package k8s

import(
	"fmt"
)

type k8sError struct {
	secret error
	configMap error
	persistentVolumeClaim error
	deployment error
	service error
}

func (e *k8sError) Error() string{
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\n",e.secret,e.configMap,e.persistentVolumeClaim,e.deployment,e.service)
}

