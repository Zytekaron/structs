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
type List[V any] struct {
	eq   structs.EqualFunc[V]
	head *Node[V]
	tail *Node[V]
	size int
}

// Node is a linked list node.
type Node[V any] struct {
	Value V
	Next  *Node[V]
	Prev  *Node[V]
}

func newNode[V any](value V) *Node[V] {
	return &Node[V]{Value: value}
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
func (l *List[V]) Add(value V) bool {
	return l.AddLast(value)
}

// AddNode adds a node to the end of the list.
func (l *List[V]) AddNode(node *Node[V]) {
	l.AddLastNode(node)
}

// AddIndex adds a value at the specified index in the list.
func (l *List[V]) AddIndex(index int, value V) {
	l.AddIndexNode(index, newNode(value))
}

// AddIndexNode adds a node at the specified index in the list.
func (l *List[V]) AddIndexNode(index int, node *Node[V]) {
	if index == 0 {
		l.AddFirstNode(node)
		return
	}
	target := l.GetNode(index - 1)
	l.InsertNodeAfterNode(target, node)
}

// AddFirst adds a value to the front of the list.
func (l *List[V]) AddFirst(value V) bool {
	l.AddFirstNode(newNode(value))
	return true
}

// AddFirstNode adds a node to the front of the list.
func (l *List[V]) AddFirstNode(node *Node[V]) {
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

// AddLast adds a value to the end of the list.
func (l *List[V]) AddLast(value V) bool {
	l.AddLastNode(newNode(value))
	return true
}

// AddLastNode adds a node to the end of the list.
func (l *List[V]) AddLastNode(node *Node[V]) {
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

// AddAll adds all the values in the other collection to the end of the list.
func (l *List[V]) AddAll(other structs.Collection[V]) bool {
	return l.AddIterator(other.Iterator())
}

// AddIterator adds all the values in the iterator to the list.
func (l *List[V]) AddIterator(iter structs.Iterator[V]) bool {
	changed := iter.HasNext()
	for iter.HasNext() {
		l.Add(iter.Next())
	}
	return changed
}

// Clear clears the list.
func (l *List[V]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// Contains returns whether the value is present the list.
func (l *List[V]) Contains(value V) bool {
	return l.IndexOf(value) >= 0
}

// ContainsAll returns whether all the values in the other collection are present the list.
func (l *List[V]) ContainsAll(other structs.Collection[V]) bool {
	it := other.Iterator()
	for it.HasNext() {
		if !l.Contains(it.Next()) {
			return false
		}
	}
	return true
}

// FindNode attempts to find and return the first node with the specified value.
func (l *List[V]) FindNode(value V) *Node[V] {
	node := l.head
	for node != nil {
		if l.eq(node.Value, value) {
			return node
		}
		node = node.Next
	}
	return nil
}

// FindLastNode attempts to find and return the last node with the specified value.
func (l *List[V]) FindLastNode(value V) *Node[V] {
	node := l.tail
	for node != nil {
		if l.eq(node.Value, value) {
			return node
		}
		node = node.Prev
	}
	return nil
}

// Get returns the value at the specified index.
func (l *List[V]) Get(index int) V {
	return l.GetNode(index).Value
}

// GetFirst returns the value at the front of the list.
func (l *List[V]) GetFirst() V {
	return l.head.Value
}

// GetFirstNode returns the node at the front of the list.
func (l *List[V]) GetFirstNode() *Node[V] {
	return l.head
}

// GetLast returns the value at the end of the list.
func (l *List[V]) GetLast() V {
	return l.tail.Value
}

// GetLastNode returns the node at the head of the list.
func (l *List[V]) GetLastNode() *Node[V] {
	return l.tail
}

// GetNode returns the node at the specified index.
func (l *List[V]) GetNode(index int) *Node[V] {
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

// IndexOf returns the first index of a value in the list, otherwise -1.
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

// LastIndexOf returns the last index of a value in the list, otherwise -1.
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

// IndexOfNode returns the index of a node in the list, otherwise -1.
func (l *List[V]) IndexOfNode(node *Node[V]) int {
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

// InsertAfter inserts a node after the first occurrence of an existing value in the list.
func (l *List[V]) InsertAfter(before, value V) bool {
	target := l.FindNode(before)
	return l.InsertNodeAfterNode(target, newNode(value))
}

// InsertAfterNode inserts a value after an existing node in the list.
func (l *List[V]) InsertAfterNode(target *Node[V], value V) bool {
	return l.InsertNodeAfterNode(target, newNode(value))
}

// InsertNodeAfterNode inserts a node after an existing node in the list.
func (l *List[V]) InsertNodeAfterNode(target *Node[V], node *Node[V]) bool {
	node.Prev = target
	node.Next = target.Next
	if target.Next != nil {
		target.Next.Prev = node
	}
	target.Next = node
	if target == l.tail {
		l.tail = node
	}
	l.size++
	return true
}

// InsertBefore inserts a node before the first occurrence of an existing value in the list.
func (l *List[V]) InsertBefore(before, value V) bool {
	target := l.FindNode(before)
	fmt.Println(target, value)
	return l.InsertNodeBeforeNode(target, newNode(value))
}

// InsertBeforeNode inserts a value before an existing node in the list.
func (l *List[V]) InsertBeforeNode(target *Node[V], value V) bool {
	return l.InsertNodeBeforeNode(target, newNode(value))
}

// InsertNodeBeforeNode inserts a node before an existing node in the list.
func (l *List[V]) InsertNodeBeforeNode(target *Node[V], node *Node[V]) bool {
	node.Next = target
	node.Prev = target.Prev
	if target.Prev != nil {
		target.Prev.Next = node
	}
	target.Prev = node
	if target == l.head {
		l.head = node
	}
	l.size++
	return true
}

// IsEmpty returns whether this list is empty.
func (l *List[V]) IsEmpty() bool {
	return l.size == 0
}

// Iterator returns a LinkedListIterator for the list.
func (l *List[V]) Iterator() structs.Iterator[V] {
	return &LinkedListIterator[V]{
		list: l,
		next: l.head,
	}
}

// DescendingIterator returns a LinkedListDescendingIterator for the list.
func (l *List[V]) DescendingIterator() structs.Iterator[V] {
	return &LinkedListDescendingIterator[V]{
		iter: &LinkedListIterator[V]{
			list:  l,
			prev:  l.tail,
			index: l.size,
		},
	}
}

// Offer offers a value to the linked list, adding it to the end of the list.
func (l *List[V]) Offer(value V) bool {
	return l.OfferLast(value)
}

// OfferFirst offers a value to the linked list, adding it to the front of the list.
func (l *List[V]) OfferFirst(value V) bool {
	return l.AddFirst(value)
}

// OfferLast offers a value to the linked list, adding it to the end of the list.
func (l *List[V]) OfferLast(value V) bool {
	return l.AddLast(value)
}

// Peek returns the first value in the list, or
// the zero value of the type if the list is empty.
func (l *List[V]) Peek() V {
	return l.PeekFirst()
}

// PeekFirst returns the first value in the list, or
// the zero value of the type if the list is empty.
func (l *List[V]) PeekFirst() V {
	if l.IsEmpty() {
		var null V
		return null
	}
	return l.head.Value
}

// PeekLast returns the last value in the list, or
// the zero value of the type if the list is empty.
func (l *List[V]) PeekLast() V {
	if l.IsEmpty() {
		var null V
		return null
	}
	return l.tail.Value
}

// Poll removes and returns the first value in the list,
// or returns the zero value of the type if the list is empty.
func (l *List[V]) Poll() V {
	return l.PollFirst()
}

// PollFirst removes and returns the first value in the list,
// or returns the zero value of the type if the list is empty.
func (l *List[V]) PollFirst() V {
	if l.IsEmpty() {
		var null V
		return null
	}
	return l.RemoveFirst()
}

// PollLast removes and returns the last value in the list,
// or returns the zero value of the type if the list is empty.
func (l *List[V]) PollLast() V {
	if l.IsEmpty() {
		var null V
		return null
	}
	return l.RemoveLast()
}

// Push appends a value to the front of the list.
func (l *List[V]) Push(value V) {
	l.AddFirst(value)
}

// Pop removes and returns the first value in the list.
func (l *List[V]) Pop() V {
	return l.PollFirst()
}

// RemoveHead removes and returns the first value in the list.
func (l *List[V]) RemoveHead() V {
	return l.RemoveFirst()
}

// RemoveFirst removes and returns the first value in the list.
func (l *List[V]) RemoveFirst() V {
	node := l.head
	l.RemoveNode(node)
	return node.Value
}

// RemoveFirstOccurrence removes the first occurrence of a value from the list.
func (l *List[V]) RemoveFirstOccurrence(value V) bool {
	node := l.FindNode(value)
	if node == nil {
		return false
	}
	l.RemoveNode(node)
	return true
}

// RemoveLast removes and returns the last value in the list.
func (l *List[V]) RemoveLast() V {
	node := l.tail
	l.RemoveNode(node)
	return node.Value
}

// RemoveLastOccurrence removes the last occurrence of a value from the list.
func (l *List[V]) RemoveLastOccurrence(value V) bool {
	node := l.FindLastNode(value)
	if node == nil {
		return false
	}
	l.RemoveNode(node)
	return true
}

// Remove removes a value from the list.
func (l *List[V]) Remove(value V) bool {
	node := l.FindNode(value)
	if node == nil {
		return false
	}
	l.RemoveNode(node)
	return true
}

// RemoveNode removes a node from the list.
func (l *List[V]) RemoveNode(node *Node[V]) {
	l.size--
	if l.size == 0 {
		l.head = nil
		l.tail = nil
		return
	}
	if l.head == node {
		l.head = l.head.Next
		node.Next.Prev = nil
		return
	}
	if l.tail == node {
		l.tail = l.tail.Prev
		node.Prev.Next = nil
		return
	}
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

// RemoveIndex removes the element at the specified index from the list and returns its value.
func (l *List[V]) RemoveIndex(index int) V {
	return l.RemoveIndexNode(index).Value
}

// RemoveIndexNode removes the node at the specified index from the list and returns it.
func (l *List[V]) RemoveIndexNode(index int) *Node[V] {
	node := l.GetNode(index)
	l.RemoveNode(node)
	return node
}

// RemoveAll removes all the values in the other collection from the list.
func (l *List[V]) RemoveAll(other structs.Collection[V]) bool {
	return l.RemoveIterator(other.Iterator())
}

// RemoveIterator removes all the values in the iterator from the list.
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
func (l *List[V]) RetainAll(other structs.Collection[V]) bool {
	changed := false
	node := l.head
	for node != nil {
		if !other.Contains(node.Value) {
			l.RemoveNode(node)
			changed = true
		}
		node = node.Next
	}
	return changed
}

// Reverse reverses elements in the list.
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
func (l *List[V]) Set(index int, value V) V {
	node := l.GetNode(index)
	oldValue := node.Value
	node.Value = value
	return oldValue
}

// Sort sorts the list based on a comparator.
func (l *List[V]) Sort(cmp structs.LessFunc[V]) {
	sorted := l.Values()
	slices.SortFunc(sorted, cmp)

	l.Clear()
	for _, value := range sorted {
		l.Add(value)
	}
}

// Size returns the number of elements in the list.
func (l *List[V]) Size() int {
	return l.size
}

func (l *List[V]) Values() []V {
	values := make([]V, 0, l.size)
	node := l.head
	for node != nil {
		values = append(values, node.Value)
		node = node.Next
	}
	return values
}

// Clone creates a new list with the same values, starting at the front.
func (l *List[V]) Clone() *List[V] {
	list := New[V](l.eq)
	node := l.head
	for node != nil {
		list.Add(node.Value)
		node = node.Next
	}
	return list
}

// CloneBack creates a new list with the same values, starting at the back.
func (l *List[V]) CloneBack() *List[V] {
	list := New[V](l.eq)
	node := l.tail
	for node != nil {
		list.Add(node.Value)
		node = node.Prev
	}
	return list
}

// Each calls a function for each value in the list, starting at the front.
func (l *List[V]) Each(f func(value V)) {
	node := l.head
	for node != nil {
		f(node.Value)
		node = node.Next
	}
}

// EachNode calls a function for each element in the list, starting at the front.
func (l *List[V]) EachNode(f func(node *Node[V])) {
	node := l.head
	for node != nil {
		f(node)
		node = node.Next
	}
}

// EachBack calls a function for each value in the list, starting at the end.
func (l *List[V]) EachBack(f func(value V)) {
	node := l.tail
	for node != nil {
		f(node.Value)
		node = node.Prev
	}
}

// EachNodeBack calls a function for each element in the list, starting at the end.
func (l *List[V]) EachNodeBack(f func(node *Node[V])) {
	node := l.tail
	for node != nil {
		f(node)
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
