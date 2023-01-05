# Error service

This service transforms generic error messages into specific messages. 

## Build

```zsh
VERSION=0.2
podman build \
    --no-cache \
    --build-arg version=$VERSION \
    -t b0rr3g0/error-service:$VERSION \
    -f Containerfile.run
```