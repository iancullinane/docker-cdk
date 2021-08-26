# Script executed when running the build container
# https://github.turbine.com/MGP-Server/docker-go-build

# The default gocache location is inside the container's filesystem, not the host's.
# Change the build cache location to a shared location.
export GOCACHE=$(pwd)/.gobuildcache

# Build the binary to /bin/app
# CGO_ENABLED=0 means build the binary with cgo instead of go

BUILD_NUMBER=0
BUILD_VERSION=$(git rev-parse --short HEAD)
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.VersionString=${BUILD_NUMBER}-${BUILD_VERSION}" -installsuffix cgo -o bin/docker-cdk src/main/*.go

