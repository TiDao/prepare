package k8s

import(
	corev1 "k8s.io/api/core/v1"
	"encoding/json"
)

const svcTemplate = `{
    "apiVersion": "v1",
    "kind": "Service",
    "metadata": {
        "name": "",
        "namespace": ""
    },
    "spec": {
        "type": "LoadBalancer",
        "ports": [
            {
                "name": "p2p",
                "port": 11300,
                "targetPort": 11300
            },
            {
                "name": "rpc",
                "port": 12300,
                "targetPort": 12300
            },
            {
                "name": "trusted",
                "port": 13300,
                "targetPort": 13300
            },
            {
                "name": "pprof",
                "port": 24330,
                "targetPort": 24330
            },
            {
                "name": "monitor",
                "port": 14320,
                "targetPort": 14320
            }
        ],
        "selector": {
            "chainmaker": ""
        }
    }
}`

func initSVC(name string,namespace string) (*corev1.Secret, error) {
	svc := &corev1.Secret
	err := json.Unmarshal([]byte(svcTemplate),svc)
	if err != nil {
		return nil,err
	}

	svc.ObjectMeta.Name = name
	svc.ObjectMeta.Namespace = namespace
	svc.Spec.Selector["chainmaker"] =  name

	return svc,nil
}

