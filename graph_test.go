package main

import (
	"fmt"
	"testing"
)

func TestDfs(t *testing.T) {
	g := NewDirectedGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)
	g.AddVertex(8)
	g.AddVertex(9)
	g.AddVertex(10)

	g.AddEdge(1, 9)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 2)
	g.AddEdge(3, 4)
	g.AddEdge(5, 6)
	g.AddEdge(5, 8)
	g.AddEdge(6, 7)
	g.AddEdge(9, 10)

	visitedOrder := []int{}
	cb := func(i int) {
		visitedOrder = append(visitedOrder, i)
	}
	DFS(g, g.Vertices[1], cb)

	// add assertions here
	fmt.Println(visitedOrder)
}

func TestBfs(t *testing.T) {
	g := NewDirectedGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)
	g.AddVertex(8)
	g.AddVertex(9)
	g.AddVertex(10)

	g.AddEdge(1, 9)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 2)
	g.AddEdge(3, 4)
	g.AddEdge(5, 6)
	g.AddEdge(5, 8)
	g.AddEdge(6, 7)
	g.AddEdge(9, 10)

	visitedOrder := []int{}
	cb := func(i int) {
		visitedOrder = append(visitedOrder, i)
	}
	BFS(g, g.Vertices[1], cb)

	// add assertions here
	fmt.Println(visitedOrder)
}
