package list

import (
	"fmt"
	"github.com/zytekaron/structs"
	"github.com/zytekaron/structs/wrap"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"strings"
)

// List is an implementation of a doubly-linked list.
//
// The value equality function may be omitted (nil)
// if no methods are called which use it.
type List[V any] struct {
	eq   structs.EqualFunc[V]
	head *listNode[V]
	tail *listNode[V]
	size int
}

// listNode is a linked list listNode.
type listNode[V any] struct {
	Value V
	Prev  *listNode[V]
	Next  *listNode[V]
}

func newNode[V any](value V) *listNode[V] {
	return &listNode[V]{Value: value}
}

func newFullNode[V any](value V, prev, next *listNode[V]) *listNode[V] {
	return &listNode[V]{
		Value: value,
		Prev:  prev,
		Next:  next,
	}
}

// New creates an empty List.
func New[V any](equal structs.EqualFunc[V]) *List[V] {
	return &List[V]{
		eq:   equal,
		head: nil,
		tail: nil,
		size: 0,
	}
}

// NewOrdered creates an empty List from a type that implements constraints.Ordered.
func NewOrdered[V constraints.Ordered]() *List[V] {
	return &List[V]{
		eq:   structs.EqualOrdered[V],
		head: nil,
		tail: nil,
		size: 0,
	}
}

// From creates a List from an existing collection.
func From[V any](equal structs.EqualFunc[V], other structs.Collection[V]) *List[V] {
	list := New[V](equal)
	list.AddAll(other)
	return list
}

// FromOrdered creates a List from an existing collection of a type that implements constraints.Ordered.
func FromOrdered[V constraints.Ordered](other structs.Collection[V]) *List[V] {
	list := NewOrdered[V]()
	list.AddAll(other)
	return list
}

// Of creates a List from an existing slice.
func Of[V any](equal structs.EqualFunc[V], values ...V) *List[V] {
	list := New[V](equal)
	list.AddIterator(wrap.SliceIterator[V](values))
	return list
}

// OfOrdered creates a List from an existing slice of a type that implements constraints.Ordered.
func OfOrdered[V constraints.Ordered](values ...V) *List[V] {
	list := NewOrdered[V]()
	list.AddIterator(wrap.SliceIterator[V](values))
	return list
}

// Add adds a value to the end of the list.
//
// Time Complexity: O(1)
func (l *List[V]) Add(value V) bool {
	return l.AddLast(value)
}

// AddAt adds a value at the specified index in the list.
//
// Panics if the index is out of bounds.
//
// Time Complexity: O(n)
func (l *List[V]) AddAt(index int, value V) {
	l.checkBounds(index, true)

	if index == 0 {
		l.AddFirst(value)
		return
	}
	if index == l.size {
		l.AddLast(value)
		return
	}

	target := l.getNodeAt(index - 1)
	l.insertAfterNode(target, value)
}

// AddAll adds all the values in the other collection to the end of the list.
//
// Time Complexity: O(n)
func (l *List[V]) AddAll(other structs.Collection[V]) bool {
	return l.AddIterator(other.Iterator())
}

// AddAllAt adds all the values in the other collection to the list,
// starting at the index provided and shifting existing elements right.
//
// Panics if the index is out of bounds.
//
// Time Complexity: O(n)
func (l *List[V]) AddAllAt(index int, other structs.Collection[V]) bool {
	return l.AddIteratorAt(index, other.Iterator())
}

// AddFirst adds a value to the front of the list.
//
// Time Complexity: O(1)
func (l *List[V]) AddFirst(value V) bool {
	l.addFirstNode(newNode(value))
	return true
}

// AddLast adds a value to the end of the list.
//
// Time Complexity: O(1)
func (l *List[V]) AddLast(value V) bool {
	l.addLastNode(newNode(value))
	return true
}

// AddIterator adds all the values in the iterator to the list.
//
// Time Complexity: O(n)
func (l *List[V]) AddIterator(iter structs.Iterator[V]) bool {
	changed := iter.HasNext()
	for iter.HasNext() {
		l.Add(iter.Next())
	}
	return changed
}

// AddIteratorAt adds all the values in the iterator to the list,
// starting at the index provided and shifting existing elements right.
//
// Panics if the index is out of bounds.
//
// Time Complexity: O(n)
func (l *List[V]) AddIteratorAt(index int, iter structs.Iterator[V]) bool {
	l.checkBounds(index, true)

	changed := iter.HasNext()
	node := l.getNodeAt(index)
	for iter.HasNext() {
		node = l.insertAfterNode(node, iter.Next())
	}
	return changed
}

// Clear clears the list.
//
// Time Complexity: O(1)
func (l *List[V]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Contains returns whether the value is present the list.
//
// Time Complexity: O(n)
func (l *List[V]) Contains(value V) bool {
	return l.IndexOf(value) >= 0
}

// ContainsAll returns whether all the values in the other collection are present the list.
//
// Time Complexity: O(nm)
//
//	n = size of the current list
//	m = size of the other collection
func (l *List[V]) ContainsAll(other structs.Collection[V]) bool {
	return l.ContainsIterator(other.Iterator())
}

// ContainsIterator returns whether all the values in the iterator are present the list.
//
// Time Complexity: O(nm)
//
//	n = size of the current list
//	m = number of values in the iterator
func (l *List[V]) ContainsIterator(iter structs.Iterator[V]) bool {
	for iter.HasNext() {
		if !l.Contains(iter.Next()) {
			return false
		}
	}
	return true
}

// Get returns the value at the specified index.
//
// Panics if the index is out of bounds.
//
// Time Complexity: O(n)
func (l *List[V]) Get(index int) V {
	l.checkBounds(index, false)

	return l.getNodeAt(index).Value
}

// GetFirst returns the value at the front of the list.
//
// Time Complexity: O(1)
func (l *List[V]) GetFirst() V {
	l.checkEmpty()

	return l.head.Value
}

// GetLast returns the value at the end of the list.
//
// Time Complexity: O(1)
func (l *List[V]) GetLast() V {
	l.checkEmpty()

	return l.tail.Value
}

// IndexOf returns the first index of a value in the list,
// or -1 if the value is not present in the list.
//
// Time Complexity: O(n)
func (l *List[V]) IndexOf(value V) int {
	index := 0
	node := l.head
	for node != nil {
		if l.eq(node.Value, value) {
			return index
		}
		index++
		node = node.Next
	}
	return -1
}

// LastIndexOf returns the last index of a value in the list,
// or -1 if the value is not present in the list.
//
// Time Complexity: O(n)
func (l *List[V]) LastIndexOf(value V) int {
	index := l.size - 1
	node := l.tail
	for node != nil {
		if l.eq(node.Value, value) {
			return index
		}
		index--
		node = node.Prev
	}
	return -1
}

// InsertAfter inserts a value after the first occurrence of an existing value in the list.
//
// Time Complexity: O(n)
func (l *List[V]) InsertAfter(before, value V) bool {
	target := l.findNode(before)
	l.insertAfterNode(target, value)
	return true
}

// InsertBefore inserts a value before the first occurrence of an existing value in the list.
//
// Time Complexity: O(n)
func (l *List[V]) InsertBefore(before, value V) bool {
	target := l.findNode(before)
	l.insertBeforeNode(target, value)
	return true
}

// IsEmpty returns whether this list is empty.
func (l *List[V]) IsEmpty() bool {
	return l.size == 0
}

// Iterator returns an Iterator for the list.
func (l *List[V]) Iterator() structs.Iterator[V] {
	return &Iterator[V]{
		list: l,
		next: l.head,
	}
}

// DescendingIterator returns a DescendingIterator for the list.
func (l *List[V]) DescendingIterator() structs.Iterator[V] {
	return &DescendingIterator[V]{
		iter: &Iterator[V]{
			list:  l,
			prev:  l.tail,
			index: l.size,
		},
	}
}

// Offer offers a value to the linked list, adding it to the end of the list.
//
// Time Complexity: O(1)
func (l *List[V]) Offer(value V) bool {
	return l.OfferLast(value)
}

// OfferFirst offers a value to the linked list, adding it to the front of the list.
//
// Time Complexity: O(1)
func (l *List[V]) OfferFirst(value V) bool {
	return l.AddFirst(value)
}

// OfferLast offers a value to the linked list, adding it to the end of the list.
//
// Time Complexity: O(1)
func (l *List[V]) OfferLast(value V) bool {
	return l.AddLast(value)
}

// Peek returns the first value in the list, or
// the zero value of the type if the list is empty.
//
// Time Complexity: O(1)
func (l *List[V]) Peek() V {
	return l.PeekFirst()
}

// PeekFirst returns the first value in the list, or
// the zero value of the type if the list is empty.
//
// Time Complexity: O(1)
func (l *List[V]) PeekFirst() V {
	if l.IsEmpty() {
		var null V
		return null
	}
	return l.head.Value
}

// PeekLast returns the last value in the list, or
// the zero value of the type if the list is empty.
//
// Time Complexity: O(1)
func (l *List[V]) PeekLast() V {
	if l.IsEmpty() {
		var null V
		return null
	}
	return l.tail.Value
}

// Poll removes and returns the first value in the list,
// or returns the zero value of the type if the list is empty.
//
// Time Complexity: O(1)
func (l *List[V]) Poll() V {
	return l.PollFirst()
}

// PollFirst removes and returns the first value in the list,
// or returns the zero value of the type if the list is empty.
//
// Time Complexity: O(1)
func (l *List[V]) PollFirst() V {
	if l.IsEmpty() {
		var null V
		return null
	}
	return l.removeNode(l.head)
}

// PollLast removes and returns the last value in the list,
// or returns the zero value of the type if the list is empty.
//
// Time Complexity: O(1)
func (l *List[V]) PollLast() V {
	if l.IsEmpty() {
		var null V
		return null
	}
	return l.removeNode(l.tail)
}

// Push appends a value to the front of the list.
//
// Time Complexity: O(1)
func (l *List[V]) Push(value V) {
	l.AddFirst(value)
}

// Pop removes and returns the first value in the list.
//
// Panics if the list is empty.
//
// Time Complexity: O(1)
func (l *List[V]) Pop() V {
	return l.RemoveFirst()
}

// RemoveHead removes and returns the first value in the list.
//
// Panics if the list is empty.
//
// Time Complexity: O(1)
func (l *List[V]) RemoveHead() V {
	return l.RemoveFirst()
}

// RemoveFirst removes and returns the first value in the list.
//
// Panics if the list is empty.
//
// Time Complexity: O(1)
func (l *List[V]) RemoveFirst() V {
	l.checkEmpty()

	return l.removeNode(l.head)
}

// RemoveFirstOccurrence removes the first occurrence of a value
// from the list and returns whether the value was present.
//
// Time Complexity: O(n)
func (l *List[V]) RemoveFirstOccurrence(value V) bool {
	node := l.findNode(value)
	if node == nil {
		return false
	}
	l.removeNode(node)
	return true
}

// RemoveLast removes and returns the last value in the list.
//
// Panics if the list is empty.
//
// Time Complexity: O(1)
func (l *List[V]) RemoveLast() V {
	l.checkEmpty()

	return l.removeNode(l.tail)
}

// RemoveLastOccurrence removes the last occurrence of a value
// from the list and returns whether the value was present.
//
// Time Complexity: O(n)
func (l *List[V]) RemoveLastOccurrence(value V) bool {
	node := l.findLastNode(value)
	if node == nil {
		return false
	}
	l.removeNode(node)
	return true
}

// Remove removes a value from the list and returns whether the value was present.
//
// Time Complexity: O(n)
func (l *List[V]) Remove(value V) bool {
	node := l.findNode(value)
	if node == nil {
		return false
	}
	l.removeNode(node)
	return true
}

// RemoveAt removes the value at the specified index from the list and returns its value.
//
// Panics if the index is out of bounds.
//
// Time Complexity: O(n)
func (l *List[V]) RemoveAt(index int) V {
	l.checkBounds(index, false)

	return l.removeNodeAt(index).Value
}

// RemoveAll removes all the values in the other collection from the list.
//
// Time Complexity: O(nm)
//
//	n = size of the current list
//	m = size of the other collection
func (l *List[V]) RemoveAll(other structs.Collection[V]) bool {
	return l.RemoveIterator(other.Iterator())
}

// RemoveIterator removes all the values in the iterator from the list.
//
// Time Complexity: O(nm)
//
//	n = size of the current list
//	m = number of values in the iterator
func (l *List[V]) RemoveIterator(iter structs.Iterator[V]) bool {
	changed := false
	for iter.HasNext() {
		value := iter.Next()
		if l.Remove(value) {
			changed = true
		}
	}
	return changed
}

// RetainAll removes all the values not present in the other collection from the list.
//
// Time Complexity: O(nm)
//
//	n = size of the current list
//	m = time complexity of Contains on the other collection
func (l *List[V]) RetainAll(other structs.Collection[V]) bool {
	changed := false
	node := l.head
	for node != nil {
		if !other.Contains(node.Value) {
			l.removeNode(node)
			changed = true
		}
		node = node.Next
	}
	return changed
}

// Reverse reverses the list.
//
// Time Complexity: O(n)
func (l *List[V]) Reverse() {
	if l.head == l.tail {
		return
	}
	l.head, l.tail = l.tail, l.head
	node := l.head
	for node != nil {
		prev := node.Prev
		node.Prev = node.Next
		node.Next = prev
		node = prev
	}
}

// Set sets the value at the specified index, returning the old value.
//
// Panics if the index is out of bounds.
//
// Time Complexity: O(n)
func (l *List[V]) Set(index int, value V) V {
	l.checkBounds(index, false)

	node := l.getNodeAt(index)
	oldValue := node.Value
	node.Value = value
	return oldValue
}

// Sort sorts the list based on a comparator.
//
// This method extracts an array using Values and reconstructs
// the entire list from scratch. If the list cannot be fully
// represented as an array, this method won't work.
//
// Time Complexity: best case O(nlogn), worst case O(n^2).
// See slices.SortFunc (algorithm: pattern-defeating quicksort).
func (l *List[V]) Sort(cmp structs.LessFunc[V]) {
	sorted := l.Values()
	slices.SortFunc(sorted, cmp)

	l.Clear()
	for _, value := range sorted {
		l.Add(value)
	}
}

// Size returns the length of the list.
//
// Time Complexity: O(1)
func (l *List[V]) Size() int {
	return l.size
}

// Values return a slice of the values in the list, starting at the front of the list.
//
// Time Complexity: O(n)
//
// Space Complexity: O(n)
func (l *List[V]) Values() []V {
	values := make([]V, 0, l.size)
	node := l.head
	for node != nil {
		values = append(values, node.Value)
		node = node.Next
	}
	return values
}

// ValuesReverse return a slice of the values in the list, starting at the end of the list.
//
// Time Complexity: O(n)
//
// Space Complexity: O(n)
func (l *List[V]) ValuesReverse() []V {
	values := make([]V, 0, l.size)
	node := l.tail
	for node != nil {
		values = append(values, node.Value)
		node = node.Prev
	}
	return values
}

// Clone creates a new list with the same values, starting at the front of the list.
//
// Time Complexity: O(n)
//
// Space Complexity: O(n)
func (l *List[V]) Clone() *List[V] {
	list := New[V](l.eq)
	node := l.head
	for node != nil {
		list.Add(node.Value)
		node = node.Next
	}
	return list
}

// CloneReverse creates a new list with the same values, starting at the end of the list.
//
// Time Complexity: O(n)
//
// Space Complexity: O(n)
func (l *List[V]) CloneReverse() *List[V] {
	list := New[V](l.eq)
	node := l.tail
	for node != nil {
		list.Add(node.Value)
		node = node.Prev
	}
	return list
}

// Each calls a function for each value in the list, starting at the front of the list.
func (l *List[V]) Each(f func(value V)) {
	node := l.head
	for node != nil {
		f(node.Value)
		node = node.Next
	}
}

// EachReverse calls a function for each value in the list, starting at the end of the list.
func (l *List[V]) EachReverse(f func(value V)) {
	node := l.tail
	for node != nil {
		f(node.Value)
		node = node.Prev
	}
}

// String returns a string representation of the list, with brackets and comma separated.
func (l *List[V]) String() string {
	if l.head == nil {
		return "[]"
	}

	var buf strings.Builder
	buf.WriteRune('[')
	buf.WriteString(fmt.Sprint(l.head.Value))

	node := l.head
	for node.Next != nil {
		node = node.Next
		buf.WriteString(", ")
		buf.WriteString(fmt.Sprint(node.Value))
	}

	buf.WriteRune(']')
	return buf.String()
}

// addFirstNode adds a listNode to the front of the list.
//
// Time Complexity: O(1)
func (l *List[V]) addFirstNode(node *listNode[V]) {
	node.Prev = nil
	if l.size == 0 {
		node.Next = nil
		l.head = node
		l.tail = node
	} else {
		node.Next = l.head
		l.head.Prev = node
		l.head = node
	}
	l.size++
}

// addLastNode adds a listNode to the end of the list.
//
// Time Complexity: O(1)
func (l *List[V]) addLastNode(node *listNode[V]) {
	node.Next = nil
	if l.size == 0 {
		node.Prev = nil
		l.head = node
		l.tail = node
	} else {
		node.Prev = l.tail
		l.tail.Next = node
		l.tail = node
	}
	l.size++
}

// insertAfterNode inserts a listNode after an existing listNode in the list.
//
// Time Complexity: O(1)
func (l *List[V]) insertAfterNode(target *listNode[V], value V) *listNode[V] {
	node := newFullNode(value, target, target.Next)
	if target.Next != nil {
		target.Next.Prev = node
	}
	target.Next = node
	if target == l.tail {
		l.tail = node
	}
	l.size++
	return node
}

// insertBeforeNode inserts a listNode before an existing listNode in the list.
//
// Time Complexity: O(1)
func (l *List[V]) insertBeforeNode(target *listNode[V], value V) *listNode[V] {
	node := newFullNode(value, target, target.Prev)
	if target.Prev != nil {
		target.Prev.Next = node
	}
	target.Prev = node
	if target == l.head {
		l.head = node
	}
	l.size++
	return node
}

// getNodeAt returns the listNode at the specified index.
//
// bounds checking should be performed prior to calling.
//
// Time Complexity: O(n)
func (l *List[V]) getNodeAt(index int) *listNode[V] {
	if index < 0 || index >= l.size {
		panic("index out of range")
	}
	node := l.head
	for index > 0 {
		node = node.Next
		index--
	}
	return node
}

// indexOfNode returns the index of a listNode in the list, otherwise -1.
//
// Time Complexity: O(n)
func (l *List[V]) indexOfNode(node *listNode[V]) int {
	index := 0
	tmp := l.head
	for tmp != nil {
		if tmp == node {
			return index
		}
		index++
		tmp = tmp.Next
	}
	return -1
}

// findNode attempts to find and return the first listNode with the specified value.
//
// Time Complexity: O(n)
func (l *List[V]) findNode(value V) *listNode[V] {
	node := l.head
	for node != nil {
		if l.eq(node.Value, value) {
			return node
		}
		node = node.Next
	}
	return nil
}

// findLastNode attempts to find and return the last listNode with the specified value.
//
// Time Complexity: O(n)
func (l *List[V]) findLastNode(value V) *listNode[V] {
	node := l.tail
	for node != nil {
		if l.eq(node.Value, value) {
			return node
		}
		node = node.Prev
	}
	return nil
}

// removeNode removes a listNode from the list.
//
// Time Complexity: O(1)
func (l *List[V]) removeNode(node *listNode[V]) V {
	l.size--
	if l.size == 0 {
		l.head = nil
		l.tail = nil
		return node.Value
	}
	if l.head == node {
		l.head = l.head.Next
		node.Next.Prev = nil
		return node.Value
	}
	if l.tail == node {
		l.tail = l.tail.Prev
		node.Prev.Next = nil
		return node.Value
	}
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
	return node.Value
}

// removeNodeAt removes the listNode at the specified index from the list and returns it.
//
// bounds checking should be performed prior to calling.
//
// Time Complexity: O(n)
func (l *List[V]) removeNodeAt(index int) *listNode[V] {
	node := l.getNodeAt(index)
	l.removeNode(node)
	return node
}

// checkEmpty panics if the list is empty. used to display useful
// errors in methods which rely on the head/tail being non-nil.
func (l *List[V]) checkEmpty() {
	if l.IsEmpty() {
		panic("list is empty")
	}
}

// checkBounds checks whether an index is within the checkBounds of the list.
//
// if allowEnd is true, index == l.size is allowed. this is useful
// for operations which may insert after the list's tail.
func (l *List[V]) checkBounds(index int, allowEnd bool) {
	if index < 0 || index > l.size || (!allowEnd && index == l.size) {
		panic("index out of bounds")
	}
}
