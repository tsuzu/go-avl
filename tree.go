package avl

import (
	"golang.org/x/exp/constraints"
)

// Comparator is to compare keys
type Comparator[Key any] func(a, b Key) int

// Tree : AVL tree
type Tree[Key any, Value any] struct {
	root *Node[Key, Value]

	comparator Comparator[Key]
}

// NewTree creates a new AVL tree
func NewTree[Key any, Value any](comp Comparator[Key]) *Tree[Key, Value] {
	return &Tree[Key, Value]{nil, comp}
}

// NewTree creates a new AVL tree for comparable key types
func NewTreeOrdered[Key constraints.Ordered, Value any]() *Tree[Key, Value] {
	return &Tree[Key, Value]{nil, func(a, b Key) int {
		if a < b {
			return 1
		} else if a > b {
			return -1
		}

		return 0
	}}
}

func (t *Tree[Key, Value]) get(n *Node[Key, Value], key Key) (*Node[Key, Value], bool) {
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

func (t *Tree[Key, Value]) index(n *Node[Key, Value], rank int) (*Node[Key, Value], bool) {
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

func (t *Tree[Key, Value]) rank(n *Node[Key, Value], key Key, prev int) int {
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

func (t *Tree[Key, Value]) set(n, p *Node[Key, Value], force bool) (*Node[Key, Value], bool) {
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

func (t *Tree[Key, Value]) erase(n *Node[Key, Value], key Key) (*Node[Key, Value], bool) {
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
func (t *Tree[Key, Value]) Get(key Key) (*Node[Key, Value], bool) {
	return t.get(t.root, key)
}

// Index returns the node whose key is 'rank' th and true.
// If the key is not found, it returns nil and false.
func (t *Tree[Key, Value]) Index(rank int) (*Node[Key, Value], bool) {
	return t.index(t.root, rank)
}

// Rank returns the index of the node whose key is 'key'.
// If the key is not found, it returns -1.
func (t *Tree[Key, Value]) Rank(key Key) int {
	return t.rank(t.root, key, 0)
}

// Set adds a new node with 'key' and 'val' if a node with the key isn't found.
// Set updates the node with 'key' and 'val' if a node with the key is found.
func (t *Tree[Key, Value]) Set(key Key, val Value) {
	t.root, _ = t.set(t.root, newNode(key, val), true)
}

// Insert adds a new node with 'key' and 'val' if a node with the key isn't found.
// Insert do nothing if a node with the key is found.
func (t *Tree[Key, Value]) Insert(key Key, val Value) bool {
	var res bool
	t.root, res = t.set(t.root, newNode(key, val), false)

	return res
}

// Erase erases the node whose key is 'key'
// If the key is not found, it returns false.
func (t *Tree[Key, Value]) Erase(key Key) bool {
	var res bool
	t.root, res = t.erase(t.root, key)

	return res
}

// Size returns the size of the tree
func (t *Tree[Key, Value]) Size() int {
	return nodeSizeOr(t.root)
}
