package tree

// a generic iterator for any type of items
func Iterator[T interface{}](items []T) func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for _, v := range items {
			if !yield(v) {
				return
			}
		}
	}
}

// an iterator for a BST
func (t *BST[K, V]) BSTIterator() func(yield func(K, V) bool) {
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
