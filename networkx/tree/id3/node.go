package id3

import (
	"mervin.me/algset/common"
	"mervin.me/algset/networkx"
)

//Node the struct of the ID3 Node
type Node struct {
	networkx.Node

	parent   *Node
	children []*Edge
}

//NewNode create a new Node struct
func NewNode(s string) *Node {
	n := new(Node)
	n.Label = s
	n.Attributes = common.NewAttributes()
	n.children = make([]*Edge, 0)
	return n
}

//AddEdge add a new edge from the node n to the other
func (n *Node) AddEdge(other *Node) *Edge {
	e := &Edge{
		Source: n,
		Target: other,
	}
	e.Attributes = common.NewAttributes()
	n.children = append(n.children, e)
	return e
}

// Edges return all the out edge from current n
func (n *Node) Edges() []*Edge {
	return n.children
}
