package id3

import (
	"mervin.me/algset/common"
)

//Edge define the relationship between two nodes
type Edge struct {
	Attributes *common.Attributes
	Source     *Node
	Target     *Node
}

//NewEdge create a new edge
func NewEdge(s, t *Node) *Edge {
	e := &Edge{
		Source: s,
		Target: t,
	}
	e.Attributes = common.NewAttributes()
	return e
}
