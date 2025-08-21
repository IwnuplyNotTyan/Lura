#!/usr/bin/env bash
set -euo pipefail

# This script creates a Flutter app (if missing), wires channels, adds xterm, and
# prepares Android Gradle to consume the gomobile AAR built from ./termbridge.

APP_DIR=${APP_DIR:-"flutter_app"}
CHANNEL_METHOD=${CHANNEL_METHOD:-"game.method"}
CHANNEL_EVENTS=${CHANNEL_EVENTS:-"game.events"}

command -v flutter >/dev/null 2>&1 || {
  echo "[setup] flutter not found. Install Flutter SDK and ensure it is on PATH." >&2
  exit 1
}

echo "[setup] Ensuring Flutter app exists at ${APP_DIR}"
if [[ ! -d "${APP_DIR}" ]]; then
  flutter create --org com.example --project-name lura_flutter "${APP_DIR}"
fi

pushd "${APP_DIR}" >/dev/null

echo "[setup] Adding xterm dependency"
flutter pub add xterm >/dev/null

echo "[setup] Writing lib/main.dart"
cat > lib/main.dart <<'DART'
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:xterm/xterm.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Lura CLI',
      theme: ThemeData(colorSchemeSeed: Colors.teal, useMaterial3: true),
      home: const GameScreen(),
    );
  }
}

class GameScreen extends StatefulWidget {
  const GameScreen({super.key});

  @override
  State<GameScreen> createState() => _GameScreenState();
}

class _GameScreenState extends State<GameScreen> {
  final _terminal = Terminal(maxLines: 2000, platform: TerminalTargetPlatform.android);
  final _method = const MethodChannel('game.method');
  final _events = const EventChannel('game.events');
  final _controller = TextEditingController();

  @override
  void initState() {
    super.initState();
    _method.invokeMethod('start');
    _events.receiveBroadcastStream().listen((event) {
      if (event is String) {
        _terminal.write(event);
        _terminal.write('\r\n');
      }
    });
  }

  @override
  void dispose() {
    _method.invokeMethod('stop');
    _controller.dispose();
    super.dispose();
  }

  void _send(String text) {
    _method.invokeMethod('send', {'text': text});
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Lura CLI')),
      body: Column(
        children: [
          Expanded(child: TerminalView(_terminal)),
          Row(children: [
            Expanded(
              child: TextField(
                controller: _controller,
                onSubmitted: (v) { if (v.isNotEmpty) { _send(v); _controller.clear(); }},
                decoration: const InputDecoration(hintText: 'Введите команду…', contentPadding: EdgeInsets.symmetric(horizontal: 12)),
              ),
            ),
            IconButton(icon: const Icon(Icons.send), onPressed: () { final v = _controller.text; if (v.isNotEmpty) { _send(v); _controller.clear(); } }),
          ]),
          const SizedBox(height: 8),
          Wrap(spacing: 8, children: [
            ElevatedButton(onPressed: () => _send('look'), child: const Text('look')),
            ElevatedButton(onPressed: () => _send('attack'), child: const Text('attack')),
            ElevatedButton(onPressed: () => _send('exit'), child: const Text('exit')),
          ]),
          const SizedBox(height: 12),
        ],
      ),
    );
  }
}
DART

echo "[setup] Ensuring android/app/libs exists"
mkdir -p android/app/libs

APP_BUILD_GRADLE=android/app/build.gradle
echo "[setup] Patching ${APP_BUILD_GRADLE} to include AAR and packaging options"

if ! grep -q "packagingOptions" "$APP_BUILD_GRADLE"; then
  cat >> "$APP_BUILD_GRADLE" <<'GRADLE'

android {
    packagingOptions {
        pickFirst "lib/**/libgojni.so"
    }
}
GRADLE
fi

APP_BUILD_GRADLE_PROJ=android/build.gradle
if ! grep -q "flatDir" "$APP_BUILD_GRADLE_PROJ"; then
  cat >> "$APP_BUILD_GRADLE_PROJ" <<'GRADLE'

allprojects {
    repositories {
        flatDir { dirs 'app/libs' }
    }
}
GRADLE
fi

if ! grep -q "implementation(name: 'lura', ext: 'aar')" "$APP_BUILD_GRADLE"; then
  cat >> "$APP_BUILD_GRADLE" <<'GRADLE'

dependencies {
    implementation(name: 'lura', ext: 'aar')
}
GRADLE
fi

echo "[setup] Updating MainActivity to wire channels"
MAIN_ACTIVITY_FILE=$(find android/app/src/main/kotlin -name MainActivity.kt | head -n 1)
if [[ -z "${MAIN_ACTIVITY_FILE}" ]]; then
  echo "[setup] ERROR: MainActivity.kt not found" >&2
  exit 1
fi

PKG_LINE=$(grep -E '^package ' "${MAIN_ACTIVITY_FILE}" | head -n 1 || true)
PKG_NAME=${PKG_LINE#package }
if [[ -z "${PKG_NAME}" ]]; then
  PKG_NAME="com.example.lura_flutter"
fi

cat > "${MAIN_ACTIVITY_FILE}" <<KOTLIN
package ${PKG_NAME}

import io.flutter.embedding.android.FlutterActivity
import io.flutter.plugin.common.EventChannel
import io.flutter.plugin.common.MethodChannel
import go.termbridge.Mobile
import go.termbridge.Output

class MainActivity: FlutterActivity() {
    private val method = "${CHANNEL_METHOD}"
    private val events = "${CHANNEL_EVENTS}"

    private var mobile: Mobile? = null
    private var eventSink: EventChannel.EventSink? = null

    private inner class NativeOut: Output {
        override fun OnLine(line: String?) {
            line?.let { eventSink?.success(it) }
        }
    }

    override fun configureFlutterEngine(engine: io.flutter.embedding.engine.FlutterEngine) {
        super.configureFlutterEngine(engine)
        EventChannel(engine.dartExecutor.binaryMessenger, events).setStreamHandler(object: EventChannel.StreamHandler {
            override fun onListen(args: Any?, sink: EventChannel.EventSink?) { eventSink = sink }
            override fun onCancel(args: Any?) { eventSink = null }
        })
        MethodChannel(engine.dartExecutor.binaryMessenger, method).setMethodCallHandler { call, result ->
            when (call.method) {
                "start" -> {
                    if (mobile == null) mobile = Mobile(NativeOut())
                    mobile?.Start()
                    result.success(null)
                }
                "send" -> {
                    val text = call.argument<String>("text") ?: ""
                    mobile?.SendLine(text)
                    result.success(null)
                }
                "stop" -> {
                    mobile?.Close()
                    mobile = null
                    result.success(null)
                }
                else -> result.notImplemented()
            }
        }
    }
}
KOTLIN

popd >/dev/null

echo "[setup] Flutter app is ready: ${APP_DIR}"

