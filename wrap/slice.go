package wrap

import (
	"github.com/zytekaron/structs"
	"golang.org/x/exp/constraints"
)

// SliceWrap is a read-only wrapper around a slice,
// used in cases where you want to use it in place
// of a structs.Collection. Attempting to mutate
// the collection will result in a panic.
//
// The value equality function may be omitted (nil)
// if no methods are called which use it.
type SliceWrap[V any] struct {
	eq   structs.EqualFunc[V]
	data []V
}

// Slice creates a SliceWrap of the passed slice.
func Slice[V any](eq structs.EqualFunc[V], s []V) *SliceWrap[V] {
	return &SliceWrap[V]{
		eq:   eq,
		data: s,
	}
}

// OrderedSlice creates a SliceWrap of the passed slice of a type that implements constraints.Ordered.
func OrderedSlice[V constraints.Ordered](s []V) *SliceWrap[V] {
	return &SliceWrap[V]{
		eq:   structs.EqualOrdered[V],
		data: s,
	}
}

// Values creates a SliceWrap of the passed values.
func Values[V any](eq structs.EqualFunc[V], values ...V) *SliceWrap[V] {
	return Slice(eq, values)
}

// OrderedValues creates a SliceWrap of the passed values of a type that implements constraints.Ordered.
func OrderedValues[V constraints.Ordered](values ...V) *SliceWrap[V] {
	return OrderedSlice(values)
}

// Add panics when called.
func (s *SliceWrap[V]) Add(V) bool {
	panic(structs.PanicUnsupportedOperation)
}

// AddAt panics when called.
func (s *SliceWrap[V]) AddAt(int, V) {
	panic(structs.PanicUnsupportedOperation)
}

// AddAll panics when called.
func (s *SliceWrap[V]) AddAll(structs.Collection[V]) bool {
	panic(structs.PanicUnsupportedOperation)
}

// AddIterator panics when called.
func (s *SliceWrap[V]) AddIterator(structs.Iterator[V]) bool {
	panic(structs.PanicUnsupportedOperation)
}

// Clear panics when called.
func (s *SliceWrap[V]) Clear() {
	panic(structs.PanicUnsupportedOperation)
}

// Contains returns whether the value is present the wrapped slice.
func (s *SliceWrap[V]) Contains(value V) bool {
	for _, elem := range s.data {
		if s.eq(elem, value) {
			return true
		}
	}
	return false
}

// ContainsAll returns whether all the values in the
// other collection are present the wrapped slice.
func (s *SliceWrap[V]) ContainsAll(other structs.Collection[V]) bool {
	return s.ContainsIterator(other.Iterator())
}

// ContainsIterator returns whether all the values
// in the iterator are present the wrapped slice.
func (s *SliceWrap[V]) ContainsIterator(iter structs.Iterator[V]) bool {
	for iter.HasNext() {
		if !s.Contains(iter.Next()) {
			return false
		}
	}
	return true
}

// Get returns the value at the specified index.
func (s *SliceWrap[V]) Get(index int) V {
	s.checkBounds(index)

	return s.data[index]
}

// IndexOf returns the first index of a value in the
// wrapped slice, or -1 if the value is not present.
func (s *SliceWrap[V]) IndexOf(value V) int {
	for i, val := range s.data {
		if s.eq(val, value) {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the last index of a value in the
// wrapped slice, or -1 if the value is not present.
func (s *SliceWrap[V]) LastIndexOf(value V) int {
	for i := s.Size() - 1; i >= 0; i-- {
		if s.eq(s.data[i], value) {
			return i
		}
	}
	return -1
}

// IsEmpty returns whether the wrapped slice is empty.
func (s *SliceWrap[V]) IsEmpty() bool {
	return s.Size() == 0
}

// Iterator returns a SliceWrapIterator over the wrapped slice.
func (s *SliceWrap[V]) Iterator() structs.Iterator[V] {
	return &SliceWrapIterator[V]{
		data: s.data,
	}
}

// Remove panics when called.
func (s *SliceWrap[V]) Remove(V) bool {
	panic(structs.PanicUnsupportedOperation)
}

// RemoveAt panics when called.
func (s *SliceWrap[V]) RemoveAt(int) V {
	panic(structs.PanicUnsupportedOperation)
}

// RemoveAll panics when called.
func (s *SliceWrap[V]) RemoveAll(structs.Collection[V]) bool {
	panic(structs.PanicUnsupportedOperation)
}

// RemoveIterator panics when called.
func (s *SliceWrap[V]) RemoveIterator(structs.Iterator[V]) bool {
	panic(structs.PanicUnsupportedOperation)
}

// RetainAll panics when callec.
func (s *SliceWrap[V]) RetainAll(structs.Collection[V]) bool {
	panic(structs.PanicUnsupportedOperation)
}

// Size returns the size of the wrapped slice.
func (s *SliceWrap[V]) Size() int {
	return len(s.data)
}

// Cap returns the capacity of the wrapped slice.
func (s *SliceWrap[V]) Cap() int {
	return cap(s.data)
}

// Set panics when called.
func (s *SliceWrap[V]) Set(int, V) V {
	panic(structs.PanicUnsupportedOperation)
}

// Sort panics when called.
func (s *SliceWrap[V]) Sort(structs.LessFunc[V]) {
	panic(structs.PanicUnsupportedOperation)
}

// Values returns a copy of the wrapped slice.
func (s *SliceWrap[V]) Values() []V {
	out := make([]V, len(s.data))
	copy(out, s.data)
	return out
}

func (s *SliceWrap[V]) checkBounds(index int) {
	if index < 0 || index >= s.Size() {
		panic(structs.PanicIndexOutOfBounds)
	}
}
