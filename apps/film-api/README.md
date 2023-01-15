Microservice used to improve a legacy system

# Configuration
| Variable | Description |
| ------ | ------ |
|SERVICE_VERSION|Service version for info api|
|SERVICE_NAME|Service name for info api|
|SERVICE_BUILD_TIME|Build application time|
|FILM_SERVICE_URL|Film service url|
|USER_SERVICE_URL|User service url|

# API

This microservice exposes a few endpoints:

| Endpoint | Description |
| ------ | ------ |
|/api/v1/films|Retrieves a JWT token with user information, from it get the user films|
|/api/v1/health|Healthcheck uses to ensure that the application is running up rightly|
|/api/v1/info|Exposes the microservice information|

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

