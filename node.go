package avl

// Node : node for AVL tree
type Node struct {
	Key interface{}
	Val interface{}

	size, height int
	child        [2]*Node
}

func nodeSizeOr(n *Node) int {
	if n != nil {
		return n.size
	}

	return 0
}

func nodeHeightOr(n *Node) int {
	if n != nil {
		return n.height
	}

	return 0
}

/*
   to rotate AVL Tree
*/

func rotate(n *Node, l, r int) *Node {
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

func balance(n *Node) *Node {
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

func moveDown(n, m *Node) *Node {
	if n == nil {
		return m
	}

	n.child[1] = moveDown(n.child[1], m)
	return balance(n)
}

func newNode(k, v interface{}) *Node {
	return &Node{
		Key:    k,
		Val:    v,
		size:   1,
		height: 1,
	}
}
