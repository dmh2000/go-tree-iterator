package main

import (
	"fmt"

	tree "sqirvy.xyz/go-tree-iterator/tree"
)

func copilot() {

	// create a new red-black tree
	bst := tree.NewCopilotRBT[int, string]()

	// create a list of  key/value pairs in a random order
	kvs := []struct {
		k int
		v string
	}{
		{5, "five"},
		{3, "three"},
		{1, "one"},
		{6, "six"},
		{4, "four"},
		{2, "two"},
	}

	// insert the key/value pairs into the tree
	for _, kv := range kvs {
		bst.Put(kv.k, kv.v)
	}

	// iterate over the tree in order
	iter := bst.CopilotIterator()
	iter(func(k int, v string) bool {
		fmt.Println(k, v)
		return true
	})
}

func gemini() {

	// create a new red-black tree
	bst := tree.NewGeminiRBT[int, string]()

	// create a list of  key/value pairs in a random order
	kvs := []struct {
		k int
		v string
	}{
		{5, "five"},
		{3, "three"},
		{1, "one"},
		{6, "six"},
		{4, "four"},
		{2, "two"},
	}

	// insert the key/value pairs into the tree
	for _, kv := range kvs {
		bst.Put(kv.k, kv.v)
	}

	// iterate over the tree in order
	iter := bst.GeminiIterator()
	iter(func(k int, v string) bool {
		fmt.Println(k, v)
		return true
	})
}

func main() {
	copilot()
	gemini()
}
