package gemini

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestEmptyRbt(t *testing.T) {
	rbt := NewRBT[int, string]()
	if !rbt.IsEmpty() {
		t.Errorf("IsEmpty() == %v; want true", rbt.IsEmpty())
	}
}

func TestPutOneRbt(t *testing.T) {
	rbt := NewRBT[int, string]()
	rbt.Put(1, "one")
	if rbt.IsEmpty() {
		t.Errorf("IsEmpty() == %v; want false", rbt.IsEmpty())
	}

	v, ok := rbt.Get(1)
	if !ok {
		t.Errorf("Get(1) = %v; want 'one'", v)
	}
	if v != "one" {
		t.Errorf("Get(1) = %v; want 'one'", v)
	}
}

func TestPutThreeRbt(t *testing.T) {
	rbt := NewRBT[int, string]()
	rbt.Put(1, "one")
	rbt.Put(2, "two")
	rbt.Put(3, "three")
	if rbt.IsEmpty() {
		t.Errorf("IsEmpty() == %v; want false", rbt.IsEmpty())
	}
}

func TestContains3Rbt(t *testing.T) {
	rbt := NewRBT[int, string]()
	rbt.Put(1, "one")
	rbt.Put(2, "two")
	rbt.Put(3, "three")

	if x, ok := rbt.Get(1); !ok {
		t.Errorf("Get(1) == %v; want true : %v", x, ok)
	}

	if x, ok := rbt.Get(2); !ok {
		t.Errorf("Get(2) == %v; want true : %v", x, ok)
	}

	if x, ok := rbt.Get(3); !ok {
		t.Errorf("Get(3) == %v; want true : %v", x, ok)
	}

	if x, ok := rbt.Get(4); ok {
		t.Errorf("Get(4) == %v; want false", x)
	}
}

// test that the keys are returned in order
func TestContainsKeysRbt(t *testing.T) {
	rbt := NewRBT[int, string]()
	rbt.Put(1, "one")
	rbt.Put(2, "two")
	rbt.Put(3, "three")

	v, ok := rbt.Get(1)
	if !ok {
		t.Errorf("Get(%v) = %v; want 'one'", 1, v)
	}
	if v != "one" {
		t.Errorf("Get(%v) = %v; want 'one'", 1, v)
	}

	v, ok = rbt.Get(2)
	if !ok {
		t.Errorf("Get(%v) = %v; want 'two'", 2, v)
	}
	if v != "two" {
		t.Errorf("Get(%v) = %v; want 'two'", 2, v)
	}

	v, ok = rbt.Get(3)
	if !ok {
		t.Errorf("Get(%v) = %v; want 'three'", 3, v)
	}
	if v != "three" {
		t.Errorf("Get(%v) = %v; want 'three'", 3, v)
	}

}

// create a map of random keys and values
func TestRandomKeystRbt(t *testing.T) {
	rbt := NewRBT[int, string]()
	keys := make([]int, 100)
	values := make([]string, 100)
	for i := range keys {
		keys[i] = rand.Intn(100)
		values[i] = strconv.Itoa(keys[i])
	}

	for i := 0; i < len(keys); i++ {
		rbt.Put(keys[i], values[i])
	}

	for i := 0; i < len(keys); i++ {
		v, ok := rbt.Get(keys[i])
		if !ok {
			t.Errorf("Get(%v) = %v; want %v", keys[i], v, values[i])
		}
		if v != values[i] {
			t.Errorf("Get(%v) = %v; want %v", keys[i], v, values[i])
		}
	}
}

// test the Iterator with a large number of random keys
func TestIteratorRandom(t *testing.T) {
	rbt := NewRBT[int, string]()

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
		rbt.Put(k, v)
	}

	k := -1
	for r := range rbt.Iterator() {
		if r.Key < k {
			t.Errorf("Out of order(%v) = %v; ", k, r.Key)
		}
		t.Log(r.Key, r.Val)
		k = r.Key
	}
}
