package gemini

import (
	"fmt"

	"golang.org/x/exp/constraints"

	rbt "sqirvy.xyz/go-tree-iterator/rbt"
)

// Node represents a node in the Red-Black BST.
type Node[K constraints.Ordered, V any] struct {
	key         K
	val         V
	color       bool // color of parent link
	left, right *Node[K, V]
	N           int // subtree count
}

// GeminiRbt represents a Red-Black BST.
type GeminiRbt[K constraints.Ordered, V any] struct {
	root *Node[K, V]
}

func NewRBT[K constraints.Ordered, V any]() *GeminiRbt[K, V] {
	return &GeminiRbt[K, V]{}
}

// isEmpty returns true if the BST is empty.
func (bst *GeminiRbt[K, V]) IsEmpty() bool {
	return bst.root == nil
}

// size returns the number of key-value pairs in the BST.
func (bst *GeminiRbt[K, V]) Size() int {
	return bst.size(bst.root)
}

func (bst *GeminiRbt[K, V]) size(x *Node[K, V]) int {
	if x == nil {
		return 0
	}
	return x.N
}

// get returns the value associated with the given key.
func (bst *GeminiRbt[K, V]) Get(key K) (V, bool) {
	x := bst.get(bst.root, key)
	if x == nil {
		var zero V
		return zero, false
	}
	return x.val, true
}

func (bst *GeminiRbt[K, V]) get(x *Node[K, V], key K) *Node[K, V] {
	for x != nil {
		switch {
		case key < x.key:
			x = x.left
		case key > x.key:
			x = x.right
		default:
			return x
		}
	}
	return nil
}

// put inserts the key-value pair into the BST.
func (bst *GeminiRbt[K, V]) Put(key K, val V) {
	bst.root = bst.put(bst.root, key, val)
}

func (bst *GeminiRbt[K, V]) put(h *Node[K, V], key K, val V) *Node[K, V] {
	if h == nil {
		return &Node[K, V]{key: key, val: val, color: false, N: 1}
	}
	switch {
	case key < h.key:
		h.left = bst.put(h.left, key, val)
	case key > h.key:
		h.right = bst.put(h.right, key, val)
	default:
		h.val = val
		return h
	}
	if !isRed(h.left) && isRed(h.right) {
		h = bst.rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = bst.rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		bst.flipColors(h)
	}
	h.N = 1 + bst.size(h.left) + bst.size(h.right)
	return h
}

// min returns the smallest key in the BST.
func (bst *GeminiRbt[K, V]) Min() (K, bool) {
	x := bst.min(bst.root)
	if x == nil {
		var zero K
		return zero, false
	}
	return x.key, true
}

func (bst *GeminiRbt[K, V]) min(x *Node[K, V]) *Node[K, V] {
	if x == nil {
		return nil
	}
	for x.left != nil {
		x = x.left
	}
	return x
}

// max returns the largest key in the BST.
func (bst *GeminiRbt[K, V]) Max() (K, bool) {
	x := bst.max(bst.root)
	if x == nil {
		var zero K
		return zero, false
	}
	return x.key, true
}

func (bst *GeminiRbt[K, V]) max(x *Node[K, V]) *Node[K, V] {
	if x == nil {
		return nil
	}
	for x.right != nil {
		x = x.right
	}
	return x
}

// floor returns the largest key less than or equal to the given key.
func (bst *GeminiRbt[K, V]) Floor(key K) (K, bool) {
	x := bst.floor(bst.root, key)
	if x == nil {
		var zero K
		return zero, false
	}
	return x.key, true
}

func (bst *GeminiRbt[K, V]) floor(x *Node[K, V], key K) *Node[K, V] {
	if x == nil {
		return nil
	}
	switch {
	case key < x.key:
		return bst.floor(x.left, key)
	case key > x.key:
		t := bst.floor(x.right, key)
		if t != nil {
			return t
		}
		return x
	default:
		return x
	}
}

// ceiling returns the smallest key greater than or equal to the given key.
func (bst *GeminiRbt[K, V]) Ceiling(key K) (K, bool) {
	x := bst.ceiling(bst.root, key)
	if x == nil {
		var zero K
		return zero, false
	}
	return x.key, true
}

func (bst *GeminiRbt[K, V]) ceiling(x *Node[K, V], key K) *Node[K, V] {
	if x == nil {
		return nil
	}
	switch {
	case key > x.key:
		return bst.ceiling(x.right, key)
	case key < x.key:
		t := bst.ceiling(x.left, key)
		if t != nil {
			return t
		}
		return x
	default:
		return x
	}
}

// select returns the key of rank k.
func (bst *GeminiRbt[K, V]) Select(k int) (K, bool) {
	x := bst.selectKey(bst.root, k)
	if x == nil {
		var zero K
		return zero, false
	}
	return x.key, true
}

func (bst *GeminiRbt[K, V]) selectKey(x *Node[K, V], k int) *Node[K, V] {
	if x == nil {
		return nil
	}
	t := bst.size(x.left)
	switch {
	case t > k:
		return bst.selectKey(x.left, k)
	case t < k:
		return bst.selectKey(x.right, k-t-1)
	default:
		return x
	}
}

// rank returns the number of keys less than key.
func (bst *GeminiRbt[K, V]) Rank(key K) int {
	return bst.rank(bst.root, key)
}

func (bst *GeminiRbt[K, V]) rank(x *Node[K, V], key K) int {
	if x == nil {
		return 0
	}
	switch {
	case key < x.key:
		return bst.rank(x.left, key)
	case key > x.key:
		return 1 + bst.size(x.left) + bst.rank(x.right, key)
	default:
		return bst.size(x.left)
	}
}

// deleteMin deletes the smallest key in the BST.
func (bst *GeminiRbt[K, V]) DeleteMin() {
	bst.root = bst.deleteMin(bst.root)
}

func (bst *GeminiRbt[K, V]) deleteMin(h *Node[K, V]) *Node[K, V] {
	if h.left == nil {
		return h.right
	}
	if !isRed(h.left) && !isRed(h.left.left) {
		h = bst.moveRedLeft(h)
	}
	h.left = bst.deleteMin(h.left)
	return bst.balance(h)
}

// deleteMax deletes the largest key in the BST.
func (bst *GeminiRbt[K, V]) DeleteMax() {
	bst.root = bst.deleteMax(bst.root)
}

func (bst *GeminiRbt[K, V]) deleteMax(h *Node[K, V]) *Node[K, V] {
	if isRed(h.left) {
		h = bst.rotateRight(h)
	}
	if h.right == nil {
		return h.left
	}
	if !isRed(h.right) && !isRed(h.right.left) {
		h = bst.moveRedRight(h)
	}
	h.right = bst.deleteMax(h.right)
	return bst.balance(h)
}

// delete deletes the key-value pair with the given key.
func (bst *GeminiRbt[K, V]) Delete(key K) {
	bst.root = bst.delete(bst.root, key)
}

func (bst *GeminiRbt[K, V]) delete(h *Node[K, V], key K) *Node[K, V] {
	if h == nil {
		return nil
	}
	if key < h.key {
		if !isRed(h.left) && !isRed(h.left.left) {
			h = bst.moveRedLeft(h)
		}
		h.left = bst.delete(h.left, key)
	} else {
		if isRed(h.left) {
			h = bst.rotateRight(h)
		}
		if key == h.key && h.right == nil {
			return nil
		}
		if !isRed(h.right) && !isRed(h.right.left) {
			h = bst.moveRedRight(h)
		}
		if key == h.key {
			x := bst.min(h.right)
			h.key = x.key
			h.val = x.val
			h.right = bst.deleteMin(h.right)
		} else {
			h.right = bst.delete(h.right, key)
		}
	}
	return bst.balance(h)
}

// keys returns an iterator over the keys in the BST in ascending order.
func (bst *GeminiRbt[K, V]) Keys() []K {
	queue := make([]K, 0)
	bst.keys(bst.root, &queue)
	return queue
}

func (bst *GeminiRbt[K, V]) keys(x *Node[K, V], queue *[]K) {
	if x == nil {
		return
	}
	bst.keys(x.left, queue)
	*queue = append(*queue, x.key)
	bst.keys(x.right, queue)
}

// keys returns an iterator over the keys in the BST in ascending order.
func (bst *GeminiRbt[K, V]) KeysInOrder() []K {
	queue := make([]K, 0)
	bst.keysInOrder(bst.root, &queue)
	return queue
}

func (bst *GeminiRbt[K, V]) keysInOrder(x *Node[K, V], queue *[]K) {
	if x == nil {
		return
	}
	bst.keysInOrder(x.left, queue)
	*queue = append(*queue, x.key)
	bst.keysInOrder(x.right, queue)
}

// keys returns an iterator over the keys in the BST in level order.
func (bst *GeminiRbt[K, V]) KeysLevelOrder() []K {
	queue := make([]K, 0)
	bst.keysLevelOrder(bst.root, &queue)
	return queue
}

func (bst *GeminiRbt[K, V]) keysLevelOrder(x *Node[K, V], queue *[]K) {
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
func (bst *GeminiRbt[K, V]) KeysPreOrder() []K {
	queue := make([]K, 0)
	bst.keysPreOrder(bst.root, &queue)
	return queue
}

func (bst *GeminiRbt[K, V]) keysPreOrder(x *Node[K, V], queue *[]K) {
	if x == nil {
		return
	}
	*queue = append(*queue, x.key)
	bst.keysPreOrder(x.left, queue)
	bst.keysPreOrder(x.right, queue)
}

// keys returns an iterator over the keys in the BST in post order.
func (bst *GeminiRbt[K, V]) KeysPostOrder() []K {
	queue := make([]K, 0)
	bst.keysPostOrder(bst.root, &queue)
	return queue
}

func (bst *GeminiRbt[K, V]) keysPostOrder(x *Node[K, V], queue *[]K) {
	if x == nil {
		return
	}
	bst.keysPostOrder(x.left, queue)
	bst.keysPostOrder(x.right, queue)
	*queue = append(*queue, x.key)
}

// isRed returns true if the given node is red.
func isRed[K constraints.Ordered, V any](x *Node[K, V]) bool {
	if x == nil {
		return false
	}
	return !x.color
}

// rotateLeft performs a left rotation.
func (bst *GeminiRbt[K, V]) rotateLeft(h *Node[K, V]) *Node[K, V] {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = true
	x.N = h.N
	h.N = 1 + bst.size(h.left) + bst.size(h.right)
	return x
}

// rotateRight performs a right rotation.
func (bst *GeminiRbt[K, V]) rotateRight(h *Node[K, V]) *Node[K, V] {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = true
	x.N = h.N
	h.N = 1 + bst.size(h.left) + bst.size(h.right)
	return x
}

// flipColors flips the colors of a node and its two children.
func (bst *GeminiRbt[K, V]) flipColors(h *Node[K, V]) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

// moveRedLeft moves a red link to the left.
func (bst *GeminiRbt[K, V]) moveRedLeft(h *Node[K, V]) *Node[K, V] {
	bst.flipColors(h)
	if isRed(h.right.left) {
		h.right = bst.rotateRight(h.right)
		h = bst.rotateLeft(h)
		bst.flipColors(h)
	}
	return h
}

// moveRedRight moves a red link to the right.
func (bst *GeminiRbt[K, V]) moveRedRight(h *Node[K, V]) *Node[K, V] {
	bst.flipColors(h)
	if isRed(h.left.left) {
		h = bst.rotateRight(h)
		bst.flipColors(h)
	}
	return h
}

// balance restores red-black tree invariant.
func (bst *GeminiRbt[K, V]) balance(h *Node[K, V]) *Node[K, V] {
	if isRed(h.right) {
		h = bst.rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = bst.rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		bst.flipColors(h)
	}
	h.N = 1 + bst.size(h.left) + bst.size(h.right)
	return h
}

// Print prints the BST in level order.
func (bst *GeminiRbt[K, V]) Print() {
	bst.print(bst.root, 0)
}

func (bst *GeminiRbt[K, V]) print(x *Node[K, V], depth int) {
	if x == nil {
		return
	}
	bst.print(x.right, depth+1)
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
	if x.color {
		fmt.Printf("R: %v -> %v\n", x.key, x.val)
	} else {
		fmt.Printf("B: %v -> %v\n", x.key, x.val)
	}
	bst.print(x.left, depth+1)
}

// PrintInOrder prints the BST in inorder.
func (bst *GeminiRbt[K, V]) PrintInOrder() {
	bst.printInOrder(bst.root)
}

func (bst *GeminiRbt[K, V]) printInOrder(x *Node[K, V]) {
	if x == nil {
		return
	}
	bst.printInOrder(x.left)
	fmt.Printf("%v -> %v\n", x.key, x.val)
	bst.printInOrder(x.right)
}

// PrintPreOrder prints the BST in preorder.
func (bst *GeminiRbt[K, V]) PrintPreOrder() {
	bst.printPreOrder(bst.root)
}

func (bst *GeminiRbt[K, V]) printPreOrder(x *Node[K, V]) {
	if x == nil {
		return
	}
	fmt.Printf("%v -> %v\n", x.key, x.val)
	bst.printPreOrder(x.left)
	bst.printPreOrder(x.right)
}

// PrintPostOrder prints the BST in postorder.
func (bst *GeminiRbt[K, V]) PrintPostOrder() {
	bst.printPostOrder(bst.root)
}

func (bst *GeminiRbt[K, V]) printPostOrder(x *Node[K, V]) {
	if x == nil {
		return
	}
	bst.printPostOrder(x.left)
	bst.printPostOrder(x.right)
	fmt.Printf("%v -> %v\n", x.key, x.val)
}

// PrintLevelOrder prints the BST in level order.
func (bst *GeminiRbt[K, V]) PrintLevelOrder() {
	bst.printLevelOrder(bst.root)
}

func (bst *GeminiRbt[K, V]) printLevelOrder(x *Node[K, V]) {
	if x == nil {
		return
	}
	q := []*Node[K, V]{x}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		fmt.Printf("%v -> %v\n", n.key, n.val)
		if n.left != nil {
			q = append(q, n.left)
		}
		if n.right != nil {
			q = append(q, n.right)
		}
	}
}

// Height returns the height of the BST.
func (bst *GeminiRbt[K, V]) Height() int {
	return bst.height(bst.root)
}

func (bst *GeminiRbt[K, V]) height(x *Node[K, V]) int {
	if x == nil {
		return -1
	}
	return 1 + max(bst.height(x.left), bst.height(x.right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (bst *GeminiRbt[K, V]) GetAll() []rbt.KeyValuePair[K, V] {
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

func (t *GeminiRbt[K, V]) Iterator() func(yield func(rbt.KeyValuePair[K, V]) bool) {
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
