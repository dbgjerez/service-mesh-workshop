# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17-buster AS build

WORKDIR /app

COPY src/go.mod ./
COPY src/go.sum ./
RUN go mod download

COPY src ./

RUN CGO_ENABLED=0 go build -o /go/bin/app

##
## Deploy
##
FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=build --chown=nonroot:nonroot /go/bin/app /

EXPOSE 8080
USER nonroot:nonroot
ENV GIN_MODE=release

CMD ["/app"]