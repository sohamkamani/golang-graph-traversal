# Graph Traversal Algorithms

This is the example repository for my [blog post](https://codetree.dev/golang-graph-traversal/) on graph traversal algorithms.

To run the tests, clone the repository and run:

```
go test -v ./...
```

There are no assertions in the tests, but you'll be able to see the output showing the order of traversal for BFS and DFS:

```
❯ go test -v ./src/posts/golang-graph-algorithms/example
=== RUN   TestDfs
[1 9 10 5 6 7 8 2]
--- PASS: TestDfs (0.00s)
=== RUN   TestBfs
[1 9 5 2 10 8 6 7]
--- PASS: TestBfs (0.00s)
PASS
```
