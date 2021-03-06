package graph

import "testing"

func TestNew(t *testing.T) {
	g := NewGraph(2)
	if g.nbVertices != 2 ||
		g.nbEdges != 0 ||
		len(g.adj) != 2 {
		t.Error("Expected {2 0 [[] []]} got", *g)
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph(2)
	g.AddEdge(0, 1)
	if g.nbVertices != 2 ||
		g.nbEdges != 1 ||
		g.adj[0][0] != 1 ||
		g.adj[1][0] != 0 {
		t.Error("Expected {2 1 [[1] [0]]} got", *g)
	}
}

func TestString(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)

	expectedResult := "3 vertices, 2 edges\n0: 1,2,\n1: 0,\n2: 0,\n"
	res := g.String()
	if res != expectedResult {
		t.Error("Expected", expectedResult, "got", res)
	}
}
