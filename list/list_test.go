package list

import (
	"github.com/zytekaron/structs"
	"github.com/zytekaron/structs/wrap"
	"golang.org/x/exp/slices"
	"testing"
)

func TestNew(t *testing.T) {
	var l structs.List[int] = NewOrdered[int]()
	if l.Size() != 0 {
		t.Errorf("expected new linked list size to be 0 but got %d", l.Size())
	}
}

func TestOf(t *testing.T) {
	l := OfOrdered(1, 2, 3)

	expect := []int{1, 2, 3}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_Add(t *testing.T) {
	l := NewOrdered[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	expect := []int{1, 2, 3}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_AddAt(t *testing.T) {
	l := OfOrdered(1, 3, 5)
	l.AddAt(1, 2) // 1, 2, 3, 5       // mid
	l.AddAt(3, 4) // 1, 2, 3, 4, 5    // mid (shifted)
	l.AddAt(5, 6) // 1, 2, 3, 4, 5, 6 // end (one after)

	expect := []int{1, 2, 3, 4, 5, 6}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_AddFirst(t *testing.T) {
	l := NewOrdered[int]()
	l.AddFirst(3)
	l.AddFirst(2)
	l.AddFirst(1)

	expect := []int{1, 2, 3}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_AddFirstNode(t *testing.T) {
	l := NewOrdered[int]()
	l.addFirstNode(newNode(3))
	l.addFirstNode(newNode(2))
	l.addFirstNode(newNode(1))

	expect := []int{1, 2, 3}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_AddLast(t *testing.T) {
	l := NewOrdered[int]()
	l.AddLast(1)
	l.AddLast(2)
	l.AddLast(3)

	expect := []int{1, 2, 3}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_AddLastNode(t *testing.T) {
	l := NewOrdered[int]()
	l.addLastNode(newNode(1))
	l.addLastNode(newNode(2))
	l.addLastNode(newNode(3))

	expect := []int{1, 2, 3}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_AddAll(t *testing.T) {
	l := NewOrdered[int]()
	l.AddAll(wrap.OrderedValues(1, 2, 3))
	l.AddAll(wrap.OrderedValues(4, 5, 6))

	expect := []int{1, 2, 3, 4, 5, 6}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_AddIterator(t *testing.T) {
	l := NewOrdered[int]()
	l.AddIterator(wrap.ValueIterator(1, 2))
	l.AddIterator(wrap.ValueIterator(3, 4, 5))

	expect := []int{1, 2, 3, 4, 5}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_Clear(t *testing.T) {
	l := OfOrdered(1, 2, 3)
	l.Clear()

	var expect []int
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
	if l.Size() != 0 {
		t.Errorf("expected list to be of size 0 but got %v", l.Size())
	}
}

func TestList_Contains(t *testing.T) {
	l := OfOrdered[int]()
	if l.Contains(0) {
		t.Errorf("expected list to not contain 0")
	}

	l = OfOrdered(1, 2, 3)
	if !l.Contains(1) {
		t.Errorf("expected list to contain 1")
	}
	if !l.Contains(2) {
		t.Errorf("expected list to contain 2")
	}
	if !l.Contains(3) {
		t.Errorf("expected list to contain 3")
	}
	if l.Contains(4) {
		t.Errorf("expected list to not contain 4")
	}
}

func TestList_ContainsAll(t *testing.T) {
	l := OfOrdered(1, 2, 3)
	if !l.ContainsAll(wrap.OrderedValues(1)) {
		t.Errorf("expected list to contain all of: 1")
	}
	if !l.ContainsAll(wrap.OrderedValues(2, 3)) {
		t.Errorf("expected list to contain all of: 2, 3")
	}
	if !l.ContainsAll(wrap.OrderedValues(1, 2, 3)) {
		t.Errorf("expected list to contain all of: 1, 2, 3")
	}
	if l.ContainsAll(wrap.OrderedValues(2, 4)) {
		t.Errorf("expected list to not contain all of: 2, 4")
	}
	if l.ContainsAll(wrap.OrderedValues(1, 2, 3, 4)) {
		t.Errorf("expected list to not contain all of: 1, 2, 3, 4")
	}
}

func TestList_Get(t *testing.T) {
	l := OfOrdered(1, 2, 3)
	if l.Get(0) != 1 {
		t.Errorf("expected value at index 0 to have the value 1")
	}
	if l.Get(1) != 2 {
		t.Errorf("expected value at index 1 to have the value 2")
	}
	if l.Get(2) != 3 {
		t.Errorf("expected value at index 2 to have the value 3")
	}
}

func TestList_GetFirst(t *testing.T) {
	l := OfOrdered(1, 2)
	if l.GetFirst() != 1 {
		t.Errorf("expected first value to be 1")
	}
}

func TestList_GetLast(t *testing.T) {
	l := OfOrdered(1, 2)
	if l.GetLast() != 2 {
		t.Errorf("expected last value to be 2")
	}
}

func TestList_IndexOf(t *testing.T) {
	l := OfOrdered(1, 2, 1)
	index := l.IndexOf(1)
	if index != 0 {
		t.Errorf("expected first index of 1 to be 0 but got %v", index)
	}
	index = l.IndexOf(2)
	if index != 1 {
		t.Errorf("expected first index of 2 to be 1 but got %v", index)
	}
	index = l.IndexOf(3)
	if index != -1 {
		t.Errorf("expected first index of 3 to be -1 but got %v", index)
	}
}

func TestList_LastIndexOf(t *testing.T) {
	l := OfOrdered(1, 2, 1)
	index := l.LastIndexOf(1)
	if index != 2 {
		t.Errorf("expected first index of 1 to be 2 but got %v", index)
	}
	index = l.LastIndexOf(2)
	if index != 1 {
		t.Errorf("expected first index of 2 to be 1 but got %v", index)
	}
	index = l.LastIndexOf(3)
	if index != -1 {
		t.Errorf("expected first index of 3 to be -1 but got %v", index)
	}
}

func TestList_InsertAfter(t *testing.T) {
	l := OfOrdered(1, 3)
	l.InsertAfter(1, 2) // 1, 2, 3
	l.InsertAfter(3, 4) // 1, 2, 3, 4

	expect := []int{1, 2, 3, 4}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_InsertBefore(t *testing.T) {
	l := OfOrdered(2, 4)
	l.InsertBefore(2, 1) // 1, 2, 4
	l.InsertBefore(4, 3) // 1, 2, 3, 4

	expect := []int{1, 2, 3, 4}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_IsEmpty(t *testing.T) {
	l := NewOrdered[int]()
	if !l.IsEmpty() {
		t.Errorf("expected new list to be empty, instead got size of %v", l.Size())
	}
	l.Add(1)
	if l.IsEmpty() {
		t.Errorf("expected new list to be non-empty")
	}
}

func TestList_Size(t *testing.T) {
	l := NewOrdered[int]()
	if l.Size() != 0 {
		t.Errorf("expected size to be 0 but got %v", l.Size())
	}
	l.Add(1)
	if l.Size() != 1 {
		t.Errorf("expected size to be 1 but got %v", l.Size())
	}
	l.Add(2)
	if l.Size() != 2 {
		t.Errorf("expected size to be 2 but got %v", l.Size())
	}
	l.Clear()
	if l.Size() != 0 {
		t.Errorf("expected size to be 0 after clear but got %v", l.Size())
	}
}

func TestList_Values(t *testing.T) {
	l := NewOrdered[int]()

	var expect []int
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}

	l.Add(1)
	expect = []int{1}
	got = l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}

	l.Add(2)
	expect = []int{1, 2}
	got = l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_Offer(t *testing.T) {
	l := NewOrdered[int]()
	l.Offer(1)
	l.Offer(2)
	l.Offer(3)

	expect := []int{1, 2, 3}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_OfferFirst(t *testing.T) {
	l := NewOrdered[int]()
	l.OfferFirst(3)
	l.OfferFirst(2)
	l.OfferFirst(1)

	expect := []int{1, 2, 3}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_OfferLast(t *testing.T) {
	l := NewOrdered[int]()
	l.OfferLast(1)
	l.OfferLast(2)
	l.OfferLast(3)

	expect := []int{1, 2, 3}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_Peek(t *testing.T) {
	l := OfOrdered(1, 2)
	if l.Peek() != 1 {
		t.Errorf("expected first value to be 1")
	}

	l = NewOrdered[int]()
	if l.Peek() != 0 {
		t.Errorf("expected peek for empty list to be 0")
	}
}

func TestList_PeekFirst(t *testing.T) {
	l := OfOrdered(1, 2)
	if l.PeekFirst() != 1 {
		t.Errorf("expected first value to be 1")
	}

	l = NewOrdered[int]()
	if l.PeekFirst() != 0 {
		t.Errorf("expected peek for empty list to be 0")
	}
}

func TestList_PeekLast(t *testing.T) {
	l := OfOrdered(1, 2)
	if l.PeekLast() != 2 {
		t.Errorf("expected last value to be 2")
	}

	l = NewOrdered[int]()
	if l.PeekLast() != 0 {
		t.Errorf("expected peek for empty list to be 0")
	}
}

func TestList_Poll(t *testing.T) {
	l := OfOrdered(1, 2)
	if l.Poll() != 1 {
		t.Errorf("expected first value to be 1")
	}
}

func TestList_PollFirst(t *testing.T) {
	l := OfOrdered(1, 2)
	if l.PollFirst() != 1 {
		t.Errorf("expected first value to be 1")
	}
}

func TestList_PollLast(t *testing.T) {
	l := OfOrdered(1, 2)
	if l.PollLast() != 2 {
		t.Errorf("expected last value to be 2")
	}
}

func TestList_Push(t *testing.T) {
	l := NewOrdered[int]()
	l.Push(3)
	l.Push(2)
	l.Push(1)

	expect := []int{1, 2, 3}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_Pop(t *testing.T) {
	l := OfOrdered(1, 2, 3)
	value := l.Pop()
	if value != 1 {
		t.Errorf("expected first popped value to be 1 but got %v", value)
	}
	value = l.Pop()
	if value != 2 {
		t.Errorf("expected first popped value to be 2 but got %v", value)
	}
	value = l.Pop()
	if value != 3 {
		t.Errorf("expected first popped value to be 3 but got %v", value)
	}

	if !l.IsEmpty() {
		t.Errorf("expected popped out list to be empty, ")
	}
}

func TestList_RemoveHead(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)

	value := l.RemoveHead()
	if value != 1 {
		t.Errorf("expected first removed value to be 1 but got %v", value)
	}
	value = l.RemoveHead()
	if value != 2 {
		t.Errorf("expected second removed value to be 2 but got %v", value)
	}
	value = l.RemoveHead()
	if value != 3 {
		t.Errorf("expected third removed value to be 3 but got %v", value)
	}

	expect := []int{4, 5}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_RemoveFirst(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)

	value := l.RemoveFirst()
	if value != 1 {
		t.Errorf("expected first removed value to be 1 but got %v", value)
	}
	value = l.RemoveFirst()
	if value != 2 {
		t.Errorf("expected second removed value to be 2 but got %v", value)
	}
	value = l.RemoveFirst()
	if value != 3 {
		t.Errorf("expected third removed value to be 3 but got %v", value)
	}

	expect := []int{4, 5}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_RemoveLast(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)

	value := l.RemoveLast()
	if value != 5 {
		t.Errorf("expected first removed value to be 5 but got %v", value)
	}
	value = l.RemoveLast()
	if value != 4 {
		t.Errorf("expected second removed value to be 4 but got %v", value)
	}
	value = l.RemoveLast()
	if value != 3 {
		t.Errorf("expected third removed value to be 3 but got %v", value)
	}

	expect := []int{1, 2}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_RemoveFirstOccurrence(t *testing.T) {
	l := OfOrdered(1, 2, 3, 2, 1)
	l.RemoveFirstOccurrence(1)
	l.RemoveFirstOccurrence(2)

	expect := []int{3, 2, 1}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_RemoveLastOccurrence(t *testing.T) {
	l := OfOrdered(1, 2, 3, 2, 1)
	l.RemoveLastOccurrence(1)
	l.RemoveLastOccurrence(2)

	expect := []int{1, 2, 3}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_Remove(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)
	l.Remove(1) // head
	l.Remove(3) // center
	l.Remove(5) // tail

	expect := []int{2, 4}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_RemoveAt(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)
	l.RemoveAt(1) // 2; 1, 3, 4, 5
	l.RemoveAt(1) // 2; 1, 4, 5    (shifted)
	l.RemoveAt(2) // 2; 1, 4       (end, shifted)

	expect := []int{1, 4}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_RemoveAll(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)
	l.RemoveAll(wrap.OrderedValues(2, 3))
	l.RemoveAll(wrap.OrderedValues(4, 1))

	expect := []int{5}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_RemoveIterator(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)
	l.RemoveIterator(wrap.ValueIterator(2, 3))
	l.RemoveIterator(wrap.ValueIterator(4, 1))

	expect := []int{5}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_RetainAll(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)
	l.RetainAll(wrap.OrderedValues(2, 3, 4, 5))
	l.RetainAll(wrap.OrderedValues(2, 4, 6))

	expect := []int{2, 4}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_Set(t *testing.T) {
	l := OfOrdered(0, 2, 0, 4, 0)
	l.Set(0, 1)
	l.Set(2, 3)
	l.Set(4, 5)

	expect := []int{1, 2, 3, 4, 5}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_Sort(t *testing.T) {
	l := OfOrdered(2, 4, 3, 1, 5)

	l.Sort(structs.LessOrdered[int])
	expect := []int{1, 2, 3, 4, 5}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}

	l.Sort(structs.GreaterOrdered[int])
	expect = []int{5, 4, 3, 2, 1}
	got = l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_Clone(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)
	c := l.Clone()

	expect := []int{1, 2, 3, 4, 5}
	got := c.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_CloneReverse(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)
	c := l.CloneReverse()

	expect := []int{5, 4, 3, 2, 1}
	got := c.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_String(t *testing.T) {
	l := NewOrdered[int]()
	expect := "[]"
	got := l.String()
	if got != expect {
		t.Errorf("expected string '%v' but got '%v'", expect, got)
	}

	l = OfOrdered(1)
	expect = "[1]"
	got = l.String()
	if got != expect {
		t.Errorf("expected string '%v' but got '%v'", expect, got)
	}

	l = OfOrdered(1, 2, 3, 4, 5)
	expect = "[1, 2, 3, 4, 5]"
	got = l.String()
	if got != expect {
		t.Errorf("expected string '%v' but got '%v'", expect, got)
	}
}
