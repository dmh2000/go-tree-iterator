package main

import (
	"fmt"

	ch "sqirvy.xyz/go-tree-iterator/chatgpt"
	cp "sqirvy.xyz/go-tree-iterator/copilot"
	gm "sqirvy.xyz/go-tree-iterator/gemini"
	rbt "sqirvy.xyz/go-tree-iterator/rbt"
)

func runRbt(s string, t rbt.RBT[int, string]) {

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
		t.Put(kv.k, kv.v)
	}

	fmt.Println(s, "Iterator")
	// iterate over the tree in order
	for r := range t.Iterator() {
		fmt.Println(r)
	}

	fmt.Println(s, "GetAll")
	a := t.GetAll()
	for _, r := range a {
		fmt.Println(r)
	}

}

func main() {
	runRbt("=== Copilot ===", cp.NewRBT[int, string]())
	runRbt("=== Gemini ===", gm.NewRBT[int, string]())
	runRbt("=== ChatGpt ===", ch.NewRBT[int, string]())
}
