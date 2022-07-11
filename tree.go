package nbds

import "golang.org/x/exp/constraints"

func inverseTreeNode[D constraints.Ordered](tn *node[D]) {
	if tn == nil {
		return
	}

	tn.left, tn.right = tn.right, tn.left

	inverseTreeNode(tn.left)
	inverseTreeNode(tn.right)
}

type Tree[D constraints.Ordered] struct {
	root *node[D]
}

func (t *Tree[D]) Inverse() {
	inverseTreeNode(t.root)
}
