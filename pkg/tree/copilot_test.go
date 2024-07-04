package tree

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestCreateRBT(t *testing.T) {
	bst := NewCopilotRBT[int, string]()
	if bst == nil {
		t.Errorf("NewGeminiRBT() = %v; want a new red-black tree", bst)
	}
}

func TestEmptyRBT(t *testing.T) {
	bst := NewCopilotRBT[int, string]()
	if !bst.IsEmpty() {
		t.Errorf("IsEmpty() == %v; want true", bst.IsEmpty())
	}
}

func TestPutOneRBT(t *testing.T) {
	bst := NewCopilotRBT[int, string]()
	bst.Put(1, "one")
	if bst.IsEmpty() {
		t.Errorf("IsEmpty() == %v; want false", bst.IsEmpty())
	}

	v, err := bst.Get(1)
	if err != nil {
		t.Errorf("Get(1) = %v; want 'one'", v)
	}
	if v != "one" {
		t.Errorf("Get(1) = %v; want 'one'", v)
	}
}

func TestPutThreeRBT(t *testing.T) {
	bst := NewCopilotRBT[int, string]()
	bst.Put(1, "one")
	bst.Put(2, "two")
	bst.Put(3, "three")
	if bst.IsEmpty() {
		t.Errorf("IsEmpty() == %v; want false", bst.IsEmpty())
	}
}

func TestContains3RBT(t *testing.T) {
	bst := NewCopilotRBT[int, string]()
	bst.Put(1, "one")
	bst.Put(2, "two")
	bst.Put(3, "three")

	if x, err := bst.Get(1); err != nil {
		t.Errorf("Get(1) == %v; want true : %v", x, err)
	}

	if x, err := bst.Get(2); err != nil {
		t.Errorf("Get(2) == %v; want true : %v", x, err)
	}

	if x, err := bst.Get(3); err != nil {
		t.Errorf("Get(3) == %v; want true : %v", x, err)
	}

	if x, err := bst.Get(4); err == nil {
		t.Errorf("Get(4) == %v; want false", x)
	}
}

// test that the keys are returned in order
func TestContainsKeys(t *testing.T) {
	bst := NewCopilotRBT[int, string]()
	bst.Put(1, "one")
	bst.Put(2, "two")
	bst.Put(3, "three")

	x, err := bst.Keys()
	if err != nil {
		t.Errorf("Keys() == %v; want [1, 2, 3]", x)
	}

	if x[0] != 1 {
		t.Errorf("Keys() == %v; want [1, 2, 3]", x)
	}

	v, err := bst.Get(x[0])
	if err != nil {
		t.Errorf("Get(%v) = %v; want 'one'", x[0], v)
	}
	if v != "one" {
		t.Errorf("Get(%v) = %v; want 'one'", x[0], v)
	}

	if x[1] != 2 {
		t.Errorf("Keys() == %v; want [1, 2, 3]", x)
	}
	if x[2] != 3 {
		t.Errorf("Keys() == %v; want [1, 2, 3]", x)
	}
}

// create a map of random keys and values
func TestRandomKeys(t *testing.T) {
	bst := NewCopilotRBT[int, string]()
	keys := make([]int, 100)
	values := make([]string, 100)
	for i := range keys {
		keys[i] = rand.Intn(100)
		values[i] = strconv.Itoa(keys[i])
	}

	for i := 0; i < len(keys); i++ {
		bst.Put(keys[i], values[i])
	}

	for i := 0; i < len(keys); i++ {
		v, err := bst.Get(keys[i])
		if err != nil {
			t.Errorf("Get(%v) = %v; want %v", keys[i], v, values[i])
		}
		if v != values[i] {
			t.Errorf("Get(%v) = %v; want %v", keys[i], v, values[i])
		}
	}
}

// test the BSTIterator
func TestBSTIterator(t *testing.T) {
	bst := NewCopilotRBT[int, string]()
	bst.Put(1, "one")
	bst.Put(2, "two")
	bst.Put(3, "three")

	keys := make([]int, 0)
	values := make([]string, 0)
	bst.CopilotIterator()(func(k int, v string) bool {
		keys = append(keys, k)
		values = append(values, v)
		return true
	})

	if keys[0] != 1 {
		t.Errorf("BSTIterator() == %v; want [1, 2, 3]", keys)
	}
	if keys[1] != 2 {
		t.Errorf("BSTIterator() == %v; want [1, 2, 3]", keys)
	}
	if keys[2] != 3 {
		t.Errorf("BSTIterator() == %v; want [1, 2, 3]", keys)
	}
}

// test the BSTIterator with a large number of random keys
func TestBSTIteratorRandom(t *testing.T) {
	bst := NewCopilotRBT[int, string]()

	m := make(map[int]string)
	for i := 0; i < 100; i++ {
		k := rand.Intn(100)
		v := strconv.Itoa(k)
		m[k] = v
	}

	// iterate over the map m
	t.Log("--- random keys")
	for k, v := range m {
		t.Log(k, v)
		bst.Put(k, v)
	}

	// iterate over the BST and get the keys in order
	keys2 := make([]int, 0)
	values2 := make([]string, 0)

	iter := bst.CopilotIterator()
	iter(func(k int, v string) bool {
		keys2 = append(keys2, k)
		values2 = append(values2, v)
		return true
	})

	t.Log("--- in order keys")

	// check that the keys are in order
	k := keys2[0]
	v := values2[0]
	t.Log(k, v)
	for i := 1; i < len(keys2); i++ {
		if keys2[i] <= k {
			t.Errorf("BSTIterator() == %v; want %v", keys2, keys2)
		}
		k = keys2[i]
		v = values2[i]
		t.Log(k, v)
	}
}
