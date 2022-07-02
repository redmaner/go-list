package nbds

type node[D any] struct {
	data D
	next *node[D]
	prev *node[D]
}

func newNode[D any]() *node[D] {
	return &node[D]{}
}

func (n *node[D]) setData(data D) *node[D] {
	n.data = data
	return n
}

func (n *node[D]) setNext(next *node[D]) *node[D] {
	n.next = next
	return n
}

func (n *node[D]) setPrev(prev *node[D]) *node[D] {
	n.prev = prev
	return n
}

func (n *node[D]) updateData(updateFunc func(data D) D) *node[D] {
	n.data = updateFunc(n.data)
	return n
}

func (n *node[D]) unlink() {
	if n.prev != nil {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}
}
