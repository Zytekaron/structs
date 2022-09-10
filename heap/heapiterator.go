package heap

import "github.com/zytekaron/structs/list"

// todo: thoroughly test

type Iterator[V any] struct {
	heap  *Heap[V]
	index int

	unforgotten *list.List[V]

	lastReturned      *V
	lastReturnedIndex int
}

func (i *Iterator[V]) HasNext() bool {
	return i.index < i.heap.size ||
		(i.unforgotten != nil && i.unforgotten.IsEmpty())
}

func (i *Iterator[V]) Next() V {
	if i.index < i.heap.size {
		i.lastReturnedIndex = i.index
		i.index++
		return i.heap.data[i.lastReturnedIndex]
	}
	if i.unforgotten != nil {
		i.lastReturnedIndex = -1
		*i.lastReturned = i.unforgotten.Poll()
		return *i.lastReturned
	}
	panic("next called on empty iterator")
}

func (i *Iterator[V]) Remove() {
	if i.lastReturnedIndex != -1 {
		moved := i.heap.RemoveIndex(i.lastReturnedIndex)
		i.lastReturnedIndex = -1

		if i.unforgotten == nil {
			i.unforgotten = list.New[V](nil)
		}
		i.unforgotten.Add(moved)
		return
	}
	if i.lastReturned != nil {
		i.heap.Remove(*i.lastReturned)
		i.lastReturned = nil
	}
}
