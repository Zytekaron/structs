package structs

type Iterator[V any] interface {
	HasNext() bool
	Next() V
	Remove()
}

type ListIterator[V any] interface {
	HasNext() bool
	HasPrevious() bool
	NextIndex() int
	Next() V
	PreviousIndex() int
	Previous() V
	Remove()
}

type Collection[V any] interface {
	Add(value V) bool
	AddAll(other Collection[V]) bool
	AddIterator(iter Iterator[V]) bool
	Clear()
	Contains(value V) bool
	ContainsAll(other Collection[V]) bool
	ContainsIterator(iter Iterator[V]) bool
	IsEmpty() bool
	Iterator() Iterator[V]
	Remove(value V) bool
	RemoveAll(other Collection[V]) bool
	RemoveIterator(iter Iterator[V]) bool
	RetainAll(other Collection[V]) bool
	Size() int
	Values() []V
}

type List[V any] interface {
	Collection[V]

	AddIndex(index int, value V)
	Get(index int) V
	IndexOf(value V) int
	LastIndexOf(value V) int
	RemoveIndex(index int) V
	Set(index int, value V) V
	Sort(cmp LessFunc[V])
	//SubList(from, to int) List[V]
}

//type SortedList[V comparable] interface {
//	List[V]
//
//	First() V
//	Last() V
//	HeadList(value V) SortedList[V]
//	TailList(value V) SortedList[V]
//}

type Set[V any] interface {
	Collection[V]
}

//type SortedSet[V any] interface {
//	Set[V]
//
//	First() V
//	Last() V
//	HeadSet(value V) SortedSet[V]
//	SubSet(from, to V) SortedSet[V]
//	TailSet(value V) SortedSet[V]
//}

type Queue[V any] interface {
	Collection[V]

	Element() V
	Peek() V
	Poll() V
	RemoveHead() V
}

type Deque[V any] interface {
	Collection[V]
	Queue[V]

	AddFirst(value V)
	AddLast(value V)
	GetFirst() V
	GetLast() V
	Offer(value V) bool
	OfferFirst(value V) bool
	OfferLast(value V) bool
	Peek() V
	PeekFirst() V
	PeekLast() V
	Poll() V
	PollFirst() V
	PollLast() V
	Pop() V
	Push(value V)
	RemoveHead() V
	RemoveFirst() V
	RemoveFirstOccurrence(value V) bool
	RemoveLast() V
	RemoveLastOccurrence(value V) bool
}
