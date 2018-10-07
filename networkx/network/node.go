package network

import (
	"mervin.me/algset/common"
	"mervin.me/algset/networkx"
)

//Node
type Node struct {
	networkx.Node

	adj []*Edge
}

//NewNode create a new Node
func NewNode(name string) *Node {
	n := new(Node)
	n.Label = name
	//objectId
	return n
}

//GetAdjEdges
func (n *Node) AddAdjNode(dst *Node) *Edge {
	if n.HasAdjNode(dst.ID) {
		return n.GetAdjEdge(dst.ID)
	}
	e := NewEdge(n, dst)
	n.adj = append(n.adj, e)
	return e
}

//GetAdjNodes
func (n *Node) GetAdjNodes() []*Node {
	ns := make([]*Node, 0)
	for _, e := range n.adj {
		if e.Source.ID.Equal(n.ID) {
			ns = append(ns, e.Target)
		} else {
			ns = append(ns, e.Source)
		}
	}
	return ns
}

//GetAdjEdges
func (n *Node) GetAdjEdges() []*Edge {
	return n.adj
}

//GetAdjNodes
func (n *Node) GetAdjNode(id *common.ObjectID) *Node {
	if n.adj == nil {
		return nil
	}
	for _, e := range n.adj {
		if e.Source.ID.Equal(id) {
			return e.Source
		} else if e.Target.ID.Equal(id) {
			return e.Target
		}
	}
	return nil
}

//GetAdjNodes
func (n *Node) GetAdjEdge(id *common.ObjectID) *Edge {
	if n.adj == nil {
		return nil
	}
	for _, e := range n.adj {
		if e.Source.ID.Equal(id) || e.Target.ID.Equal(id) {
			return e
		}
	}
	return nil
}

//HasAdjNode
func (n *Node) HasAdjNode(id *common.ObjectID) bool {
	if n.adj == nil {
		return false
	}
	for _, e := range n.adj {
		if e.Source.ID.Equal(id) {
			return true
		} else if e.Target.ID.Equal(id) {
			return true
		}
	}
	return false
}
