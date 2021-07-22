kubectl create configmap chainmaker-1-config --from-file=./chainmaker/wx-org1.chainmaker.org/config
kubectl create configmap chainmaker-2-config --from-file=./chainmaker/wx-org2.chainmaker.org/config
kubectl create configmap chainmaker-3-config --from-file=./chainmaker/wx-org2.chainmaker.org/config
kubectl create configmap chainmaker-4-config --from-file=./chainmaker/wx-org2.chainmaker.org/config
kubectl create secret generic chainmaker-ca --from-file=./chainmaker/ca
kubectl	create secret generic chainmaker-1-node --from-file=./chainmaker/wx-org1.chainmaker.org/node
kubectl	create secret generic chainmaker-2-node --from-file=./chainmaker/wx-org2.chainmaker.org/node
kubectl	create secret generic chainmaker-3-node --from-file=./chainmaker/wx-org3.chainmaker.org/node
kubectl	create secret generic chainmaker-4-node --from-file=./chainmaker/wx-org4.chainmaker.org/node
kubectl	create secret generic chainmaker-1-user --from-file=./chainmaker/wx-org1.chainmaker.org/user
kubectl	create secret generic chainmaker-2-user --from-file=./chainmaker/wx-org2.chainmaker.org/user
kubectl	create secret generic chainmaker-3-user --from-file=./chainmaker/wx-org3.chainmaker.org/user
kubectl	create secret generic chainmaker-4-user --from-file=./chainmaker/wx-org4.chainmaker.org/user
