package queue

import (
	"github.com/zytekaron/structs"
	"math/rand"
	"sort"
	"testing"
)

func TestNewPriority(t *testing.T) {
	pq := NewPriority[int](structs.CompareOrdered[int])
	if !pq.IsEmpty() {
		t.Error("expected instantiated priority queue to be empty, got size", pq.Size())
	}
}

func TestNewPriorityCap(t *testing.T) {
	pq := NewPriorityCap[int](32, structs.CompareOrdered[int])
	if !pq.IsEmpty() {
		t.Error("expected instantiated priority queue to be empty, got size", pq.Size())
	}
}

func TestPriorityQueue(t *testing.T) {
	// if randMax is less than count, duplicates are unavoidable. good for testing.
	const count = 256
	const randMax = 128

	h := NewPriority[int](structs.CompareOrdered[int])

	// test pq ordering using many elements, probably containing duplicates
	random := make([]int, count)
	for i := 0; i < len(random); i++ {
		random[i] = rand.Intn(randMax)
	}
	for _, num := range random {
		h.Enqueue(num)
	}
	// sort the slice to ensure dequeued values will be in order
	sort.Slice(random, func(i, j int) bool {
		return random[i] < random[j]
	})

	for i := 0; i < len(random); i++ {
		value := h.Dequeue()
		expect := random[i]
		if value != expect {
			t.Errorf("expected %d but got %d at index %d", value, expect, i)
		}
	}
}
