package list

type node[D comparable] struct {
	data D
	next *node[D]
}

func NewNode[D comparable]() *node[D] {
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
