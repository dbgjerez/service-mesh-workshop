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
TOKEN_PREMIUM_USER=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoicHJlbWl1bSJ9.mtZhdDIN6fpmWV0pOFeGotL6UJwVkrQ5gkYk6FHiED8
curl -H 'Authorization:'$TOKEN_PREMIUM_USER $IP/api/v1/films
```

## Kiali

Kiali shows the information about the service mesh traffic. This information is helpful since we can see the Istio objects configuration, possible issues and the traffic dashboard. 

Working with Minikube, we can open the Kiali dashboard with the following command: 

```bash
istioctl dashboard kiali
```

Now, we can go to the Graph tab and select all the namespaces to see our traffic ecosystem.

![Architecture image](images/kiali-01.gif)

> NOTE: If you can't see traffic in the animation, you have to generate traffic recalling the service, for example: using the previous call with the ```watch``` command.

# Architecture

![Architecture image](images/architecture.png)

## Components

This repository contains four applications that will be used to demonstrate Istio capabilities.

### film-frontend

### film-storage

### user-storage

### film-api
