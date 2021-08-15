package prefix_tree


type node struct {
	unit byte
	Sons map[byte]*node
}

func newNode(c byte) *node {
	return &node{
		unit: c,
		Sons:make(map[byte]*node),
	}
}