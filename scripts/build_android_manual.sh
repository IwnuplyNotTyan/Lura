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

command -v gomobile >/dev/null 2>&1 || {
  echo "[build] gomobile not found; installing..." >&2
  go install golang.org/x/mobile/cmd/gomobile@latest
  go install golang.org/x/mobile/cmd/gobind@latest
  gomobile init
}

echo "[build] Building AAR from ./termbridge"
gomobile bind -target=${TARGETS} -androidapi=${ANDROID_API} -o "/tmp/${AAR_NAME}" ./termbridge

echo "[build] Ensuring Flutter app exists and libs folder present"
[[ -d "${APP_DIR}" ]] || { echo "[build] Flutter app '${APP_DIR}' not found. Run scripts/setup_flutter_app.sh first." >&2; exit 1; }
mkdir -p "${APP_DIR}/android/app/libs"
cp -f "/tmp/${AAR_NAME}" "${APP_DIR}/android/app/libs/${AAR_NAME}"

echo "[build] Building Flutter APK (release)"
pushd "${APP_DIR}" >/dev/null
flutter build apk --release
popd >/dev/null

echo "[build] Done. APK under ${APP_DIR}/build/app/outputs/flutter-apk/"

