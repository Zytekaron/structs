package wrap

import "reflect"

type MapSetIterator[V comparable] struct {
	data map[V]struct{}
	iter *reflect.MapIter
}

func MapSet[V comparable](m map[V]struct{}) *MapSetIterator[V] {
	return &MapSetIterator[V]{
		data: m,
		iter: reflect.ValueOf(m).MapRange(),
	}
}

func (i *MapSetIterator[V]) HasNext() bool {
	return i.iter.Next()
}

func (i *MapSetIterator[V]) Next() V {
	return i.key()
}

func (i *MapSetIterator[V]) Remove() {
	delete(i.data, i.key())
}

func (i *MapSetIterator[V]) key() V {
	return i.iter.Key().Interface().(V)
}
