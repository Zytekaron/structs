package treemap

type NodeIterator[K, V any] struct {
	treemap *TreeMap[K, V]
	next    *TreeNode[K, V]
	last    *TreeNode[K, V]
}

func (it *NodeIterator[K, V]) HasNext() bool {
	return it.next != nil
}

func (it *NodeIterator[K, V]) NextNode() *TreeNode[K, V] {
	e := it.next
	if e == nil {
		panic("nextnode called on exhausted iterator")
	}

	it.next = e.successor()
	it.last = e
	return e
}

func (it *NodeIterator[K, V]) PreviousNode() *TreeNode[K, V] {
	e := it.next
	if e == nil {
		panic("previousnode called on exhausted iterator")
	}

	it.next = e.predecessor()
	it.last = e
	return e
}

func (it *NodeIterator[K, V]) Remove() {
	if it.last == nil {
		panic("remove called on unused iterator")
	}

	if it.last.Left != nil && it.last.Right != nil {
		it.next = it.last
		it.treemap.RemoveNode(it.last)
		it.last = nil
	}
}

type KeyIterator[K, V any] struct {
	iter *NodeIterator[K, V]
}

func (it *KeyIterator[K, V]) HasNext() bool {
	return it.iter.HasNext()
}

func (it *KeyIterator[K, V]) Next() K {
	return it.iter.NextNode().Key
}

func (it *KeyIterator[K, V]) Remove() {
	it.iter.Remove()
}

type DescendingKeyIterator[K, V any] struct {
	iter *NodeIterator[K, V]
}

func (it *DescendingKeyIterator[K, V]) HasNext() bool {
	return it.iter.HasNext()
}

func (it *DescendingKeyIterator[K, V]) Next() K {
	return it.iter.PreviousNode().Key
}

func (it *DescendingKeyIterator[K, V]) Remove() {
	if it.iter.last == nil {
		panic("remove called on unused iterator")
	}

	it.iter.treemap.RemoveNode(it.iter.last)
	it.iter.last = nil
}

type ValueIterator[K, V any] struct {
	iter *NodeIterator[K, V]
}

func (it *ValueIterator[K, V]) HasNext() bool {
	return it.iter.HasNext()
}

func (it *ValueIterator[K, V]) Next() K {
	return it.iter.NextNode().Key
}

func (it *ValueIterator[K, V]) Remove() {
	it.iter.Remove()
}

type DescendingValueIterator[K, V any] struct {
	iter *NodeIterator[K, V]
}

func (it *DescendingValueIterator[K, V]) HasNext() bool {
	return it.iter.HasNext()
}

func (it *DescendingValueIterator[K, V]) Next() V {
	return it.iter.PreviousNode().Value
}

func (it *DescendingValueIterator[K, V]) Remove() {
	if it.iter.last == nil {
		panic("remove called on unused iterator")
	}

	it.iter.treemap.RemoveNode(it.iter.last)
	it.iter.last = nil
}

type EntryIterator[K, V any] struct {
	iter *NodeIterator[K, V]
}

func (it *EntryIterator[K, V]) HasNext() bool {
	return it.iter.HasNext()
}

func (it *EntryIterator[K, V]) Next() (K, V) {
	next := it.iter.NextNode()
	return next.Key, next.Value
}

func (it *EntryIterator[K, V]) Remove() {
	it.iter.Remove()
}

type DescendingEntryIterator[K, V any] struct {
	iter *NodeIterator[K, V]
}

func (it *DescendingEntryIterator[K, V]) HasNext() bool {
	return it.iter.HasNext()
}

func (it *DescendingEntryIterator[K, V]) Next() (K, V) {
	prev := it.iter.PreviousNode()
	return prev.Key, prev.Value
}

func (it *DescendingEntryIterator[K, V]) Remove() {
	if it.iter.last == nil {
		panic("remove called on unused iterator")
	}

	it.iter.treemap.RemoveNode(it.iter.last)
	it.iter.last = nil
}
