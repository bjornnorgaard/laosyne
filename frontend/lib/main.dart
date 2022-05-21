import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        body: Container(
          child: Center(
            child: Image.network(
              "http://10.0.2.2:8080/p?id=3",
              loadingBuilder: ((context, child, loadingProgress) {
                return loadingProgress == null
                    ? child
                    : const CircularProgressIndicator();
              }),
              height: 700,
              width: 700,
              fit: BoxFit.fitWidth,
            ),
          ),
        ),
      ),
    );
  }
}
