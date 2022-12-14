PWD=$(pwd)
FILE_ARGOCD_OPERATOR=/resources/operators/argocd-operator.yaml
FILE_MONGODB_OPERATOR=/resources/operators/mongodb-operator.yaml
FILE_ARGOCD_SERVER=/resources/argocd/server.yaml
FILE_ARGOCD_BOOTSTRAP=/resources/argocd/bootstrap.yaml
SLEEP=5

if [[ ! $PWD$FILE_ARGOCD_SERVER ]] ; then
	echo "âYou must execute the script from root git folder"
	exit
else
	echo "ð $FILE_ARGOCD_SERVER found"
fi

if [[ ! $PWD$FILE_ARGOCD_OPERATOR ]] ; then
	echo "âYou must execute the script from root git folder"
	exit
else
	echo "ð $FILE_ARGOCD_OPERATOR found"
fi

if [[ ! $PWD$FILE_MONGODB_OPERATOR ]] ; then
	echo "âYou must execute the script from root git folder"
	exit
else
	echo "ð $FILE_MONGODB_OPERATOR found"
fi

echo "ð [All checks ok]"
echo "-------"

minikube start --cpus=4 --memory='16g' --vm-driver=kvm2

curl -sL https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.22.0/install.sh | bash -s v0.22.0

kubectl create -f $PWD$FILE_ARGOCD_OPERATOR
kubectl create -f $PWD$FILE_MONGODB_OPERATOR

while  
	! kubectl -n operators wait \
		--for condition=established \
		--timeout=60s \
		crd/argocds.argoproj.io
do 
	echo "â Waiting for CRD creations"
	sleep $SLEEP 
done

while  
	! kubectl --namespace operators wait \
    	--for=condition=ready pod \
    	--selector=control-plane=controller-manager \
    	--timeout=60s
do 
	echo "â Waiting for ArgoCD controller"
	sleep $SLEEP 
done

kubectl create ns argocd
kubectl apply -f $PWD$FILE_ARGOCD_SERVER

kubectl apply -f $PWD$FILE_ARGOCD_BOOTSTRAP -n argocd

## user=admin
## password
# kubectl -n argocd get secret argocd-cluster -o json | jq -r '.data["admin.password"]' | base64 -d
