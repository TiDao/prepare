package k8s

import(
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"encoding/json"
	"path/filepath"
	"strings"
)


const deploymentTemplate = `{
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
        "name": "chainmaker-1",
        "namespace": "test"
    },
    "spec": {
        "replicas": 1,
        "selector": {
            "matchLabels": {
                "chainmaker": "chainmaker-1"
            }
        },
        "template": {
            "metadata": {
                "name": "chainmaker-1",
                "labels": {
                    "chainmaker": "chainmaker-1"
                }
            },
            "spec": {
                "securityContext": {
                    "runAsUser": 2021,
                    "runAsGroup": 2021,
                    "fsGroup": 2021
                },
                "imagePullSecrets": [
                    {
                        "name": "regcred"
                    }
                ],
                "containers": [
                    {
                        "image": "registry.docker.heyue/chainmaker:v1.0.0",
                        "name": "chainmaker",
                        "command": [
                            "/home/heyue/bin/chainmaker"
                        ],
                        "args": [
                            "start",
                            "-c",
                            "/home/heyue/config/chainmaker.yml"
                        ],
                        "volumeMounts":[
							{
								"name": "data-pvc"
								"mountPath": "/home/heyue/data"
							}
						]
                    }
                ],
                "volumes": [
                    {
                       "name": "data-pvc",
                       "persistentVolumeClaim": {
                            "claimName": null
                        }
                    }
                ]
            }
        }
    }
}`

type volumeType string

const (
	volumeConfigMap volumeType = "configMap"
	volumeSecret    volumeType = "secret"
)


func appendVolume(Type volumeType,nodeName string,files []string,deployment *appsv1.Deployment) {

	volume := corev1.Volume{}

	for _,file := range files{

		var volumeName string

		nameSplits := strings.Split(file,".")
		if len(nameSplits) == 2{
			volumeName = nameSplits[0] + "-" + nameSplits[1]
		}else{
			volumeName = nameSplits[0] + "-" + nameSplits[1] + "-" + nameSplits[2]
		}

		switch Type{
		case volumeConfigMap:
			volume.Name = volumeName
			volume.VolumeSource.ConfigMap.Name = nodeName + "-" + volumeName
		case volumeSecret:
			volume.Name = volumeName
			volume.VolumeSource.Secret.SecretName = nodeName + "-" + volumeName
		}

		deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, volume)
	}
}

func appendVolumeMount(mountPath string,files []string,deployment *appsv1.Deployment) {

	volumeMount = corev1.VolumeMount{}

	for _,file := range files{

		var volumeName string

		nameSplits := strings.Split(file,".")
		if len(nameSplits) == 2{
			volumeName = nameSplits[0] + "-" + nameSplits[1]
		}else{
			volumeName = nameSplits[0] + "-" + nameSplits[1] + "-" + nameSplits[2]
		}

		volumeMount.Name = volumeName
		volumeMount.MountPath = mountPath

		deployment.Sepc.Template.Spec.Containers[0].VolumeMOunts = append(deployment.Spec.Template.Spec.Containers[0].VolumeMounts,volumeMount)
	}

}

func deploymentInit(name,namespace,configPath,caPath,nodePath,userPath string) (*appsv1.Deployment,error) {
	deployment := &appsv1.Deployment{}
	err := json.Unmarshal([]byte(deploymentTemplate),deployment)
	if err != nil{
		return nil,err
	}

	configFileNames,err := filepath.Glob(filepath.Join(configPath,"*"))
	if err != nil {
		return nil,err
	}
	appendVolume(volumeConfigMap,name,configFileNames,deployment)
	appendVolumeMount("/home/heyue/config",configFileNames,deployment)


	caFileNames,err := filepath.Glob(filepath.Join(caPath,"*"))
	if err != nil {
		return nil,err
	}
	appendVolume(volumeSecret,name,caFileNames,deployment)
	appendVolume("/home/heyue/ca",caFileNames,deployment)

	nodeFileNames,err := filepath.Glob(filepath.Join(nodePath,"*"))
	if err != nil {
		return nil,err
	}
	appendVolume(volumeSecret,name,nodeFileNames,deployment)
	appendVolume("/home/heyue/node",nodeFileNames,deployment)

	userFileNames,err := filepath.Glob(filepath.Join(userPath,"*"))
	if err != nil {
		return nil,err
	}
	appendVolume(volumeSecret,name,userFileNames,deployment)
	appendVolume("/home/heyue/user",userFileNames,deployment)

}

