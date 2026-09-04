import 'dart:convert';
import 'dart:io';

import 'package:web_socket_channel/io.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class AISStreamWebsocketClient {
  final String _serverUri;

  AISStreamWebsocketClient(this._serverUri);

  void connect() async {
    // Enables compression, which aisstream.io requires to serve full message bandwidth.
    final socket = await WebSocket.connect(
      _serverUri,
      compression: const CompressionOptions(
        clientNoContextTakeover: true,
        serverNoContextTakeover: true,
      ),
    );
    final WebSocketChannel channel = IOWebSocketChannel(socket);
    await channel.ready;
    channel.stream.listen(onMessage);
    channel.sink.add(
      '{"APIKey":"<YOUR API KEY>","BoundingBoxes":[[[-90,-180],[90,180]]]}',
    );
  }

  void onMessage(dynamic message) {
    final jsonString = utf8.decode(message);
    print(jsonString);
  }
}
