import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  final int cId = 3;
  final int oId = 6;

  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        body: Center(
          child: Row(
            children: [
              Picture(id: cId),
              Picture(id: oId),
            ],
          ),
        ),
      ),
    );
  }
}

class Picture extends StatelessWidget {
  final String api = "http://localhost:8080";
  final int id;

  const Picture({
    required this.id,
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: Image.network(
        "$api/p?id=$id",
        loadingBuilder: ((context, child, loadingProgress) {
          if (loadingProgress != null) {
            return const CircularProgressIndicator();
          }
          return child;
        }),
        fit: BoxFit.fitWidth,
      ),
    );
  }
}
