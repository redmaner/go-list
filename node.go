package list

type node struct {
	data interface{}
	next *node
}

func NewNode() *node {
	return &node{}
}

func (n *node) setData(data interface{}) *node {
	n.data = data
	return n
}

func (n *node) setNext(next *node) *node {
	n.next = next
	return n
}
