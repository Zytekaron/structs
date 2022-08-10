package wrap

type SliceWrapIterator[V any] struct {
	data  []V
	index int
}

func SliceIterator[V any](s []V) *SliceWrapIterator[V] {
	return &SliceWrapIterator[V]{
		data: s,
	}
}

func ValueIterator[V any](values ...V) *SliceWrapIterator[V] {
	return SliceIterator(values)
}

func (s *SliceWrapIterator[V]) HasNext() bool {
	return s.index < len(s.data)
}

func (s *SliceWrapIterator[V]) Next() V {
	value := s.data[s.index]
	s.index++
	return value
}

func (s *SliceWrapIterator[V]) Remove() {
	panic("remove called on slice wrapping iterator")
}
