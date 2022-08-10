package list

type LinkedListIterator[V any] struct {
	list  *List[V]
	prev  *Node[V]
	next  *Node[V]
	last  *Node[V] // last value returned
	index int
}

func (i *LinkedListIterator[V]) NextIndex() int {
	return i.index
}

func (i *LinkedListIterator[V]) PreviousIndex() int {
	return i.index - 1
}

func (i *LinkedListIterator[V]) HasNext() bool {
	return i.next != nil
}

func (i *LinkedListIterator[V]) HasPrevious() bool {
	return i.prev != nil
}

func (i *LinkedListIterator[V]) Next() V {
	if i.next == nil {
		panic("next called on exhausted iterator")
	}
	i.index++

	this := i.next // value to return
	i.last, i.prev = this, this
	i.next = this.Next
	return this.Value
}

func (i *LinkedListIterator[V]) Previous() V {
	if i.prev == nil {
		panic("previous called on unused iterator")
	}
	i.index--

	this := i.prev // value to return
	i.last, i.next = this, this
	i.prev = this.Prev
	return this.Value
}

func (i *LinkedListIterator[V]) Remove() {
	if i.prev == nil {
		panic("remove called on unused iterator")
	}
	// if removing the previous value (by list order),
	// decrease the index due to a left shift of data
	if i.last == i.prev {
		i.index--
	}

	this := i.last
	i.next = this.Next
	i.prev = this.Prev

	i.list.RemoveNode(this)
	i.last = nil
}

type LinkedListDescendingIterator[V any] struct {
	iter *LinkedListIterator[V]
}

func (i *LinkedListDescendingIterator[V]) NextIndex() int {
	return i.iter.PreviousIndex()
}

func (i *LinkedListDescendingIterator[V]) PreviousIndex() int {
	return i.iter.NextIndex()
}

func (i *LinkedListDescendingIterator[V]) HasNext() bool {
	return i.iter.HasPrevious()
}

func (i *LinkedListDescendingIterator[V]) HasPrevious() bool {
	return i.iter.HasNext()
}

func (i *LinkedListDescendingIterator[V]) Next() V {
	return i.iter.Previous()
}

func (i *LinkedListDescendingIterator[V]) Previous() V {
	return i.iter.Next()
}

func (i *LinkedListDescendingIterator[V]) Remove() {
	i.iter.Remove()
}
