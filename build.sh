#!/usr/bin/env bash

# set some environment variables
ORG_PATH="github.com/templexxx"
REPO_PATH="${ORG_PATH}/etcd"

GIT_SHA=$(git rev-parse --short HEAD || echo "GitNotFound")
if [[ -n "$FAILPOINTS" ]]; then
	GIT_SHA="$GIT_SHA"-FAILPOINTS
fi

# Set GO_LDFLAGS="-s" for building without symbols for debugging.
GO_LDFLAGS="$GO_LDFLAGS -X ${REPO_PATH}/version.GitSHA=${GIT_SHA}"

out="bin"

GO111MODULE=on CGO_ENABLED=0 go build $GO_BUILD_FLAGS \
		-installsuffix cgo \
		-ldflags "$GO_LDFLAGS" \
		-o "${out}/etcd" ${REPO_PATH}

GO111MODULE=on CGO_ENABLED=0 go build $GO_BUILD_FLAGS \
		-installsuffix cgo \
		-ldflags "$GO_LDFLAGS" \
		-o "${out}/etcdctl" ${REPO_PATH}/etcdctl
