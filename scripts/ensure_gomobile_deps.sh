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
# Add mobile and bind packages into the module graph so gobind-generated imports resolve
go get golang.org/x/mobile@latest
go get golang.org/x/mobile/bind@latest || true
go mod tidy

echo "[deps] Initializing gomobile"
gomobile init -androidapi=${ANDROID_API}

echo "[deps] Ready"

