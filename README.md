This repository represents a Service Mesh workshop based on Istio. 

# Steps

## Start up

We can test the application by calling the api endpoint. It mandatory some steps if we're using minikube:

Firstly, we must execute the ```minikube tunnel``` command.

Secondly, we're going to request against the mesh, so we need to know the IP: 

```bash
IP=$(kubectl get svc -n istio-system istio-ingress -o json | jq -r '.status.loadBalancer.ingress[0].ip')
```

Once we have the correct IP, we can execute the request:

```bash
TOKEN_NORMAL_USER=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoibm9ybWFsIn0.laJaqfrA8WXGx3VOUaYilgT3j0aWT1VmDeb394zlwKw
curl -H 'Authorization:'$TOKEN_NORMAL_USER $IP/api/v1/films
```



# Architecture

![Architecture image](images/architecture.png)

## Components

This repository contains four applications that will be used to demonstrate Istio capabilities.

### film-frontend

### film-storage

### user-storage

### film-api
