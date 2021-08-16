package k8s
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
                        "volumeMounts": null
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
