#!/usr/bin/env bash
set -euo pipefail

# Ensure gomobile tools and golang.org/x/mobile packages are available in module mode.

ANDROID_API=${ANDROID_API:-23}

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")"/.. && pwd)"
cd "${REPO_ROOT}"

echo "[deps] Installing gomobile tools"
go install golang.org/x/mobile/cmd/gomobile@latest
go install golang.org/x/mobile/cmd/gobind@latest

echo "[deps] Ensuring golang.org/x/mobile modules present in go.mod"
GO111MODULE=on GOFLAGS=-mod=mod GOWORK=off GOPROXY="https://proxy.golang.org,direct" \
  go get golang.org/x/mobile@latest
GO111MODULE=on GOFLAGS=-mod=mod GOWORK=off GOPROXY="https://proxy.golang.org,direct" \
  go get golang.org/x/mobile/cmd/gobind@latest
GO111MODULE=on GOFLAGS=-mod=mod GOWORK=off GOPROXY="https://proxy.golang.org,direct" \
  go mod tidy

echo "[deps] Initializing gomobile"
gomobile init

echo "[deps] Ready"

