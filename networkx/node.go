package networkx

import "mervin.me/algset/common"

//Node the struct of the node
type Node struct {
	ID         *common.ObjectID
	Label      string
	Attributes *common.Attributes
	Weight     float64
}

//NewNode create a new Node
func NewNode(name string) *Node {
	n := new(Node)
	n.Label = name
	//objectId
	return n
}

//SetLabel set the label of a node
func (n *Node) SetLabel(s string) {

}

//SetWeight set the weight of a node
func (n *Node) SetWeight(w float64) {

}

/*
Attributes
*/

//InitAttributes initialize the attributes of the node
func (n *Node) InitAttributes() {
	n.Attributes = common.NewAttributes()
}

//AddAttribute add one attribute
func (n *Node) AddAttribute(key string, v interface{}) {
	n.Attributes.Put(key, v)
}

//GetAttribute add one attribute
func (n *Node) GetAttribute(key string) interface{} {
	return n.Attributes.Get(key)
}
