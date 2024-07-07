package copilot

import (
	"fmt"

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

// CopilotRBT is a red-black tree
type CopilotRBT[K constraints.Ordered, V any] struct {
	root *Node[K, V]
}

// create a new red-black tree
func NewRBT[K constraints.Ordered, V any]() *CopilotRBT[K, V] {
	return &CopilotRBT[K, V]{}
}

// get the size of the tree from the root
func (t *CopilotRBT[K, V]) Size() int {
	return t.root.Size()
}

// check if the tree is empty
func (t *CopilotRBT[K, V]) IsEmpty() bool {
	return t.Size() == 0
}

// get the value of a key
func (t *CopilotRBT[K, V]) Get(key K) (V, error) {
	return t.get(t.root, key)
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

// get the value of a key from a specified subtree
func (t *CopilotRBT[K, V]) get(x *Node[K, V], key K) (V, error) {
	if x == nil {
		return t.root.val, fmt.Errorf("input node is nil")
	}

	for x != nil {
		cmp := compare(key, x.key)
		if cmp < 0 {
			x = x.left
		} else if cmp > 0 {
			x = x.right
		} else {
			return x.val, nil
		}
	}
	return t.root.val, fmt.Errorf("key not found")

}

// does this tree contain the given key?
func (t *CopilotRBT[K, V]) Contains(key K) (bool, error) {
	_, err := t.Get(key)
	if err != nil {
		return false, err
	}
	return true, nil
}

// insert a key-value pair into the red-black tree
func (t *CopilotRBT[K, V]) Put(key K, val V) {
	t.root = t.put(t.root, key, val)
	t.root.color = black
}

// insert the key-value pair in the subtree rooted at h
func (t *CopilotRBT[K, V]) put(h *Node[K, V], key K, val V) *Node[K, V] {
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
func (t *CopilotRBT[K, V]) rotateRight(h *Node[K, V]) *Node[K, V] {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = x.right.color
	x.right.color = red
	x.size = h.size
	h.size = h.left.Size() + h.right.Size() + 1
	return x
}

func (t *CopilotRBT[K, V]) rotateLeft(h *Node[K, V]) *Node[K, V] {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = x.left.color
	x.left.color = red
	x.size = h.size
	h.size = h.left.Size() + h.right.Size() + 1
	return x
}

func (t *CopilotRBT[K, V]) flipColors(h *Node[K, V]) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

// ************ Ordered Symbol Table Functions ***********
func (t *CopilotRBT[K, V]) Min() (K, error) {
	if t.IsEmpty() {
		return t.root.key, fmt.Errorf("tree is empty")
	}
	n, err := t.min(t.root)
	if err != nil {
		return t.root.key, err
	}
	return n.key, nil
}

func (t *CopilotRBT[K, V]) min(x *Node[K, V]) (*Node[K, V], error) {
	if x.left == nil {
		return x, nil
	}
	return t.min(x.left)
}

func (t *CopilotRBT[K, V]) Max() (K, error) {
	if t.IsEmpty() {
		return t.root.key, fmt.Errorf("tree is empty")
	}
	n, err := t.max(t.root)
	if err != nil {
		return t.root.key, err
	}
	return n.key, nil
}

func (t *CopilotRBT[K, V]) max(x *Node[K, V]) (*Node[K, V], error) {
	if x.right == nil {
		return x, nil
	}
	return t.max(x.right)
}

// return the largest key in the rbt tree less than or equal to key
func (t *CopilotRBT[K, V]) Floor(key K) (K, error) {
	if t.IsEmpty() {
		return t.root.key, fmt.Errorf("tree is empty")
	}
	x, err := t.floor(t.root, key)
	if err != nil {
		return t.root.key, err
	}
	return x.key, nil
}

// the largest key in the subtree rooted at x less than or equal to the given key
func (t *CopilotRBT[K, V]) floor(x *Node[K, V], key K) (*Node[K, V], error) {
	if x == nil {
		return x, fmt.Errorf("input node is nil")
	}

	cmp := compare(key, x.key)
	if cmp == 0 {
		return x, nil
	}
	if cmp < 0 {
		return t.floor(x.left, key)
	}

	n, err := t.floor(x.right, key)
	if err != nil {
		return nil, err
	}
	if n != nil {
		return n, nil
	}
	return x, err
}

// Returns the smallest key in the symbol table greater than or equal to key
func (t *CopilotRBT[K, V]) Ceiling(key K) (K, error) {
	if t.IsEmpty() {
		return t.root.key, fmt.Errorf("tree is empty")
	}
	x, err := t.ceiling(t.root, key)
	if err != nil {
		return t.root.key, err
	}
	return x.key, nil
}

// the smallest key in the subtree rooted at x greater than or equal to the given key
func (t *CopilotRBT[K, V]) ceiling(x *Node[K, V], key K) (*Node[K, V], error) {
	if x == nil {
		return x, fmt.Errorf("input node is nil")
	}

	cmp := compare(key, x.key)
	if cmp == 0 {
		return x, nil
	}
	if cmp > 0 {
		return t.ceiling(x.right, key)
	}

	n, err := t.ceiling(x.left, key)
	if err != nil {
		return nil, err
	}
	if n != nil {
		return n, nil
	}
	return x, err
}

// ************ RBT Range Functions ***********

// return all keys in ascending order
func (t *CopilotRBT[K, V]) Keys() ([]K, error) {
	if t.IsEmpty() {
		return nil, fmt.Errorf("tree is empty")
	}

	m, err := t.Min()
	if err != nil {
		return nil, err
	}

	n, err := t.Max()
	if err != nil {
		return nil, err
	}
	return t.KeysRange(m, n), nil
}

// return all keys in the range [lo..hi] in ascending order
func (t *CopilotRBT[K, V]) KeysRange(lo K, hi K) []K {
	var keys []K
	t.keys(t.root, &keys, lo, hi)
	return keys
}

// add the keys between lo and hi to the queue
func (t *CopilotRBT[K, V]) keys(x *Node[K, V], keys *[]K, lo K, hi K) {
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

func (bst *CopilotRBT[K, V]) GetAll() []rbt.KeyValuePair[K, V] {
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

func (t *CopilotRBT[K, V]) Iterator() func(yield func(K, V) bool) {
	return func(yield func(K, V) bool) {
		var inorder func(*Node[K, V])
		inorder = func(n *Node[K, V]) {
			if n == nil {
				return
			}
			inorder(n.left)
			if !yield(n.key, n.val) {
				return
			}
			inorder(n.right)
		}
		inorder(t.root)
	}
}
