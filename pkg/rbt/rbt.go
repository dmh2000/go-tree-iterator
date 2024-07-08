package rbt

import "golang.org/x/exp/constraints"

type KeyValuePair[K constraints.Ordered, V any] struct {
	Key K
	Val V
}

type RBT[K constraints.Ordered, V any] interface {
	Put(key K, val V)
	Get(key K) (V, bool)
	IsEmpty() bool
	GetAll() []KeyValuePair[K, V]
	Iterator() func(yield func(KeyValuePair[K, V]) bool)
}
