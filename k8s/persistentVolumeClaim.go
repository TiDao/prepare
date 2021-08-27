package k8s

import(
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"encoding/json"
	"log"
)

const pvcTemplate = `{
	"apiVesion": "v1",
	"kind": "PersistentVolumeClaim",
    "metadata": {
        "name": "",
		"namespace": ""
    },
    "spec": {
        "accessModes": [
            "ReadWriteMany"
        ],
        "resources": {
            "requests": {
            },
            "limits": {
            }
        },
        "storageClassName": "nfs"
    }
}`

func (chain *ChainMakerType)persistentVolumeClaimInit(size string) error {

	pvc := &corev1.PersistentVolumeClaim{}

	err := json.Unmarshal([]byte(pvcTemplate),&pvc)
	if err != nil{
		log.Println(pvc)
		return err
	}

	pvc.ObjectMeta.Name = chain.NodeName
	pvc.ObjectMeta.Namespace = chain.NameSpace
	pvc.Spec.Resources.Requests["storage"] = resource.MustParse(size)
	pvc.Spec.Resources.Limits["storage"] = resource.MustParse(size)

	chain.PersistentVolumeClaim = pvc
	return nil
}
