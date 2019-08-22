package graph

import (
	"testing"
)

var g StringGraph

func fillGraph() {
	nA := NodeStr{"A"}
	nB := NodeStr{"B"}
	nC := NodeStr{"C"}
	nD := NodeStr{"D"}
	nE := NodeStr{"E"}
	nF := NodeStr{"F"}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)
}

func TestAdd(t *testing.T) {
	fillGraph()
	g.String()
}