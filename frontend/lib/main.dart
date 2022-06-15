import 'package:flutter/material.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    const cId = 3;
    const oId = 6;

    final httpLink = HttpLink(
      "http://localhost:8080/graphql",
    );

    ValueNotifier<GraphQLClient> client = ValueNotifier(
      GraphQLClient(
        link: httpLink,
        cache: GraphQLCache(
          store: InMemoryStore(),
        ),
      ),
    );

    return GraphQLProvider(
        client: client,
        child: MaterialApp(
          title: 'GraphQL Demo',
          theme: ThemeData(
            primarySwatch: Colors.blue,
          ),
          home: const Match(
            cId: cId,
            oId: oId,
          ),
        ));
  }
}

class Match extends StatelessWidget {
  const Match({
    Key? key,
    required this.cId,
    required this.oId,
  }) : super(key: key);

  final int cId;
  final int oId;

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Row(
        children: [
          Picture(id: cId),
          Picture(id: oId),
        ],
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
