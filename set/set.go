package set

import (
	"github.com/zytekaron/structs"
	"github.com/zytekaron/structs/wrap"
)

// MapSet is an implementation of a set which uses
// Go's map type internally. It requires a type which
// is comparable; equality functions cannot be used.
type MapSet[V comparable] struct {
	data map[V]struct{}
	size int
}

func NewMapSet[V comparable]() *MapSet[V] {
	return &MapSet[V]{
		data: make(map[V]struct{}),
		size: 0,
	}
}

func From[V comparable](set map[V]struct{}) *MapSet[V] {
	size := 0
	for range set {
		size++
	}
	return &MapSet[V]{
		data: set,
		size: size,
	}
}

func (s *MapSet[V]) Add(value V) bool {
	_, ok := s.data[value]
	if !ok {
		s.data[value] = struct{}{}
	}
	return ok
}

func (s *MapSet[V]) AddAll(other structs.Collection[V]) bool {
	return s.AddIterator(other.Iterator())
}

func (s *MapSet[V]) AddIterator(iter structs.Iterator[V]) bool {
	changed := iter.HasNext()
	for iter.HasNext() {
		s.Add(iter.Next())
	}
	return changed
}

func (s *MapSet[V]) Clear() {
	s.data = make(map[V]struct{})
	s.size = 0
}

func (s *MapSet[V]) Contains(value V) bool {
	_, ok := s.data[value]
	return ok
}

func (s *MapSet[V]) ContainsAll(other structs.Collection[V]) bool {
	it := other.Iterator()
	for it.HasNext() {
		if !s.Contains(it.Next()) {
			return false
		}
	}
	return true
}

func (s *MapSet[V]) IsEmpty() bool {
	return s.size == 0
}

func (s *MapSet[V]) Iterator() structs.Iterator[V] {
	return wrap.MapSet(s.data)
}

func (s *MapSet[V]) Remove(value V) bool {
	_, ok := s.data[value]
	if ok {
		delete(s.data, value)
	}
	return ok
}

func (s *MapSet[V]) RemoveAll(other structs.Collection[V]) bool {
	changed := false
	for value := range s.data {
		if other.Contains(value) {
			delete(s.data, value)
			changed = true
		}
	}
	return changed
}

func (s *MapSet[V]) RetainAll(other structs.Collection[V]) bool {
	changed := false
	for value := range s.data {
		if !other.Contains(value) {
			delete(s.data, value)
			changed = true
		}
	}
	return changed
}

func (s *MapSet[V]) Size() int {
	return s.size
}

func (s *MapSet[V]) Values() []V {
	values := make([]V, 0, len(s.data))
	for val := range s.data {
		values = append(values, val)
	}
	return values
}
