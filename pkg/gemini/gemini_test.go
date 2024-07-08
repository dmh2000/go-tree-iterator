package gemini

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestEmptyGeminiRbt(t *testing.T) {
	bst := NewRBT[int, string]()
	if !bst.IsEmpty() {
		t.Errorf("IsEmpty() == %v; want true", bst.IsEmpty())
	}
}

func TestPutOneGeminiRbt(t *testing.T) {
	bst := NewRBT[int, string]()
	bst.Put(1, "one")
	if bst.IsEmpty() {
		t.Errorf("IsEmpty() == %v; want false", bst.IsEmpty())
	}

	v, ok := bst.Get(1)
	if !ok {
		t.Errorf("Get(1) = %v; want 'one'", v)
	}
	if v != "one" {
		t.Errorf("Get(1) = %v; want 'one'", v)
	}
}

func TestPutThreeGeminiRbt(t *testing.T) {
	bst := NewRBT[int, string]()
	bst.Put(1, "one")
	bst.Put(2, "two")
	bst.Put(3, "three")
	if bst.IsEmpty() {
		t.Errorf("IsEmpty() == %v; want false", bst.IsEmpty())
	}
}

func TestContains3GeminiRbt(t *testing.T) {
	bst := NewRBT[int, string]()
	bst.Put(1, "one")
	bst.Put(2, "two")
	bst.Put(3, "three")

	if x, ok := bst.Get(1); !ok {
		t.Errorf("Get(1) == %v; want true : %v", x, ok)
	}

	if x, ok := bst.Get(2); !ok {
		t.Errorf("Get(2) == %v; want true : %v", x, ok)
	}

	if x, ok := bst.Get(3); !ok {
		t.Errorf("Get(3) == %v; want true : %v", x, ok)
	}

	if x, ok := bst.Get(4); ok {
		t.Errorf("Get(4) == %v; want false", x)
	}
}

// test that the keys are returned in order
func TestContainsKeysGeminiRbt(t *testing.T) {
	bst := NewRBT[int, string]()
	bst.Put(1, "one")
	bst.Put(2, "two")
	bst.Put(3, "three")

	v, ok := bst.Get(1)
	if !ok {
		t.Errorf("Get(%v) = %v; want 'one'", 1, v)
	}
	if v != "one" {
		t.Errorf("Get(%v) = %v; want 'one'", 1, v)
	}

	v, ok = bst.Get(2)
	if !ok {
		t.Errorf("Get(%v) = %v; want 'two'", 2, v)
	}
	if v != "two" {
		t.Errorf("Get(%v) = %v; want 'two'", 2, v)
	}

	v, ok = bst.Get(3)
	if !ok {
		t.Errorf("Get(%v) = %v; want 'three'", 3, v)
	}
	if v != "three" {
		t.Errorf("Get(%v) = %v; want 'three'", 3, v)
	}

}

// create a map of random keys and values
func TestRandomKeysGeminitRbt(t *testing.T) {
	bst := NewRBT[int, string]()
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
		v, ok := bst.Get(keys[i])
		if !ok {
			t.Errorf("Get(%v) = %v; want %v", keys[i], v, values[i])
		}
		if v != values[i] {
			t.Errorf("Get(%v) = %v; want %v", keys[i], v, values[i])
		}
	}
}

// test the BSTIterator with a large number of random keys
func TestGeminitIteratorRandom(t *testing.T) {
	bst := NewRBT[int, string]()

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
}
