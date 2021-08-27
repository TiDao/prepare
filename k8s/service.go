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

func (chain *ChainMakerType)serviceInit() error {

	svc := &corev1.Service{}

	err := json.Unmarshal([]byte(svcTemplate),svc)
	if err != nil {
		return err
	}

	svc.ObjectMeta.Name = chain.NodeName
	svc.ObjectMeta.Namespace = chain.NameSpace
	svc.Spec.Selector["chainmaker"] =  chain.NodeName

	chain.Service = svc
	return nil
}

