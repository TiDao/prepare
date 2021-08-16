package k8s

import(
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"encoding/json"
)

const pvcTemplate = `{
    "apiVersion": "v1",
    "kind": "PersistentVolumeClaim",
    "metadata": {
        "name": "",
		"namespace": ""
    },
    "spec": {
        "resources": {
            "requests": {
				"storage": ""
            },
            "limits": {
				"storage": ""
            }
        },
        "accessModes": [
            "ReadWriteMany"
        ],
        "storageClassName": "nfs"
    }
}`

func initPVC(name string,namespace string,size string) (*corev1.PersistentVolumeClaim,error) {
	pvc := &corev1.PersistentVolumeClaim{}
	err := json.Unmarshal([]byte(pvcTempalte),&pvc)
	if err != nil{
		return nil,err
	}

	pvc.ObjectMeta.Name = name
	pvc.ObjectMeta.Namespace = namespace
	pvc.Spec.Resources.Requests["storage"] = resource.MustParse(size)
	pvc.Spec.Resuorces.Limit["storage"] = resource.MustParse(size)

	return pvc,nil
}
