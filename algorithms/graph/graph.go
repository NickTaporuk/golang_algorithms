// Package graph creates a ItemGraph data structure for the Item type
package graph

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Item=string Name=Str"
// Item the type of the binary search tree
type Item generic.Type
// NOTE: this is how easy it is to define a generic type
type Name generic.Type

// Node a single node that composes the tree
type NodeName struct {
	value Item
}

func (n *NodeName) String() string {
	return fmt.Sprintf("%v", n.value)
}

// ItemGraph the Items graph
type ItemGraph struct {
	nodes []*NodeName
	edges map[NodeName][]*NodeName
	lock  sync.RWMutex
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *NodeName) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) AddEdge(n1, n2 *NodeName) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[NodeName][]*NodeName)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
	g.lock.Unlock()
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) String() {
	g.lock.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.lock.RUnlock()
}
