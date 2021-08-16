package k8s

const secretTemplate = `{
    "apiVersion": "v1",
    "kind": "Secret",
    "metadata": {
        "name": "secret-tls"
    },
    "type": "kubernetes.io/tls",
    "data": {
    }
}`
