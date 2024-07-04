package array

func SliceIterator[T any](items []T) func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for _, v := range items {
			if !yield(v) {
				return
			}
		}
	}
}
