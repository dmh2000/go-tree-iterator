package chatgpt

import (
	"golang.org/x/exp/constraints"
	"sqirvy.xyz/go-tree-iterator/rbt"
)

// chatgpt fix : change main struct name to ChatGptRBT
// chatgpt fix : add NewChatGptRBT factory function
// chatgpt fix : add IsEmpty function

// Color constants
const (
	Red   = true
	Black = false
)

// Node represents a node in the red-black tree.
type Node[K constraints.Ordered, V any] struct {
	// chatgpt fix: key, val K (val is not used)
	key   K
	value V
	// chatgpt fix: left, right, parent *Node[K, V] (parent is not used)
	left, right *Node[K, V]
	color       bool
	size        int
}

// ChatGptRBT represents a red-black binary search tree.
type ChatGptRBT[K constraints.Ordered, V any] struct {
	root *Node[K, V]
}

func NewRBT[K constraints.Ordered, V any]() *ChatGptRBT[K, V] {
	return &ChatGptRBT[K, V]{}
}

// chatgpt fix: missing IsEmpty function, copied from GeminiRBT
func (bst *ChatGptRBT[K, V]) IsEmpty() bool {
	return bst.root == nil
}

// NewNode creates a new red node with given key, value, size, and left and right children.
func NewNode[K constraints.Ordered, V any](key K, val V, size int, color bool) *Node[K, V] {
	return &Node[K, V]{
		key:   key,
		value: val,
		color: color,
		size:  size,
	}
}

// IsRed checks if a node is red.
func IsRed[K constraints.Ordered, V any](x *Node[K, V]) bool {
	if x == nil {
		return false
	}
	// fix: return x.color == Red
	return x.color
}

// Size returns the number of nodes in the tree rooted at x.
func Size[K constraints.Ordered, V any](x *Node[K, V]) int {
	if x == nil {
		return 0
	}
	return x.size
}

// RotateLeft performs a left rotation.
func RotateLeft[K constraints.Ordered, V any](h *Node[K, V]) *Node[K, V] {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = Red
	x.size = h.size
	h.size = 1 + Size(h.left) + Size(h.right)
	return x
}

// RotateRight performs a right rotation.
func RotateRight[K constraints.Ordered, V any](h *Node[K, V]) *Node[K, V] {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = Red
	x.size = h.size
	h.size = 1 + Size(h.left) + Size(h.right)
	return x
}

// FlipColors flips the colors of a node and its two children.
func FlipColors[K constraints.Ordered, V any](h *Node[K, V]) {
	h.color = Red
	h.left.color = Black
	h.right.color = Black
}

// Put inserts the specified key-value pair into the tree, overwriting the old value with the new value if the tree already contains the specified key.
func (t *ChatGptRBT[K, V]) Put(key K, val V) {
	t.root = t.put(t.root, key, val)
	t.root.color = Black
}

func (t *ChatGptRBT[K, V]) put(h *Node[K, V], key K, val V) *Node[K, V] {
	if h == nil {
		return NewNode(key, val, 1, Red)
	}

	if key < h.key {
		h.left = t.put(h.left, key, val)
	} else if key > h.key {
		h.right = t.put(h.right, key, val)
	} else {
		h.value = val
	}

	if IsRed(h.right) && !IsRed(h.left) {
		h = RotateLeft(h)
	}
	if IsRed(h.left) && IsRed(h.left.left) {
		h = RotateRight(h)
	}
	if IsRed(h.left) && IsRed(h.right) {
		FlipColors(h)
	}

	h.size = 1 + Size(h.left) + Size(h.right)
	return h
}

// Get returns the value associated with the given key.
func (t *ChatGptRBT[K, V]) Get(key K) (V, bool) {
	x := t.root
	for x != nil {
		if key < x.key {
			x = x.left
		} else if key > x.key {
			x = x.right
		} else {
			return x.value, true
		}
	}
	var zero V
	return zero, false
}

// Contains checks if the tree contains the given key.
func (t *ChatGptRBT[K, V]) Contains(key K) bool {
	_, found := t.Get(key)
	return found
}

// DeleteMin deletes the minimum key and associated value from the tree.
func (t *ChatGptRBT[K, V]) DeleteMin() {
	if t.root == nil {
		return
	}

	if !IsRed(t.root.left) && !IsRed(t.root.right) {
		t.root.color = Red
	}

	t.root = t.deleteMin(t.root)
	if t.root != nil {
		t.root.color = Black
	}
}

func (t *ChatGptRBT[K, V]) deleteMin(h *Node[K, V]) *Node[K, V] {
	if h.left == nil {
		return nil
	}

	if !IsRed(h.left) && !IsRed(h.left.left) {
		h = MoveRedLeft(h)
	}

	h.left = t.deleteMin(h.left)
	return Balance(h)
}

// MoveRedLeft makes a left-leaning red node into a right-leaning one.
func MoveRedLeft[K constraints.Ordered, V any](h *Node[K, V]) *Node[K, V] {
	FlipColors(h)
	if IsRed(h.right.left) {
		h.right = RotateRight(h.right)
		h = RotateLeft(h)
		FlipColors(h)
	}
	return h
}

// Balance restores red-black tree properties after a deletion.
func Balance[K constraints.Ordered, V any](h *Node[K, V]) *Node[K, V] {
	if IsRed(h.right) {
		h = RotateLeft(h)
	}
	if IsRed(h.left) && IsRed(h.left.left) {
		h = RotateRight(h)
	}
	if IsRed(h.left) && IsRed(h.right) {
		FlipColors(h)
	}

	h.size = 1 + Size(h.left) + Size(h.right)
	return h
}

// DeleteMax deletes the maximum key and associated value from the tree.
func (t *ChatGptRBT[K, V]) DeleteMax() {
	if t.root == nil {
		return
	}

	if !IsRed(t.root.left) && !IsRed(t.root.right) {
		t.root.color = Red
	}

	t.root = t.deleteMax(t.root)
	if t.root != nil {
		t.root.color = Black
	}
}

func (t *ChatGptRBT[K, V]) deleteMax(h *Node[K, V]) *Node[K, V] {
	if IsRed(h.left) {
		h = RotateRight(h)
	}

	if h.right == nil {
		return nil
	}

	if !IsRed(h.right) && !IsRed(h.right.left) {
		h = MoveRedRight(h)
	}

	h.right = t.deleteMax(h.right)
	return Balance(h)
}

// MoveRedRight makes a right-leaning red node into a left-leaning one.
func MoveRedRight[K constraints.Ordered, V any](h *Node[K, V]) *Node[K, V] {
	FlipColors(h)
	if IsRed(h.left.left) {
		h = RotateRight(h)
		FlipColors(h)
	}
	return h
}

// Delete deletes the specified key and its associated value from the tree.
func (t *ChatGptRBT[K, V]) Delete(key K) {
	if !t.Contains(key) {
		return
	}

	if !IsRed(t.root.left) && !IsRed(t.root.right) {
		t.root.color = Red
	}

	t.root = t.delete(t.root, key)
	if t.root != nil {
		t.root.color = Black
	}
}

func (t *ChatGptRBT[K, V]) delete(h *Node[K, V], key K) *Node[K, V] {
	if key < h.key {
		if !IsRed(h.left) && !IsRed(h.left.left) {
			h = MoveRedLeft(h)
		}
		h.left = t.delete(h.left, key)
	} else {
		if IsRed(h.left) {
			h = RotateRight(h)
		}
		if key == h.key && h.right == nil {
			return nil
		}
		if !IsRed(h.right) && !IsRed(h.right.left) {
			h = MoveRedRight(h)
		}
		if key == h.key {
			x := Min(h.right)
			h.key = x.key
			h.value = x.value
			h.right = t.deleteMin(h.right)
		} else {
			h.right = t.delete(h.right, key)
		}
	}
	return Balance(h)
}

// Min returns the node with the minimum key.
func Min[K constraints.Ordered, V any](h *Node[K, V]) *Node[K, V] {
	for h.left != nil {
		h = h.left
	}
	return h
}

// chatgpt fix: add GetAll function
func (bst *ChatGptRBT[K, V]) GetAll() []rbt.KeyValuePair[K, V] {
	pairs := make([]rbt.KeyValuePair[K, V], 0)
	var inorder func(*Node[K, V])
	inorder = func(n *Node[K, V]) {
		if n == nil {
			return
		}
		inorder(n.left)
		pairs = append(pairs, rbt.KeyValuePair[K, V]{Key: n.key, Val: n.value})
		inorder(n.right)
	}
	inorder(bst.root)
	return pairs
}

// chatgpt fix: add range over function Iterator
func (t *ChatGptRBT[K, V]) Iterator() func(func(rbt.KeyValuePair[K, V]) bool) {
	return func(yield func(rbt.KeyValuePair[K, V]) bool) {
		var inorder func(*Node[K, V])
		inorder = func(n *Node[K, V]) {
			if n == nil {
				return
			}
			inorder(n.left)
			if !yield(rbt.KeyValuePair[K, V]{Key: n.key, Val: n.value}) {
				return
			}
			inorder(n.right)
		}
		inorder(t.root)
	}
}

// chatgpt fix : remove main function
// func main() {
// 	rb := &ChatGptRBT[int, string]{}
// 	rb.Put(1, "one")
// 	rb.Put(2, "two")
// 	rb.Put(3, "three")

// 	val, found := rb.Get(2)
// 	if found {
// 		fmt.Println("Found:", val)
// 	} else {
// 		fmt.Println("Not found")
// 	}

// 	rb.Delete(2)
// 	_, found = rb.Get(2)
// 	if found {
// 		fmt.Println("Found")
// 	} else {
// 		fmt.Println("Not found")
// 	}
// }
