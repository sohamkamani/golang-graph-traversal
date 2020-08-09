package main

// create a node that holds the graphs vertex as data
type node struct {
	v    *Vertex
	next *node
}

// create a queue data structure
type queue struct {
	head *node
	tail *node
}

// enqueue adds a new node to the tail of the queue
func (q *queue) enqueue(v *Vertex) {
	n := &node{v: v}

	// if the queue is empty, set the head and tail as the node value
	if q.tail == nil {
		q.head = n
		q.tail = n
		return
	}

	q.tail.next = n
	q.tail = n
}

// dequeue removes the head from the queue and returns it
func (q *queue) dequeue() *Vertex {
	n := q.head
	// return nil, if head is empty
	if n == nil {
		return nil
	}

	q.head = q.head.next

	// if there wasn't any next node, that
	// means the queue is empty, and the tail
	// should be set to nil
	if q.head == nil {
		q.tail = nil
	}

	return n.v
}

func BFS(g *Graph, startVertex *Vertex, visitCb func(int)) {
	// initialize queue and visited vertices map
	vertexQueue := &queue{}
	visitedVertices := map[int]bool{}

	currentVertex := startVertex
	// start a continuous loop
	for {
		// visit the current node
		visitCb(currentVertex.Key)
		visitedVertices[currentVertex.Key] = true

		// for each neighboring vertex, push it to the queue
		// if it isn't already visited
		for _, v := range currentVertex.Vertices {
			if !visitedVertices[v.Key] {
				vertexQueue.enqueue(v)
			}
		}

		// change the current vertex to the next one
		// in the queue
		currentVertex = vertexQueue.dequeue()
		// if the queue is empty, break out of the loop
		if currentVertex == nil {
			break
		}
	}
}
