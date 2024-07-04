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
func (t *CopilotRBT[K, V]) CopilotIterator() func(yield func(K, V) bool) {
	return func(yield func(K, V) bool) {
		var inorder func(*CopilotNode[K, V])
		inorder = func(n *CopilotNode[K, V]) {
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

func (t *GeminiRBT[K, V]) GeminiIterator() func(yield func(K, V) bool) {
	return func(yield func(K, V) bool) {
		var inorder func(*GeminiNode[K, V])
		inorder = func(n *GeminiNode[K, V]) {
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
