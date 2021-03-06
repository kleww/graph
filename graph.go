package graph

import "fmt"

// Graph represents an undirected graph.
// The zero value in an empty graph ready to use
type Graph struct {
	nbVertices int
	nbEdges    int
	adj        [][]int
}

// NewGraph returns an initialized graph with n vertices
func NewGraph(n int) *Graph {
	g := &Graph{}
	g.nbVertices = n
	g.nbEdges = 0
	g.adj = make([][]int, n)
	return g
}

// AddEdge links vertices v and w in both direction
func (g *Graph) AddEdge(v int, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.nbEdges++
}

// String returns a string representation of the graph's adjacency lists
func (g Graph) String() string {
	s := fmt.Sprintln(g.nbVertices, "vertices,", g.nbEdges, "edges")
	for v := 0; v < g.nbVertices; v++ {
		s += fmt.Sprint(v, ": ")
		for _, w := range g.adj[v] {
			s += fmt.Sprint(w, ",")
		}
		s += fmt.Sprintln()
	}
	return s
}
