package prefix_tree


type node struct {
	unit byte
	Sons []*node
}

func newNode(c byte) *node {
	re:= node{unit: c}
	return &re
}

func (n *node)IsEqual(a *node)bool{
	return n.unit==a.unit
}
