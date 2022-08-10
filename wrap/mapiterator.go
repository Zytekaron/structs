package wrap

import "reflect"

type MapIterator[K comparable, V any] struct {
	data map[K]V
	iter *reflect.MapIter
}

func Map[K comparable, V any](m map[K]V) *MapIterator[K, V] {
	return &MapIterator[K, V]{
		data: m,
		iter: reflect.ValueOf(m).MapRange(),
	}
}

func (i *MapIterator[K, V]) HasNext() bool {
	return i.iter.Next()
}

func (i *MapIterator[K, V]) Next() (K, V) {
	return i.key(), i.value()
}

func (i *MapIterator[K, V]) Remove() {
	delete(i.data, i.key())
}

func (i *MapIterator[K, V]) key() K {
	return i.iter.Key().Interface().(K)
}

func (i *MapIterator[K, V]) value() V {
	return i.iter.Value().Interface().(V)
}
