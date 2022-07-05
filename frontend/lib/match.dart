import 'package:flutter/material.dart';
import 'picture.dart';

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
