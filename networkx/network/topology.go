package network

import (
	"mervin.me/algset/common"
)

//Network
type Network struct {
	topo map[*common.ObjectID]*Node

	netType EdgeType
}

//NewBlankNetwork
func NewBlankNetwork() *Network {
	n := &Network{
		topo:    make(map[*common.ObjectID]*Node, 0),
		netType: Undirected,
	}
	return n
}

// //AddNode
func (net *Network) AddNodeByID(id uint32, label string) *Node {
	oi := common.NewObjectID(id)
	if n, ok := net.topo[oi]; ok {
		return n
	}
	n := NewNode(label)
	n.ID = oi
	net.topo[oi] = n
	return n
}

//AddNode
func (net *Network) AddNode(n *Node) {
	if _, ok := net.topo[n.ID]; !ok {
		net.topo[n.ID] = n
	}
}

//AddEdge
func (net *Network) AddEdge(src, dst *Node) *Edge {
	if n, ok := net.topo[src.ID]; ok {
		if n.HasAdjNode(dst.ID) {
			return n.GetAdjEdge(dst.ID)
		} else {
			return n.AddAdjNode(dst)
		}
	} else {
		net.AddNode(src)
		return n.AddAdjNode(dst)
	}
	if net.netType == Undirected {
		net.AddEdge(dst, src)
	}
	return nil
}

/*
get
*/
//GetNodeByLocal
func (net *Network) GetNodeByLocal(id int) *Node {
	oi := common.NewObjectID(uint32(id))
	if n, ok := net.topo[oi]; ok {
		return n
	}
	return nil
}
