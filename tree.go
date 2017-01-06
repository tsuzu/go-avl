package avl

// Comparator is to compare keys
type Comparator func(a, b interface{}) int

// Tree : AVL tree
type Tree struct {
	root *Node

	comparator Comparator
}

// NewTree creates a new AVL tree
func NewTree(comp Comparator) *Tree {
	return &Tree{nil, comp}
}

func (t *Tree) get(n *Node, key interface{}) (*Node, bool) {
	if n == nil {
		return nil, false
	}

	g := t.comparator(n.Key, key)

	if g == 0 {
		return n, true
	} else if g < 0 {
		return t.get(n.child[0], key)
	} else {
		return t.get(n.child[1], key)
	}
}

func (t *Tree) index(n *Node, rank int) (*Node, bool) {
	if n == nil {
		return nil, false
	}

	m := nodeSizeOr(n.child[0])

	if rank < m {
		return t.index(n.child[0], rank)
	} else if rank == m {
		return n, true
	} else {
		return t.index(n.child[1], rank-m-1)
	}
}

func (t *Tree) rank(n *Node, key interface{}, prev int) int {
	if n == nil {
		return -1
	}

	g := t.comparator(n.Key, key)

	if g == 0 {
		return prev + nodeSizeOr(n.child[0])
	} else if g < 0 {
		return t.rank(n.child[0], key, prev)
	} else {
		return t.rank(n.child[1], key, prev+nodeSizeOr(n.child[0])+1)
	}
}

func (t *Tree) set(n, p *Node, force bool) (*Node, bool) {
	if n == nil {
		return p, true
	}

	g := t.comparator(n.Key, p.Key)

	isChanged := false
	if g == 0 {
		if force {
			n.Val = p.Val
		}
		return n, false
	} else if g < 0 {
		n.child[0], isChanged = t.set(n.child[0], p, force)
	} else {
		n.child[1], isChanged = t.set(n.child[1], p, force)
	}

	if isChanged {
		n.size++

		return balance(n), true
	}

	return n, false
}

func (t *Tree) erase(n *Node, key interface{}) (*Node, bool) {
	if n == nil {
		return nil, false
	}

	g := t.comparator(n.Key, key)

	isChanged := false
	if g == 0 {
		return moveDown(n.child[0], n.child[1]), true
	} else if g < 0 {
		n.child[0], isChanged = t.erase(n.child[0], key)
	} else {
		n.child[1], isChanged = t.erase(n.child[1], key)
	}

	if isChanged {
		n.size--

		return balance(n), true
	}

	return n, false
}

// Get returns the node whose key is 'key' and true.
// If the key is not found, it returns nil and false.
func (t *Tree) Get(key interface{}) (*Node, bool) {
	return t.get(t.root, key)
}

// Index returns the node whose key is 'rank' th and true.
// If the key is not found, it returns nil and false.
func (t *Tree) Index(rank int) (*Node, bool) {
	return t.index(t.root, rank)
}

// Rank returns the index of the node whose key is 'key'.
// If the key is not found, it returns -1.
func (t *Tree) Rank(key interface{}) int {
	return t.rank(t.root, key, 0)
}

// Set adds a new node with 'key' and 'val' if a node with the key isn't found.
// Set updates the node with 'key' and 'val' if a node with the key is found.
func (t *Tree) Set(key, val interface{}) {
	t.root, _ = t.set(t.root, newNode(key, val), true)
}

// Insert adds a new node with 'key' and 'val' if a node with the key isn't found.
// Insert do nothing if a node with the key is found.
func (t *Tree) Insert(key, val interface{}) bool {
	var res bool
	t.root, res = t.set(t.root, newNode(key, val), false)

	return res
}

// Erase erases the node whose key is 'key'
// If the key is not found, it returns false.
func (t *Tree) Erase(key interface{}) bool {
	var res bool
	t.root, res = t.erase(t.root, key)

	return res
}

// Size returns the size of the tree
func (t *Tree) Size() int {
	return nodeSizeOr(t.root)
}
