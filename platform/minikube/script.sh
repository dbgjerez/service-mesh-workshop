PWD=$(pwd)
FILE_ARGOCD_SERVER=/resources/argocd/server.yaml

if [[ ! $PWD$FILE_ARGOCD_SERVER ]] ; then
	echo "❗You must execute the script from root git folder"
	exit
else
	echo "👍 $FILE_ARGOCD_SERVER found"
fi

echo "👍 [All checks ok]"
echo "-------"

minikube start --cpus=4 --memory='16g' --vm-driver=kvm2

curl -sL https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.22.0/install.sh | bash -s v0.22.0

kubectl create -f https://operatorhub.io/install/mongodb-operator.yaml
kubectl create -f https://operatorhub.io/install/argocd-operator.yaml

while  
	! kubectl -n operators wait \
		--for condition=established \
		--timeout=60s \
		crd/argocds.argoproj.io
do 
	echo "⌛ Waiting for CRD creations"
	sleep 1 
done

while  
	! kubectl --namespace operators wait \
    	--for=condition=ready pod \
    	--selector=control-plane=controller-manager \
    	--timeout=60s
do 
	echo "⌛ Waiting for ArgoCD controller"
	sleep 1 
done

kubectl create ns argocd
kubectl apply -f $PWD$FILE_ARGOCD_SERVER

## user=admin
## password
# kubectl -n argocd get secret argocd-cluster -o json | jq -r '.data["admin.password"]' | base64 -d
