package gemini

import "fmt"

type Node[K comparable, V any] struct {
	key         K
	val         V
	N           int
	color       bool // color of parent link
	left, right *Node[K, V]
}

type RedBlackBST[K comparable, V any] struct {
	root *Node[K, V]
}

func (bst *RedBlackBST[K, V]) IsEmpty() bool {
	return bst.root == nil
}

func (bst *RedBlackBST[K, V]) Size() int {
	return bst.size(bst.root)
}

func (bst *RedBlackBST[K, V]) size(x *Node[K, V]) int {
	if x == nil {
		return 0
	}
	return x.N
}

func (bst *RedBlackBST[K, V]) Get(key K) (V, bool) {
	x := bst.get(bst.root, key)
	if x == nil {
		var zero V
		return zero, false
	}
	return x.val, true
}

func (bst *RedBlackBST[K, V]) get(x *Node[K, V], key K) *Node[K, V] {
	for x != nil {
		cmp := compare(key, x.key)
		if cmp < 0 {
			x = x.left
		} else if cmp > 0 {
			x = x.right
		} else {
			return x
		}
	}
	return nil
}

func (bst *RedBlackBST[K, V]) Put(key K, val V) {
	bst.root = bst.put(bst.root, key, val)
}

func (bst *RedBlackBST[K, V]) put(h *Node[K, V], key K, val V) *Node[K, V] {
	if h == nil {
		return &Node[K, V]{key: key, val: val, N: 1, color: false}
	}
	cmp := compare(key, h.key)
	if cmp < 0 {
		h.left = bst.put(h.left, key, val)
	} else if cmp > 0 {
		h.right = bst.put(h.right, key, val)
	} else {
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

func (bst *RedBlackBST[K, V]) Min() (K, bool) {
	if bst.IsEmpty() {
		var zero K
		return zero, false
	}
	return bst.min(bst.root).key, true
}

func (bst *RedBlackBST[K, V]) min(x *Node[K, V]) *Node[K, V] {
	if x.left == nil {
		return x
	}
	return bst.min(x.left)
}

func (bst *RedBlackBST[K, V]) Max() (K, bool) {
	if bst.IsEmpty() {
		var zero K
		return zero, false
	}
	return bst.max(bst.root).key, true
}

func (bst *RedBlackBST[K, V]) max(x *Node[K, V]) *Node[K, V] {
	if x.right == nil {
		return x
	}
	return bst.max(x.right)
}

func (bst *RedBlackBST[K, V]) Floor(key K) (K, bool) {
	x := bst.floor(bst.root, key)
	if x == nil {
		var zero K
		return zero, false
	}
	return x.key, true
}

func (bst *RedBlackBST[K, V]) floor(x *Node[K, V], key K) *Node[K, V] {
	if x == nil {
		return nil
	}
	cmp := compare(key, x.key)
	if cmp < 0 {
		return bst.floor(x.left, key)
	} else if cmp == 0 {
		return x
	}
	t := bst.floor(x.right, key)
	if t != nil {
		return t
	}
	return x
}

func (bst *RedBlackBST[K, V]) Ceiling(key K) (K, bool) {
	x := bst.ceiling(bst.root, key)
	if x == nil {
		var zero K
		return zero, false
	}
	return x.key, true
}

func (bst *RedBlackBST[K, V]) ceiling(x *Node[K, V], key K) *Node[K, V] {
	if x == nil {
		return nil
	}
	cmp := compare(key, x.key)
	if cmp > 0 {
		return bst.ceiling(x.right, key)
	} else if cmp == 0 {
		return x
	}
	t := bst.ceiling(x.left, key)
	if t != nil {
		return t
	}
	return x
}

func (bst *RedBlackBST[K, V]) Select(k int) (K, bool) {
	x := bst.selectK(bst.root, k)
	if x == nil {
		var zero K
		return zero, false
	}
	return x.key, true
}

func (bst *RedBlackBST[K, V]) selectK(x *Node[K, V], k int) *Node[K, V] {
	if x == nil {
		return nil
	}
	t := bst.size(x.left)
	if t > k {
		return bst.selectK(x.left, k)
	} else if t < k {
		return bst.selectK(x.right, k-t-1)
	} else {
		return x
	}
}

func (bst *RedBlackBST[K, V]) Rank(key K) int {
	return bst.rank(bst.root, key)
}

func (bst *RedBlackBST[K, V]) rank(x *Node[K, V], key K) int {
	if x == nil {
		return 0
	}
	cmp := compare(key, x.key)
	if cmp < 0 {
		return bst.rank(x.left, key)
	} else if cmp > 0 {
		return 1 + bst.size(x.left) + bst.rank(x.right, key)
	} else {
		return bst.size(x.left)
	}
}

func (bst *RedBlackBST[K, V]) DeleteMin() {
	if bst.IsEmpty() {
		return
	}
	bst.root = bst.deleteMin(bst.root)
}

func (bst *RedBlackBST[K, V]) deleteMin(h *Node[K, V]) *Node[K, V] {
	if h.left == nil {
		return h.right
	}
	if !isRed(h.left) && !isRed(h.left.left) {
		h = bst.moveRedLeft(h)
	}
	h.left = bst.deleteMin(h.left)
	return bst.balance(h)
}

func (bst *RedBlackBST[K, V]) DeleteMax() {
	if bst.IsEmpty() {
		return
	}
	bst.root = bst.deleteMax(bst.root)
}

func (bst *RedBlackBST[K, V]) deleteMax(h *Node[K, V]) *Node[K, V] {
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

func (bst *RedBlackBST[K, V]) Delete(key K) {
	if !bst.Contains(key) {
		return
	}
	if !isRed(bst.root.left) && !isRed(bst.root.right) {
		bst.root.color = false
	}
	bst.root = bst.delete(bst.root, key)
	if !bst.IsEmpty() {
		bst.root.color = true
	}
}

func (bst *RedBlackBST[K, V]) delete(h *Node[K, V], key K) *Node[K, V] {
	if compare(key, h.key) < 0 {
		if !isRed(h.left) && !isRed(h.left.left) {
			h = bst.moveRedLeft(h)
		}
		h.left = bst.delete(h.left, key)
	} else {
		if isRed(h.left) {
			h = bst.rotateRight(h)
		}
		if compare(key, h.key) == 0 && h.right == nil {
			return nil
		}
		if !isRed(h.right) && !isRed(h.right.left) {
			h = bst.moveRedRight(h)
		}
		if compare(key, h.key) == 0 {
			var x *Node[K, V]
			x = bst.min(h.right)
			h.key = x.key
			h.val = x.val
			h.right = bst.deleteMin(h.right)
		} else {
			h.right = bst.delete(h.right, key)
		}
	}
	return bst.balance(h)
}

func (bst *RedBlackBST[K, V]) Contains(key K) bool {
	return bst.get(bst.root, key) != nil
}

func (bst *RedBlackBST[K, V]) Keys() []K {
	queue := make([]K, 0)
	bst.keys(bst.root, &queue)
	return queue
}

func (bst *RedBlackBST[K, V]) keys(x *Node[K, V], queue *[]K) {
	if x == nil {
		return
	}
	bst.keys(x.left, queue)
	*queue = append(*queue, x.key)
	bst.keys(x.right, queue)
}

func (bst *RedBlackBST[K, V]) KeysInOrder(lo K, hi K) []K {
	queue := make([]K, 0)
	bst.keysInOrder(bst.root, lo, hi, &queue)
	return queue
}

func (bst *RedBlackBST[K, V]) keysInOrder(x *Node[K, V], lo K, hi K, queue *[]K) {
	if x == nil {
		return
	}
	cmpLo := compare(lo, x.key)
	cmpHi := compare(hi, x.key)
	if cmpLo < 0 {
		bst.keysInOrder(x.left, lo, hi, queue)
	}
	if cmpLo <= 0 && cmpHi >= 0 {
		*queue = append(*queue, x.key)
	}
	if cmpHi > 0 {
		bst.keysInOrder(x.right, lo, hi, queue)
	}
}

func (bst *RedBlackBST[K, V]) SizeInOrder(lo K, hi K) int {
	return bst.sizeInOrder(bst.root, lo, hi)
}

func (bst *RedBlackBST[K, V]) sizeInOrder(x *Node[K, V], lo K, hi K) int {
	if x == nil {
		return 0
	}
	cmpLo := compare(lo, x.key)
	cmpHi := compare(hi, x.key)
	if cmpLo < 0 {
		return bst.sizeInOrder(x.left, lo, hi)
	}
	if cmpLo <= 0 && cmpHi >= 0 {
		return 1 + bst.sizeInOrder(x.left, lo, hi) + bst.sizeInOrder(x.right, lo, hi)
	}
	if cmpHi > 0 {
		return bst.sizeInOrder(x.right, lo, hi)
	}
	return 0
}

func (bst *RedBlackBST[K, V]) Height() int {
	return bst.height(bst.root)
}

func (bst *RedBlackBST[K, V]) height(x *Node[K, V]) int {
	if x == nil {
		return -1
	}
	return 1 + max(bst.height(x.left), bst.height(x.right))
}

func (bst *RedBlackBST[K, V]) IsRed(x *Node[K, V]) bool {
	if x == nil {
		return false
	}
	return x.color
}

func (bst *RedBlackBST[K, V]) rotateLeft(h *Node[K, V]) *Node[K, V] {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = true
	x.N = h.N
	h.N = 1 + bst.size(h.left) + bst.size(h.right)
	return x
}

func (bst *RedBlackBST[K, V]) rotateRight(h *Node[K, V]) *Node[K, V] {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = true
	x.N = h.N
	h.N = 1 + bst.size(h.left) + bst.size(h.right)
	return x
}

func (bst *RedBlackBST[K, V]) flipColors(h *Node[K, V]) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

func (bst *RedBlackBST[K, V]) moveRedLeft(h *Node[K, V]) *Node[K, V] {
	bst.flipColors(h)
	if isRed(h.right.left) {
		h.right = bst.rotateRight(h.right)
		h = bst.rotateLeft(h)
		bst.flipColors(h)
	}
	return h
}

func (bst *RedBlackBST[K, V]) moveRedRight(h *Node[K, V]) *Node[K, V] {
	bst.flipColors(h)
	if isRed(h.left.left) {
		h = bst.rotateRight(h)
		bst.flipColors(h)
	}
	return h
}

func (bst *RedBlackBST[K, V]) balance(h *Node[K, V]) *Node[K, V] {
	if isRed(h.right) {
		h = bst.rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = bst.rotateRight(h)
	}
	if isRed(h.left) && isRed(h.right) {
		bst.flipColors(h)
	}
	return h
}

func isRed[K comparable, V any](x *Node[K, V]) bool {
	if x == nil {
		return false
	}
	return x.color
}

func compare[K comparable](v1 K, v2 K) int {
	if v1 < v2 {
		return -1
	} else if v1 > v2 {
		return 1
	} else {
		return 0
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (bst *RedBlackBST[K, V]) Print() {
	bst.print(bst.root, 0)
}

func (bst *RedBlackBST[K, V]) print(x *Node[K, V], depth int) {
	if x == nil {
		return
	}
	bst.print(x.right, depth+1)
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
	if x.color {
		fmt.Printf("R: %v %v\n", x.key, x.val)
	} else {
		fmt.Printf("B: %v %v\n", x.key, x.val)
	}
	bst.print(x.left, depth+1)
}
