#!/usr/bin/env bash
set -euo pipefail

# Manual build script:
# 1) Build gomobile AAR from ./termbridge
# 2) Place AAR into Flutter app's android/app/libs
# 3) Build Flutter APK (release by default)

APP_DIR=${APP_DIR:-"flutter_app"}
ANDROID_API=${ANDROID_API:-23}
TARGETS=${TARGETS:-"android/arm,android/arm64,android/amd64"}
AAR_NAME=${AAR_NAME:-"lura.aar"}

if ! command -v gomobile >/dev/null 2>&1; then
  echo "[build] gomobile not found; will ensure deps"
fi

echo "[build] Ensuring gomobile deps"
bash "$(cd "$(dirname "$0")" && pwd)/ensure_gomobile_deps.sh"

echo "[build] Building AAR from ./termbridge"
GO111MODULE=on \
GOFLAGS=-mod=mod \
GOWORK=off \
GOPROXY="https://proxy.golang.org,direct" \
gomobile bind -target=${TARGETS} -androidapi=${ANDROID_API} -o "/tmp/${AAR_NAME}" ./termbridge

echo "[build] Ensuring Flutter app exists and libs folder present"
[[ -d "${APP_DIR}" ]] || { echo "[build] Flutter app '${APP_DIR}' not found. Run scripts/setup_flutter_app.sh first." >&2; exit 1; }
mkdir -p "${APP_DIR}/android/app/libs"
cp -v "/tmp/${AAR_NAME}" "${APP_DIR}/android/app/libs/${AAR_NAME}"
ls -l "${APP_DIR}/android/app/libs" || true

echo "[build] Verifying AAR package name and classes"
unzip -l "${APP_DIR}/android/app/libs/${AAR_NAME}" | sed -n '1,80p'

echo "[build] Adjusting Kotlin imports to match AAR package"
MAIN_ACTIVITY_FILE=$(find "${APP_DIR}/android/app/src/main/kotlin" -name MainActivity.kt | head -n 1)
if unzip -l "${APP_DIR}/android/app/libs/${AAR_NAME}" | grep -q " termbridge/Mobile.class"; then
  sed -i -E 's/^import (go\.)?termbridge\./import termbridge\./' "$MAIN_ACTIVITY_FILE"
elif unzip -l "${APP_DIR}/android/app/libs/${AAR_NAME}" | grep -q " go/termbridge/Mobile.class"; then
  sed -i -E 's/^import (go\.)?termbridge\./import go.termbridge\./' "$MAIN_ACTIVITY_FILE"
fi

echo "[build] Building Flutter APK (release)"
pushd "${APP_DIR}" >/dev/null
flutter build apk --release
popd >/dev/null

echo "[build] Done. APK under ${APP_DIR}/build/app/outputs/flutter-apk/"

