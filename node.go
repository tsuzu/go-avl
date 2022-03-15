package avl

// Node : node for AVL tree
type Node[Key any, Value any] struct {
	Key Key
	Val Value

	size, height int
	child        [2]*Node[Key, Value]
}

func nodeSizeOr[Key any, Value any](n *Node[Key, Value]) int {
	if n != nil {
		return n.size
	}

	return 0
}

func nodeHeightOr[Key any, Value any](n *Node[Key, Value]) int {
	if n != nil {
		return n.height
	}

	return 0
}

/*
   to rotate AVL Tree
*/

func rotate[Key any, Value any](n *Node[Key, Value], l, r int) *Node[Key, Value] {
	s := n.child[r]

	n.child[r] = s.child[l]
	s.child[l] = balance(n)

	if n != nil {
		n.size = nodeSizeOr(n.child[0]) + nodeSizeOr(n.child[1]) + 1
	}

	if s != nil {
		s.size = nodeSizeOr(s.child[0]) + nodeSizeOr(s.child[1]) + 1
	}

	return balance(s)
}

func balance[Key any, Value any](n *Node[Key, Value]) *Node[Key, Value] {
	for i := 0; i < 2; i++ {
		if nodeHeightOr(n.child[1-i])-nodeHeightOr(n.child[i]) < -1 {
			if nodeHeightOr(n.child[i].child[1-i])-nodeHeightOr(n.child[i].child[i]) > 0 {
				n.child[i] = rotate(n.child[i], i, 1-i)
			}

			return rotate(n, 1-i, i)
		}
	}

	if n != nil {
		n.height = maxInt(nodeHeightOr(n.child[0]), nodeHeightOr(n.child[1])) + 1
		n.size = nodeSizeOr(n.child[0]) + nodeSizeOr(n.child[1]) + 1
	}

	return n
}

func moveDown[Key any, Value any](n, m *Node[Key, Value]) *Node[Key, Value] {
	if n == nil {
		return m
	}

	n.child[1] = moveDown(n.child[1], m)
	return balance(n)
}

func newNode[Key any, Value any](k Key, v Value) *Node[Key, Value] {
	return &Node[Key, Value]{
		Key:    k,
		Val:    v,
		size:   1,
		height: 1,
	}
}
