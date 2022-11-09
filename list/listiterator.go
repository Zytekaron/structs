package list

type Iterator[V any] struct {
	list  *List[V]
	prev  *Node[V]
	next  *Node[V]
	last  *Node[V] // last value returned
	index int
}

func (it *Iterator[V]) HasNext() bool {
	return it.next != nil
}

func (it *Iterator[V]) HasPrevious() bool {
	return it.prev != nil
}

func (it *Iterator[V]) Next() V {
	if it.next == nil {
		panic("next called on exhausted iterator")
	}
	it.index++

	this := it.next // value to return
	it.last, it.prev = this, this
	it.next = this.Next
	return this.Value
}

func (it *Iterator[V]) NextIndex() int {
	return it.index
}

func (it *Iterator[V]) Previous() V {
	if it.prev == nil {
		panic("previous called on unused iterator")
	}
	it.index--

	this := it.prev // value to return
	it.last, it.next = this, this
	it.prev = this.Prev
	return this.Value
}

func (it *Iterator[V]) PreviousIndex() int {
	return it.index - 1
}

func (it *Iterator[V]) Remove() {
	if it.prev == nil {
		panic("remove called on unused iterator")
	}
	// if removing the previous value (by list order),
	// decrease the index due to a left shift of data
	if it.last == it.prev {
		it.index--
	}

	this := it.last
	it.next = this.Next
	it.prev = this.Prev

	it.list.RemoveNode(this)
	it.last = nil
}

func (it *Iterator[V]) Set(value V) {
	if it.last == nil {
		panic("set called on invalid iterator state")
	}
	it.last.Value = value
}

type DescendingIterator[V any] struct {
	iter *Iterator[V]
}

func (it *DescendingIterator[V]) HasNext() bool {
	return it.iter.HasPrevious()
}

func (it *DescendingIterator[V]) HasPrevious() bool {
	return it.iter.HasNext()
}

func (it *DescendingIterator[V]) Next() V {
	return it.iter.Previous()
}

func (it *DescendingIterator[V]) NextIndex() int {
	return it.iter.PreviousIndex()
}

func (it *DescendingIterator[V]) Previous() V {
	return it.iter.Next()
}

func (it *DescendingIterator[V]) PreviousIndex() int {
	return it.iter.NextIndex()
}

func (it *DescendingIterator[V]) Remove() {
	it.iter.Remove()
}

func (it *DescendingIterator[V]) Set(value V) {
	it.iter.Set(value)
}
