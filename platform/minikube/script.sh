minikube start --cpus=4 --memory='16g' --vm-driver=kvm2
minikube addons enable ingress

kubectl --namespace ingress-nginx wait \
    --for=condition=ready pod \
    --selector=app.kubernetes.io/component=controller \
    --timeout=120s

NS=argocd

kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

sleep 1

kubectl --namespace argocd wait \
    --for=condition=ready pod \
    --selector=app.kubernetes.io/component="argocd-application-controller-0" \
    --timeout=120s

## user=admin
## password
# kubectl -n argocd get secret argocd-initial-admin-secret -o json | jq -r ".data.password" | base64 -d

