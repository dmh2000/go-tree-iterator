package chatgpt

import (
	"golang.org/x/exp/constraints"

	rbt "sqirvy.xyz/go-tree-iterator/rbt"
)

const (
	RED   = true
	BLACK = false
)

type Node[K constraints.Ordered, V any] struct {
	key   K
	val   V
	left  *Node[K, V]
	right *Node[K, V]
	color bool
	size  int
}

type ChatGptRBT[K constraints.Ordered, V any] struct {
	root *Node[K, V]
}

func NewRBT[K constraints.Ordered, V any]() *ChatGptRBT[K, V] {
	return &ChatGptRBT[K, V]{}
}

func isRed[K constraints.Ordered, V any](x *Node[K, V]) bool {
	if x == nil {
		return false
	}
	return x.color
}

func (t *ChatGptRBT[K, V]) Size() int {
	return size(t.root)
}

func size[K constraints.Ordered, V any](x *Node[K, V]) int {
	if x == nil {
		return 0
	}
	return x.size
}

func (t *ChatGptRBT[K, V]) IsEmpty() bool {
	return t.root == nil
}

func (t *ChatGptRBT[K, V]) Get(key K) (V, bool) {
	return get(t.root, key)
}

func get[K constraints.Ordered, V any](x *Node[K, V], key K) (V, bool) {
	if x == nil {
		var zero V
		return zero, false
	}
	switch {
	case key == x.key:
		return x.val, true
	case key < x.key:
		return get(x.left, key)
	default:
		return get(x.right, key)
	}
}

func (t *ChatGptRBT[K, V]) Put(key K, val V) {
	t.root = put(t.root, key, val)
	t.root.color = BLACK
}

func put[K constraints.Ordered, V any](h *Node[K, V], key K, val V) *Node[K, V] {
	if h == nil {
		return &Node[K, V]{key: key, val: val, color: RED, size: 1}
	}
	switch {
	case key < h.key:
		h.left = put(h.left, key, val)
	case key > h.key:
		h.right = put(h.right, key, val)
	default:
		h.val = val
	}

	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}

	h.size = size(h.left) + size(h.right) + 1
	return h
}

func rotateRight[K constraints.Ordered, V any](h *Node[K, V]) *Node[K, V] {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
	x.size = h.size
	h.size = size(h.left) + size(h.right) + 1
	return x
}

func rotateLeft[K constraints.Ordered, V any](h *Node[K, V]) *Node[K, V] {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	x.size = h.size
	h.size = size(h.left) + size(h.right) + 1
	return x
}

func flipColors[K constraints.Ordered, V any](h *Node[K, V]) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

func (t *ChatGptRBT[K, V]) Min() (K, bool) {
	if t.IsEmpty() {
		var zero K
		return zero, false
	}
	x := min(t.root)
	return x.key, true
}

func min[K constraints.Ordered, V any](x *Node[K, V]) *Node[K, V] {
	if x.left == nil {
		return x
	}
	return min(x.left)
}

func (t *ChatGptRBT[K, V]) Max() (K, bool) {
	if t.IsEmpty() {
		var zero K
		return zero, false
	}
	x := max(t.root)
	return x.key, true
}

func max[K constraints.Ordered, V any](x *Node[K, V]) *Node[K, V] {
	if x.right == nil {
		return x
	}
	return max(x.right)
}

func (t *ChatGptRBT[K, V]) DeleteMin() {
	if t.IsEmpty() {
		panic("BST underflow")
	}
	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = RED
	}
	t.root = deleteMin(t.root)
	if !t.IsEmpty() {
		t.root.color = BLACK
	}
}

func deleteMin[K constraints.Ordered, V any](h *Node[K, V]) *Node[K, V] {
	if h.left == nil {
		return nil
	}
	if !isRed(h.left) && !isRed(h.left.left) {
		h = moveRedLeft(h)
	}
	h.left = deleteMin(h.left)
	return balance(h)
}

func (t *ChatGptRBT[K, V]) DeleteMax() {
	if t.IsEmpty() {
		panic("BST underflow")
	}
	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = RED
	}
	t.root = deleteMax(t.root)
	if !t.IsEmpty() {
		t.root.color = BLACK
	}
}

func deleteMax[K constraints.Ordered, V any](h *Node[K, V]) *Node[K, V] {
	if isRed(h.left) {
		h = rotateRight(h)
	}
	if h.right == nil {
		return nil
	}
	if !isRed(h.right) && !isRed(h.right.left) {
		h = moveRedRight(h)
	}
	h.right = deleteMax(h.right)
	return balance(h)
}

func (t *ChatGptRBT[K, V]) Delete(key K) {
	if t.IsEmpty() {
		return
	}
	if !isRed(t.root.left) && !isRed(t.root.right) {
		t.root.color = RED
	}
	t.root = delete(t.root, key)
	if !t.IsEmpty() {
		t.root.color = BLACK
	}
}

func delete[K constraints.Ordered, V any](h *Node[K, V], key K) *Node[K, V] {
	if key < h.key {
		if !isRed(h.left) && !isRed(h.left.left) {
			h = moveRedLeft(h)
		}
		h.left = delete(h.left, key)
	} else {
		if isRed(h.left) {
			h = rotateRight(h)
		}
		if key == h.key && h.right == nil {
			return nil
		}
		if !isRed(h.right) && !isRed(h.right.left) {
			h = moveRedRight(h)
		}
		if key == h.key {
			x := min(h.right)
			h.key = x.key
			h.val = x.val
			h.right = deleteMin(h.right)
		} else {
			h.right = delete(h.right, key)
		}
	}
	return balance(h)
}

// keys returns an iterator over the keys in the BST in ascending order.
func (bst *ChatGptRBT[K, V]) Keys() []K {
	queue := make([]K, 0)
	bst.keys(bst.root, &queue)
	return queue
}

func (bst *ChatGptRBT[K, V]) keys(x *Node[K, V], queue *[]K) {
	if x == nil {
		return
	}
	bst.keys(x.left, queue)
	*queue = append(*queue, x.key)
	bst.keys(x.right, queue)
}

// keys returns an iterator over the keys in the BST in ascending order.
func (bst *ChatGptRBT[K, V]) KeysInOrder() []K {
	queue := make([]K, 0)
	bst.keysInOrder(bst.root, &queue)
	return queue
}

func (bst *ChatGptRBT[K, V]) keysInOrder(x *Node[K, V], queue *[]K) {
	if x == nil {
		return
	}
	bst.keysInOrder(x.left, queue)
	*queue = append(*queue, x.key)
	bst.keysInOrder(x.right, queue)
}

// keys returns an iterator over the keys in the BST in level order.
func (bst *ChatGptRBT[K, V]) KeysLevelOrder() []K {
	queue := make([]K, 0)
	bst.keysLevelOrder(bst.root, &queue)
	return queue
}

func (bst *ChatGptRBT[K, V]) keysLevelOrder(x *Node[K, V], queue *[]K) {
	if x == nil {
		return
	}
	q := []*Node[K, V]{x}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		*queue = append(*queue, n.key)
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}
	}
}

// keys returns an iterator over the keys in the BST in pre order.
func (bst *ChatGptRBT[K, V]) KeysPreOrder() []K {
	queue := make([]K, 0)
	bst.keysPreOrder(bst.root, &queue)
	return queue
}

func (bst *ChatGptRBT[K, V]) keysPreOrder(x *Node[K, V], queue *[]K) {
	if x == nil {
		return
	}
	*queue = append(*queue, x.key)
	bst.keysPreOrder(x.left, queue)
	bst.keysPreOrder(x.right, queue)
}

// keys returns an iterator over the keys in the BST in post order.
func (bst *ChatGptRBT[K, V]) KeysPostOrder() []K {
	queue := make([]K, 0)
	bst.keysPostOrder(bst.root, &queue)
	return queue
}

func (bst *ChatGptRBT[K, V]) keysPostOrder(x *Node[K, V], queue *[]K) {
	if x == nil {
		return
	}
	bst.keysPostOrder(x.left, queue)
	bst.keysPostOrder(x.right, queue)
	*queue = append(*queue, x.key)
}

func moveRedLeft[K constraints.Ordered, V any](h *Node[K, V]) *Node[K, V] {
	flipColors(h)
	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
		flipColors(h)
	}
	return h
}

func moveRedRight[K constraints.Ordered, V any](h *Node[K, V]) *Node[K, V] {
	flipColors(h)
	if isRed(h.left.left) {
		h = rotateRight(h)
		flipColors(h)
	}
	return h
}

func balance[K constraints.Ordered, V any](h *Node[K, V]) *Node[K, V] {
	if isRed(h.right) && !isRed(h.left) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		flipColors(h)
	}
	h.size = size(h.left) + size(h.right) + 1
	return h
}

func (bst *ChatGptRBT[K, V]) GetAll() []rbt.KeyValuePair[K, V] {
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

func (t *ChatGptRBT[K, V]) Iterator() func(func(rbt.KeyValuePair[K, V]) bool) {
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
