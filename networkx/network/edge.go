package network

import (
	"mervin.me/algset/common"
)

type EdgeType uint8

const (
	Directed EdgeType = iota
	Undirected
)

//Edge
type Edge struct {
	Attributes *common.Attributes
	Source     *Node
	Target     *Node
}

//NewEdge create a new edge
func NewEdge(s, t *Node) *Edge {
	e := &Edge{}
	e.Source = s
	e.Target = t
	e.Attributes = common.NewAttributes()
	return e
}
