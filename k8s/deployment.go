package k8s

import (
	"encoding/json"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"path/filepath"
	"path"
	"strings"
)

const deploymentTemplate = `{
	"apiVesion": "apps/v1",
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
								"name": "data-pvc",
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


func getNames(pathName string) []string{
	var fileNames []string

	files,_:=  filepath.Glob(filepath.Join(pathName,"*"))

	for _,file := range files {
		fileNames = append(fileNames,path.Base(file))
	}
	return fileNames
}

func (chain *ChainMakerType)deploymentInit(caPath,configPath,nodePath, userPath string) error {

	deployment := &appsv1.Deployment{}
	err := json.Unmarshal([]byte(deploymentTemplate), deployment)
	if err != nil {
		return err
	}
	configFileNames := getNames(configPath)
	appendVolume(volumeConfigMap, chain.NodeName, configFileNames, deployment)
	appendVolumeMount("/home/heyue/config", configFileNames, deployment)


	caFileNames := getNames(caPath)
	appendVolume(volumeSecret, chain.NodeName, caFileNames, deployment)
	appendVolumeMount("/home/heyue/ca", caFileNames, deployment)

	nodeFileNames := getNames(nodePath)
	appendVolume(volumeSecret, chain.NodeName, nodeFileNames, deployment)
	appendVolumeMount("/home/heyue/node", nodeFileNames, deployment)

	userFileNames := getNames(userPath)
	appendVolume(volumeSecret, chain.NodeName, userFileNames, deployment)
	appendVolumeMount("/home/heyue/user", userFileNames, deployment)

	deployment.ObjectMeta.Name = chain.NodeName
	deployment.ObjectMeta.Namespace = chain.NameSpace
	deployment.Spec.Selector.MatchLabels["chainmaker"] = chain.NodeName
	deployment.Spec.Template.ObjectMeta.Name = chain.NodeName
	deployment.Spec.Template.ObjectMeta.Labels["chainmaker"] = chain.NodeName
	for i,_ := range deployment.Spec.Template.Spec.Volumes{
		if deployment.Spec.Template.Spec.Volumes[i].Name == "data-pvc" {
			deployment.Spec.Template.Spec.Volumes[i].PersistentVolumeClaim.ClaimName = chain.NodeName
		}
	}

	chain.Deployment = deployment
	return nil
}

func appendVolume(Type volumeType, nodeName string, files []string, deployment *appsv1.Deployment) {

	volume := corev1.Volume{}

	for _, file := range files {

		var volumeName string

		nameSplits := strings.Split(file, ".")
		if len(nameSplits) == 2 {
			volumeName = nameSplits[0] + "-" + nameSplits[1]
		} else {
			volumeName = nameSplits[0] + "-" + nameSplits[1] + "-" + nameSplits[2]
		}

		switch Type {
		case volumeConfigMap:
			volume.ConfigMap = &corev1.ConfigMapVolumeSource{}
			volume.Name = volumeName
			volume.ConfigMap.Name = nodeName + "-" + volumeName
		case volumeSecret:
			volume.Secret = &corev1.SecretVolumeSource{}
			volume.Name = volumeName
			volume.Secret.SecretName = nodeName + "-" + volumeName
		}

		deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, volume)
	}
}

func appendVolumeMount(mountPath string, files []string, deployment *appsv1.Deployment) {

	volumeMount := corev1.VolumeMount{}

	for _, file := range files {

		var volumeName string

		nameSplits := strings.Split(file, ".")
		if len(nameSplits) == 2 {
			volumeName = nameSplits[0] + "-" + nameSplits[1]
		} else {
			volumeName = nameSplits[0] + "-" + nameSplits[1] + "-" + nameSplits[2]
		}

		volumeMount.Name = volumeName
		volumeMount.MountPath = filepath.Join(mountPath,file)
		volumeMount.SubPath = file

		deployment.Spec.Template.Spec.Containers[0].VolumeMounts = append(deployment.Spec.Template.Spec.Containers[0].VolumeMounts, volumeMount)
	}

}

