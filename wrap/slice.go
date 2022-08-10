package wrap

import (
	"github.com/zytekaron/structs"
	"golang.org/x/exp/constraints"
)

// SliceWrap is a read-only wrapper around a slice,
// used in cases where you want to use it as a
// structs.Collection. Attempting to mutate the
// collection will result in a panic.
type SliceWrap[V any] struct {
	eq   structs.EqualFunc[V]
	data []V
}

func Slice[V any](eq structs.EqualFunc[V], s []V) *SliceWrap[V] {
	return &SliceWrap[V]{
		eq:   eq,
		data: s,
	}
}

func OrderedSlice[V constraints.Ordered](s []V) *SliceWrap[V] {
	return &SliceWrap[V]{
		eq:   structs.EqualOrdered[V],
		data: s,
	}
}

func Values[V any](eq structs.EqualFunc[V], values ...V) *SliceWrap[V] {
	return Slice(eq, values)
}

func OrderedValues[V constraints.Ordered](values ...V) *SliceWrap[V] {
	return OrderedSlice(values)
}

func (s *SliceWrap[V]) Add(value V) bool {
	panic("add called on immutable collection")
}

func (s *SliceWrap[V]) AddAll(other structs.Collection[V]) bool {
	panic("addall called on immutable collection")
}

func (s *SliceWrap[V]) AddIterator(iter structs.Iterator[V]) bool {
	panic("addall called on immutable collection")
}

func (s *SliceWrap[V]) Clear() {
	panic("clear called on immutable collection")
}

func (s *SliceWrap[V]) Contains(value V) bool {
	for _, elem := range s.data {
		if s.eq(elem, value) {
			return true
		}
	}
	return false
}

func (s *SliceWrap[V]) ContainsAll(other structs.Collection[V]) bool {
	it := other.Iterator()
	for it.HasNext() {
		value := it.Next()
		if !s.Contains(value) {
			return false
		}
	}
	return true
}

func (s *SliceWrap[V]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *SliceWrap[V]) Iterator() structs.Iterator[V] {
	return &SliceWrapIterator[V]{
		data: s.data,
	}
}

func (s *SliceWrap[V]) Remove(value V) bool {
	panic("remove called on immutable collection")
}

func (s *SliceWrap[V]) RemoveAll(other structs.Collection[V]) bool {
	panic("removeall called on immutable collection")
}

func (s *SliceWrap[V]) RemoveIterator(iter structs.Iterator[V]) bool {
	panic("removeiterator called on immutable collection")
}

func (s *SliceWrap[V]) RetainAll(other structs.Collection[V]) bool {
	panic("retainall called on immutable collection")
}

func (s *SliceWrap[V]) Size() int {
	return len(s.data)
}

func (s *SliceWrap[V]) Values() []V {
	// todo consider returning the slice itself
	out := make([]V, len(s.data))
	copy(out, s.data)
	return out
}
