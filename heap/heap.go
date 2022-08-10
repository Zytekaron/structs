package heap

import (
	"github.com/zytekaron/structs"
	"golang.org/x/exp/constraints"
)

type Heap[V any] struct {
	capFn structs.CapacityFunc
	cmp   structs.CompareFunc[V]
	data  []V
	size  int
}

func New[V any](cmp structs.CompareFunc[V]) *Heap[V] {
	return &Heap[V]{
		capFn: structs.DoubleCapacity,
		cmp:   cmp,
		data:  nil,
		size:  0,
	}
}

func NewOrdered[V constraints.Ordered]() *Heap[V] {
	return &Heap[V]{
		capFn: structs.DoubleCapacity,
		cmp:   structs.CompareOrdered[V],
		data:  nil,
		size:  0,
	}
}

func NewCap[V any](capacity int, cmp structs.CompareFunc[V]) *Heap[V] {
	return &Heap[V]{
		capFn: structs.DoubleCapacity,
		cmp:   cmp,
		data:  make([]V, capacity),
		size:  0,
	}
}

func NewOrderedCap[V constraints.Ordered](capacity int) *Heap[V] {
	return &Heap[V]{
		capFn: structs.DoubleCapacity,
		cmp:   structs.CompareOrdered[V],
		data:  make([]V, capacity),
		size:  0,
	}
}

func From[V any](data []V, cmp structs.CompareFunc[V]) *Heap[V] {
	h := &Heap[V]{
		capFn: structs.DoubleCapacity,
		cmp:   cmp,
		data:  data,
		size:  len(data),
	}
	for i := 0; i < h.size; i++ {
		h.bubbleDown(i)
	}
	return h
}

func FromOrdered[V constraints.Ordered](data []V) *Heap[V] {
	h := &Heap[V]{
		capFn: structs.DoubleCapacity,
		cmp:   structs.CompareOrdered[V],
		data:  data,
		size:  len(data),
	}
	for i := 0; i < h.size; i++ {
		h.bubbleDown(i)
	}
	return h
}

func FromHeapSlice[V any](data []V, cmp structs.CompareFunc[V]) *Heap[V] {
	return &Heap[V]{
		capFn: structs.DoubleCapacity,
		cmp:   cmp,
		data:  data,
		size:  len(data),
	}
}

func FromOrderedHeapSlice[V constraints.Ordered](data []V) *Heap[V] {
	return &Heap[V]{
		capFn: structs.DoubleCapacity,
		cmp:   structs.CompareOrdered[V],
		data:  data,
		size:  len(data),
	}
}

func (h *Heap[V]) SetCapFunc(capFn structs.CapacityFunc) {
	h.capFn = capFn
}

func (h *Heap[V]) IsEmpty() bool {
	return h.size == 0
}

func (h *Heap[V]) Peek() V {
	if h.IsEmpty() {
		panic("peek called on empty heap")
	}
	return h.data[0]
}

func (h *Heap[V]) Pop() V {
	if h.IsEmpty() {
		panic("pop called on empty heap")
	}
	return h.RemoveIndex(0)
}

func (h *Heap[V]) Push(value V) {
	h.size++
	if h.size > len(h.data) {
		size := h.capFn(len(h.data), len(h.data)+1)
		h.data = structs.Realloc(size, h.data)
	}
	h.data[h.size-1] = value
	h.bubbleUp(h.size - 1)
}

func (h *Heap[V]) Index(value V) int {
	for i, val := range h.data {
		if h.cmp(val, value) == 0 {
			return i
		}
	}
	return -1
}

func (h *Heap[V]) Contains(value V) bool {
	return h.Index(value) >= 0
}

func (h *Heap[V]) UpdateIndex(i int, newValue V) {
	if !h.isValidIndex(i) {
		panic("index outside of heap range")
	}
	oldValue := h.data[i]
	h.data[i] = newValue
	if h.cmp(newValue, oldValue) < 0 {
		h.bubbleUp(i)
	} else {
		h.bubbleDown(i)
	}
}

func (h *Heap[V]) Update(oldValue, newValue V) bool {
	i := h.Index(oldValue)
	if i < 0 {
		return false
	}
	h.UpdateIndex(i, newValue)
	return true
}

func (h *Heap[V]) RemoveIndex(i int) V {
	if !h.isValidIndex(i) {
		panic("index outside of heap range")
	}
	value := h.data[i]
	h.data[i] = h.data[h.size-1]
	h.size--
	h.bubbleDown(i)
	return value
}

func (h *Heap[V]) Remove(value V) bool {
	i := h.Index(value)
	if i < 0 {
		return false
	}
	h.RemoveIndex(i)
	return true
}

func (h *Heap[V]) Values() []V {
	return h.data[:h.size]
}

func (h *Heap[V]) Size() int {
	return h.size
}

func (h *Heap[V]) Clear() {
	h.size = 0
}

// todo: consider keeping heap's Clone()
//func (h *Heap[V]) Clone() *Heap[V] {
//	data := make([]V, h.size)
//	copy(data, h.data[:h.size])
//	return &Heap[V]{
//		cmp:  h.cmp,
//		data: data,
//		size: h.size,
//	}
//}

func (h *Heap[V]) isValidIndex(i int) bool {
	return i >= 0 && i < h.size
}

func (h *Heap[V]) less(i, j int) bool {
	return h.cmp(h.data[i], h.data[j]) < 0
}

func (h *Heap[V]) bubbleUp(i int) {
	parent := (i - 1) / 2
	for i > 0 && h.less(i, parent) {
		h.data[i], h.data[parent] = h.data[parent], h.data[i]

		i = parent
		parent = (i - 1) / 2
	}
}

func (h *Heap[V]) bubbleDown(i int) {
	left := 2*i + 1
	right := 2*i + 2
	for left < h.size && h.less(left, i) ||
		right < h.size && h.less(right, i) {

		min := left
		if right < h.size && h.less(right, left) {
			min = right
		}

		h.data[i], h.data[min] = h.data[min], h.data[i]

		i = min
		left = 2*i + 1
		right = 2*i + 2
	}
}
