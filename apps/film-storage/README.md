# Configuration
ENV_CONFIG_JSON_LOCATION


# Build
```zsh
SERVICE_NAME=film-storage
VERSION=$(semver info v)
podman build \
    --no-cache \
    --build-arg version=$VERSION \
    --build-arg serviceName=$SERVICE_NAME \
    -t quay.io/dborrego/$SERVICE_NAME:$VERSION \
    -f Containerfile
```

# API
