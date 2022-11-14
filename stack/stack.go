package stack

import (
	"github.com/zytekaron/structs"
	"github.com/zytekaron/structs/list"
	"golang.org/x/exp/constraints"
)

type Stack[V any] struct {
	data *list.List[V]
}

func New[V any](eq structs.EqualFunc[V]) *Stack[V] {
	return &Stack[V]{
		data: list.New[V](eq),
	}
}

func NewOrdered[V constraints.Ordered]() *Stack[V] {
	return &Stack[V]{
		data: list.NewOrdered[V](),
	}
}

func From[V any](eq structs.EqualFunc[V], values ...V) *Stack[V] {
	return &Stack[V]{
		data: list.Of(eq, values...),
	}
}

func FromOrdered[V constraints.Ordered](values ...V) *Stack[V] {
	return &Stack[V]{
		data: list.OfOrdered(values...),
	}
}

func (s *Stack[V]) Add(value V) bool {
	return s.data.Add(value)
}

func (s *Stack[V]) AddAll(other structs.Collection[V]) bool {
	return s.data.AddAll(other)
}

func (s *Stack[V]) AddIterator(iterator structs.Iterator[V]) bool {
	return s.data.AddIterator(iterator)
}

func (s *Stack[V]) Clear() {
	s.data.Clear()
}

func (s *Stack[V]) Contains(value V) bool {
	return s.data.Contains(value)
}

func (s *Stack[V]) ContainsAll(other structs.Collection[V]) bool {
	return s.data.ContainsAll(other)
}

func (s *Stack[V]) Element() V {
	return s.data.Get(0)
}

func (s *Stack[V]) IsEmpty() bool {
	return s.data.IsEmpty()
}

type Iterator[V any] struct {
	iter structs.Iterator[V]
}

func (i *Iterator[V]) HasNext() bool {
	return i.iter.HasNext()
}

func (i *Iterator[V]) Next() V {
	return i.iter.Next()
}

func (i *Iterator[V]) Remove() {
	i.iter.Remove()
}

func (s *Stack[V]) Iterator() structs.Iterator[V] {
	return &Iterator[V]{
		iter: s.data.Iterator(),
	}
}

func (s *Stack[V]) Peek() V {
	return s.data.Get(0)
}

func (s *Stack[V]) Poll() V {
	return s.data.RemoveAt(0)
}

func (s *Stack[V]) Remove(value V) bool {
	return s.data.Remove(value)
}

func (s *Stack[V]) RemoveAll(other structs.Collection[V]) bool {
	return s.data.RemoveAll(other)
}

func (s *Stack[V]) RemoveIterator(iter structs.Iterator[V]) bool {
	return s.data.RemoveIterator(iter)
}

func (s *Stack[V]) RemoveHead() V {
	return s.Poll()
}

func (s *Stack[V]) RetainAll(other structs.Collection[V]) bool {
	return s.data.RetainAll(other)
}

func (s *Stack[V]) Size() int {
	return s.data.Size()
}

func (s *Stack[V]) Values() []V {
	return s.data.Values()
}
