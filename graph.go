package main

import (
	"strconv"
)

type Vertex struct {
	// Key is the unique identifier of the vertex
	Key int
	// Vertices will describe vertices connected to this one
	// The key will be the Key value of the connected vertice
	// with the value being the pointer to it
	Vertices map[int]*Vertex
}

// We then create a constructor function for the Vertex
func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[int]*Vertex{},
	}
}

func (v *Vertex) String() string {
	s := strconv.Itoa(v.Key) + ":"

	for _, neighbor := range v.Vertices {
		s += " " + strconv.Itoa(neighbor.Key)
	}

	return s
}

type Graph struct {
	// Vertices describes all vertices contained in the graph
	// The key will be the Key value of the connected vertice
	// with the value being the pointer to it
	Vertices map[int]*Vertex
	// This will decide if it's a directed or undirected graph
	directed bool
}

// We defined constructor functions that create
// new directed or undirected graphs respectively

func NewDirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
		directed: true,
	}
}

func NewUndirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
	}
}

// AddVertex creates a new vertex with the given
// key and adds it to the graph
func (g *Graph) AddVertex(key int) {
	v := NewVertex(key)
	g.Vertices[key] = v
}

// The AddEdge method adds an edge between two vertices in the graph
func (g *Graph) AddEdge(k1, k2 int) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	// return an error if one of the vertices doesn't exist
	if v1 == nil || v2 == nil {
		panic("not all vertices exist")
	}

	// do nothing if the vertices are already connected
	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	// Add a directed edge between v1 and v2
	// If the graph is undirected, add a corresponding
	// edge back from v2 to v1, effectively making the
	// edge between v1 and v2 bidirectional
	v1.Vertices[v2.Key] = v2
	if !g.directed && v1.Key != v2.Key {
		v2.Vertices[v1.Key] = v1
	}

	// Add the vertices to the graph's vertex map
	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
}

func (g *Graph) String() string {
	s := ""
	i := 0
	for _, v := range g.Vertices {
		if i != 0 {
			s += "\n"
		}
		s += v.String()
		i++
	}
	return s
}
