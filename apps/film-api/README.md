Microservice used to improve a legacy system

# Configuration
| Variable | Description |
| ------ | ------ |
|SERVICE_VERSION|Service version for info api|
|SERVICE_NAME|Service name for info api|
|SERVICE_BUILD_TIME|Build application time|
|FILM_SERVICE_URL|Film service url|
|USER_SERVICE_URL|User service url|
|COMMENT_SERVICE_URL|Comment service url|

# API

This microservice exposes a few endpoints:

| Endpoint | Description |
| ------ | ------ |
|/api/v1/films|Retrieves a JWT token with user information, from it get the user films|
|/api/v1/health|Healthcheck uses to ensure that the application is running up rightly|
|/api/v1/info|Exposes the microservice information|


## Test
To call with a normal user, use the following jwt token:

```zsh
TOKEN_NORMAL_USER=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoibm9ybWFsIn0.laJaqfrA8WXGx3VOUaYilgT3j0aWT1VmDeb394zlwKw
HOST_API_FILMS=film-api.dev-backend.k8s/api/v1/films
curl -H 'Authorization:'$TOKEN_NORMAL_USER $HOST_API_FILMS
[{"id":4,"title":"The Hobbit - An Unexpected Journey","duration":0},{"id":5,"title":"The Hobbit - The Desolation of Smaug","duration":0},{"id":6,"title":"The Hobbit - The Battle of the Five Armies","duration":0}]
```

Another option is to call as a premium user. In this case, you have to use a premium user, for example:

```zsh
TOKEN_PREMIUM_USER=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoicHJlbWl1bSJ9.mtZhdDIN6fpmWV0pOFeGotL6UJwVkrQ5gkYk6FHiED8
curl -H 'Authorization:'$TOKEN_PREMIUM_USER film-api.dev-backend.k8s/api/v1/films
[{"id":1,"title":"The Lord of the Rings - The Fellowship of the Ring","duration":0},{"id":2,"title":"The Lord of the Rings - The Two Towers","duration":0},{"id":3,"title":"The Lord of the Rings - The Return of the King","duration":0},{"id":4,"title":"The Hobbit - An Unexpected Journey","duration":0},{"id":5,"title":"The Hobbit - The Desolation of Smaug","duration":0},{"id":6,"title":"The Hobbit - The Battle of the Five Armies","duration":0}]
```

# Build
```zsh
SERVICE_NAME=film-api
VERSION=$(semver info v)
SERVICE_BUILD_TIME=$(date '+%Y/%m/%d %H:%M:%S')
podman build \
    --no-cache \
    --build-arg version=$VERSION \
    --build-arg serviceName=$SERVICE_NAME \
    --build-arg buildTime=$SERVICE_BUILD_TIME \
    -t quay.io/dborrego/$SERVICE_NAME:$VERSION \
    -f Containerfile
```

