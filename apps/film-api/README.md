Microservice used to improve a legacy system

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
