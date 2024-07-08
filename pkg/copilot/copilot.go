package copilot

import (
	"golang.org/x/exp/constraints"

	rbt "sqirvy.xyz/go-tree-iterator/rbt"
)

const red = true
const black = false

// this implemention is derived directly from Sedgwick's Java implementation
// derived from https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/RedBlackBST.java.html

// Node represents a node in the red-black tree
type Node[K constraints.Ordered, V any] struct {
	key         K           // key
	val         V           // value
	left, right *Node[K, V] // links to left and right subtrees
	color       bool        // color of parent link
	size        int         // subtree count
}

func NewNode[K constraints.Ordered, V any](key K, val V, color bool, size int) *Node[K, V] {
	return &Node[K, V]{
		key:   key,
		val:   val,
		left:  nil,
		right: nil,
		color: color,
		size:  size,
	}
}

// get color of a node
func (n *Node[K, V]) IsRed() bool {
	if n == nil {
		return false
	}
	return n.color
}

// get size of a specified node
func (n *Node[K, V]) Size() int {
	if n == nil {
		return 0
	}
	return n.size
}

// CopilotRbt is a red-black tree
type CopilotRbt[K constraints.Ordered, V any] struct {
	root *Node[K, V]
}

// create a new red-black tree
func NewRBT[K constraints.Ordered, V any]() *CopilotRbt[K, V] {
	return &CopilotRbt[K, V]{}
}

// get the size of the tree from the root
func (t *CopilotRbt[K, V]) Size() int {
	return t.root.Size()
}

// check if the tree is empty
func (t *CopilotRbt[K, V]) IsEmpty() bool {
	return t.Size() == 0
}

// compare two Ordered values
func compare[T constraints.Ordered](a T, b T) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// get the value of a key
func (t *CopilotRbt[K, V]) Get(key K) (V, bool) {
	return t.get(t.root, key)
}

// get the value of a key from a specified subtree
func (t *CopilotRbt[K, V]) get(x *Node[K, V], key K) (V, bool) {
	if x == nil {
		return t.root.val, false
	}

	for x != nil {
		cmp := compare(key, x.key)
		if cmp < 0 {
			x = x.left
		} else if cmp > 0 {
			x = x.right
		} else {
			return x.val, true
		}
	}
	return t.root.val, false

}

// insert a key-value pair into the red-black tree
func (t *CopilotRbt[K, V]) Put(key K, val V) {
	t.root = t.put(t.root, key, val)
	t.root.color = black
}

// insert the key-value pair in the subtree rooted at h
func (t *CopilotRbt[K, V]) put(h *Node[K, V], key K, val V) *Node[K, V] {
	if h == nil {
		return NewNode(key, val, red, 1)
	}

	cmp := compare(key, h.key)
	if cmp < 0 {
		h.left = t.put(h.left, key, val)
	} else if cmp > 0 {
		h.right = t.put(h.right, key, val)
	} else {
		h.val = val
	}

	if h.right.IsRed() && !h.left.IsRed() {
		h = t.rotateLeft(h)
	}
	if h.left.IsRed() && h.left.left.IsRed() {
		h = t.rotateRight(h)
	}
	if h.left.IsRed() && h.right.IsRed() {
		t.flipColors(h)
	}

	h.size = h.left.Size() + h.right.Size() + 1
	return h
}

// ************ RBT helper functions ************

// Red-Black Rotations
func (t *CopilotRbt[K, V]) rotateRight(h *Node[K, V]) *Node[K, V] {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = x.right.color
	x.right.color = red
	x.size = h.size
	h.size = h.left.Size() + h.right.Size() + 1
	return x
}

func (t *CopilotRbt[K, V]) rotateLeft(h *Node[K, V]) *Node[K, V] {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = x.left.color
	x.left.color = red
	x.size = h.size
	h.size = h.left.Size() + h.right.Size() + 1
	return x
}

func (t *CopilotRbt[K, V]) flipColors(h *Node[K, V]) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

// ************ Ordered Symbol Table Functions ***********

// return all keys in the range [lo..hi] in ascending order
func (t *CopilotRbt[K, V]) Keys(lo K, hi K) []K {
	var keys []K
	t.keys(t.root, &keys, lo, hi)
	return keys
}

// add the keys between lo and hi to the queue
func (t *CopilotRbt[K, V]) keys(x *Node[K, V], keys *[]K, lo K, hi K) {
	if x == nil {
		return
	}
	cmplo := compare(lo, x.key)
	cmphi := compare(hi, x.key)
	if cmplo < 0 {
		t.keys(x.left, keys, lo, hi)
	}
	if cmplo <= 0 && cmphi >= 0 {
		*keys = append(*keys, x.key)
	}
	if cmphi > 0 {
		t.keys(x.right, keys, lo, hi)
	}
}

func (bst *CopilotRbt[K, V]) GetAll() []rbt.KeyValuePair[K, V] {
	pairs := make([]rbt.KeyValuePair[K, V], 0)
	var inorder func(*Node[K, V])
	inorder = func(n *Node[K, V]) {
		if n == nil {
			return
		}
		inorder(n.left)
		pairs = append(pairs, rbt.KeyValuePair[K, V]{Key: n.key, Val: n.val})
		inorder(n.right)
	}
	inorder(bst.root)
	return pairs
}

func (t *CopilotRbt[K, V]) Iterator() func(yield func(rbt.KeyValuePair[K, V]) bool) {
	return func(yield func(rbt.KeyValuePair[K, V]) bool) {
		var inorder func(*Node[K, V])
		inorder = func(n *Node[K, V]) {
			if n == nil {
				return
			}
			inorder(n.left)
			if !yield(rbt.KeyValuePair[K, V]{Key: n.key, Val: n.val}) {
				return
			}
			inorder(n.right)
		}
		inorder(t.root)
	}
}
