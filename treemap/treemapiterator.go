package treemap

type nodeIterator[K, V any] struct {
	treemap *TreeMap[K, V]
	next    *treeNode[K, V]
	last    *treeNode[K, V]
}

func (it *nodeIterator[K, V]) hasNext() bool {
	return it.next != nil
}

func (it *nodeIterator[K, V]) nextNode() *treeNode[K, V] {
	e := it.next
	if e == nil {
		panic("nextnode called on exhausted iterator")
	}

	it.next = e.successor()
	it.last = e
	return e
}

func (it *nodeIterator[K, V]) previousNode() *treeNode[K, V] {
	e := it.next
	if e == nil {
		panic("previousnode called on exhausted iterator")
	}

	it.next = e.predecessor()
	it.last = e
	return e
}

func (it *nodeIterator[K, V]) remove() {
	if it.last == nil {
		panic("remove called on unused iterator")
	}

	if it.last.Left != nil && it.last.Right != nil {
		it.next = it.last
		it.treemap.removeNode(it.last)
		it.last = nil
	}
}

type KeyIterator[K, V any] struct {
	iter *nodeIterator[K, V]
}

func (it *KeyIterator[K, V]) HasNext() bool {
	return it.iter.hasNext()
}

func (it *KeyIterator[K, V]) Next() K {
	return it.iter.nextNode().Key
}

func (it *KeyIterator[K, V]) Remove() {
	it.iter.remove()
}

type DescendingKeyIterator[K, V any] struct {
	iter *nodeIterator[K, V]
}

func (it *DescendingKeyIterator[K, V]) HasNext() bool {
	return it.iter.hasNext()
}

func (it *DescendingKeyIterator[K, V]) Next() K {
	return it.iter.previousNode().Key
}

func (it *DescendingKeyIterator[K, V]) Remove() {
	if it.iter.last == nil {
		panic("remove called on unused iterator")
	}

	it.iter.treemap.removeNode(it.iter.last)
	it.iter.last = nil
}

type ValueIterator[K, V any] struct {
	iter *nodeIterator[K, V]
}

func (it *ValueIterator[K, V]) HasNext() bool {
	return it.iter.hasNext()
}

func (it *ValueIterator[K, V]) Next() K {
	return it.iter.nextNode().Key
}

func (it *ValueIterator[K, V]) Remove() {
	it.iter.remove()
}

type DescendingValueIterator[K, V any] struct {
	iter *nodeIterator[K, V]
}

func (it *DescendingValueIterator[K, V]) HasNext() bool {
	return it.iter.hasNext()
}

func (it *DescendingValueIterator[K, V]) Next() V {
	return it.iter.previousNode().Value
}

func (it *DescendingValueIterator[K, V]) Remove() {
	if it.iter.last == nil {
		panic("remove called on unused iterator")
	}

	it.iter.treemap.removeNode(it.iter.last)
	it.iter.last = nil
}

type EntryIterator[K, V any] struct {
	iter *nodeIterator[K, V]
}

func (it *EntryIterator[K, V]) HasNext() bool {
	return it.iter.hasNext()
}

func (it *EntryIterator[K, V]) Next() (K, V) {
	next := it.iter.nextNode()
	return next.Key, next.Value
}

func (it *EntryIterator[K, V]) Remove() {
	it.iter.remove()
}

type DescendingEntryIterator[K, V any] struct {
	iter *nodeIterator[K, V]
}

func (it *DescendingEntryIterator[K, V]) HasNext() bool {
	return it.iter.hasNext()
}

func (it *DescendingEntryIterator[K, V]) Next() (K, V) {
	prev := it.iter.previousNode()
	return prev.Key, prev.Value
}

func (it *DescendingEntryIterator[K, V]) Remove() {
	if it.iter.last == nil {
		panic("remove called on unused iterator")
	}

	it.iter.treemap.removeNode(it.iter.last)
	it.iter.last = nil
}
