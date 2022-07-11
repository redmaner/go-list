package nbds

type node[D any] struct {
	data  D
	right *node[D]
	left  *node[D]
}

func newNode[D any]() *node[D] {
	return &node[D]{}
}

func (n *node[D]) setData(data D) *node[D] {
	n.data = data
	return n
}

func (n *node[D]) setRight(right *node[D]) *node[D] {
	n.right = right
	return n
}

func (n *node[D]) setLeft(left *node[D]) *node[D] {
	n.left = left
	return n
}

func (n *node[D]) updateData(updateFunc func(data D) D) *node[D] {
	n.data = updateFunc(n.data)
	return n
}

func (n *node[D]) unlink() {
	if n.left != nil {
		n.left.right = n.right
	}

	if n.right != nil {
		n.right.left = n.left
	}
}
