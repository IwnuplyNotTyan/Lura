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

