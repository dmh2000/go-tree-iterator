package main

import (
	"fmt"

	arr "sqirvy.xyz/go-tree-iterator/array"
)

func main() {
	items := []int{1, 2, 3, 4, 5}

	for v := range arr.SliceIterator(items) {
		fmt.Println(v)
	}

	s := []string{"a", "b", "c", "d", "e"}
	for u := range arr.SliceIterator(s) {
		fmt.Println(u)
	}
}
