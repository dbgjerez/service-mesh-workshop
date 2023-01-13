Microservice used to improve a legacy system

# Configuration
| Variable | Description |
| ------ | ------ |
|SERVICE_VERSION|Service version for info api|
|SERVICE_NAME|Service name for info api|
|SERVICE_BUILD_TIME|Build application time|
|FILM_SERVICE_URL|Film service url|
|USER_SERVICE_URL|User service url|


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
