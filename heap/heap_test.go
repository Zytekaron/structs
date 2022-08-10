package heap

import (
	"github.com/zytekaron/structs"
	"math/rand"
	"sort"
	"testing"
)

func TestHeap(t *testing.T) {
	// if randMax is less than count, duplicates are unavoidable. good for testing.
	const count = 256
	const randMax = 128

	heap := New[int](structs.CompareOrdered[int])

	// test heap integrity using 256 random elements, which will contain
	// duplicates of at least 2 elements (256 values from 0 to 128)
	random := make([]int, count)
	for i := 0; i < len(random); i++ {
		random[i] = rand.Intn(randMax)
	}
	for _, num := range random {
		heap.Push(num)
	}
	// sort the slice to ensure dequeued values will be in order
	sort.Slice(random, func(i, j int) bool {
		return random[i] < random[j]
	})

	for i := 0; i < len(random); i++ {
		value := heap.Pop()
		expect := random[i]
		if value != expect {
			t.Errorf("expected %d but got %d at index %d", value, expect, i)
		}
	}
}

func TestNewCap(t *testing.T) {
	const capacity = 64
	h := NewCap[int](capacity, structs.CompareOrdered[int])
	if len(h.data) != capacity {
		t.Errorf("expected capacity of %d but got %d", capacity, len(h.data))
	}
}
