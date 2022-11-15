package queue

import (
	"github.com/zytekaron/structs"
	"github.com/zytekaron/structs/list"
	"golang.org/x/exp/constraints"
)

// Queue is an implementation of a double-ended queue
// backed by *list.List
type Queue[V any] struct {
	data *list.List[V]
}

func New[V any](eq structs.EqualFunc[V]) *Queue[V] {
	return &Queue[V]{
		data: list.New[V](eq),
	}
}

func NewOrdered[V constraints.Ordered]() *Queue[V] {
	return &Queue[V]{
		data: list.NewOrdered[V](),
	}
}

func (q *Queue[V]) Add(value V) {
	q.data.AddLast(value)
}

func (q *Queue[V]) Element() V {
	return q.Peek()
}

func (q *Queue[V]) Offer(value V) bool {
	return q.data.Offer(value)
}

func (q *Queue[V]) Peek() V {
	return q.data.Peek()
}

func (q *Queue[V]) Poll() V {
	return q.data.Poll()
}

func (q *Queue[V]) RemoveHead() V {
	return q.data.RemoveHead()
}

func (q *Queue[V]) Contains(value V) bool {
	return q.data.Contains(value)
}

func (q *Queue[V]) Iterator() structs.Iterator[V] {
	return q.data.Iterator()
}

func (q *Queue[V]) DescendingIterator() structs.Iterator[V] {
	return q.data.DescendingIterator()
}

func (q *Queue[V]) IsEmpty() bool {
	return q.data.IsEmpty()
}

func (q *Queue[V]) Size() int {
	return q.data.Size()
}

func (q *Queue[V]) Clear() {
	q.data.Clear()
}

func (q *Queue[V]) Values() []V {
	return q.data.Values()
}
