package ${PACKAGE_NAME}

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

