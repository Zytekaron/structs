package queue

import (
	"github.com/zytekaron/structs"
	"github.com/zytekaron/structs/heap"
)

type PriorityQueue[V any] struct {
	heap *heap.Heap[V]
}

func NewPriority[V any](cmp structs.CompareFunc[V]) *PriorityQueue[V] {
	return &PriorityQueue[V]{
		heap: heap.New(cmp),
	}
}

func NewPriorityCap[V any](capacity int, cmp structs.CompareFunc[V]) *PriorityQueue[V] {
	return &PriorityQueue[V]{
		heap: heap.NewCap(capacity, cmp),
	}
}

func FromPriorityHeap[V any](h *heap.Heap[V]) *PriorityQueue[V] {
	return &PriorityQueue[V]{
		heap: h,
	}
}

func (p *PriorityQueue[V]) IsEmpty() bool {
	return p.heap.IsEmpty()
}

func (p *PriorityQueue[V]) Enqueue(value V) {
	p.heap.Push(value)
}

func (p *PriorityQueue[V]) Peek() V {
	return p.heap.Peek()
}

func (p *PriorityQueue[V]) Dequeue() V {
	return p.heap.Pop()
}

func (p *PriorityQueue[V]) Contains(value V) bool {
	return p.heap.Contains(value)
}

func (p *PriorityQueue[V]) Size() int {
	return p.heap.Size()
}

func (p *PriorityQueue[V]) Clear() {
	p.heap.Clear()
}

func (p *PriorityQueue[V]) Values() []V {
	return p.heap.Values()
}
