Static user list. This application shows a user list with different type of subscriptions.

# Build
```zsh
VERSION=0.2
podman build \
    --no-cache \
    --build-arg version=$VERSION \
    -t quay.io/dborrego/user-storage:$VERSION \
    -f Containerfile
```

# API
