PWD=$(pwd)
FILE_ARGOCD_BOOTSTRAP=/resources/argocd/bootstrap.yaml
SLEEP=5

if [[ ! $PWD$FILE_ARGOCD_BOOTSTRAP ]] ; then
	echo "❗You must execute the script from root git folder"
	exit
else
	echo "👍 $FILE_ARGOCD_BOOTSTRAP found"
fi

echo "👍 [All checks ok]"
echo "-------"

ansible-playbook ./playbook/openshift_gitops_deploy.yml

oc apply -f $PWD$FILE_ARGOCD_BOOTSTRAP -n openshift-gitops

## user=admin
## password
# kubectl -n argocd get secret argocd-cluster -o json | jq -r '.data["admin.password"]' | base64 -d
