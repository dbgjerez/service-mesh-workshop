minikube start --cpus=4 --memory='16g' --vm-driver=kvm2

NS=argocd
curl -sL https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.22.0/install.sh | bash -s v0.22.0

kubectl create -f https://operatorhub.io/install/mongodb-operator.yaml
kubectl create -f https://operatorhub.io/install/argocd-operator.yaml

kubectl create ns argocd
kubectl apply -f ../../resources/argocd/server.yaml

## user=admin
## password
# kubectl -n argocd get secret argocd-cluster -o json | jq -r '.data["admin.password"]' | base64 -d
