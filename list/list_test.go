package list

import (
	"fmt"
	"github.com/zytekaron/structs"
	"github.com/zytekaron/structs/wrap"
	"golang.org/x/exp/slices"
	"testing"
)

// Skipped tests:
// Size, Values
// Add, AddNode
// Offer, OfferFirst, OfferLast
// Peek, Poll, Push, Pop
// RemoveHead, RemoveFirst, RemoveLast
// Each, EachBack, EachNode, EachNodeBack

func TestNew(t *testing.T) {
	l := NewOrdered[int]()
	if l.Size() != 0 {
		t.Errorf("expected new linked list size to be 0, but got %d", l.Size())
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

func TestList_AddIndex(t *testing.T) {
	l := OfOrdered(1, 3, 5)
	l.AddIndex(1, 2) // 1, 2, 3, 5       // mid
	l.AddIndex(3, 4) // 1, 2, 3, 4, 5    // mid (shifted)
	l.AddIndex(5, 6) // 1, 2, 3, 4, 5, 6 // end (one after)

	expect := []int{1, 2, 3, 4, 5, 6}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_AddIndexNode(t *testing.T) {
	l := OfOrdered(1, 3, 5)
	a := newNode(2)
	b := newNode(4)
	c := newNode(6)
	l.AddIndexNode(1, a) // 1, 2, 3, 5       // mid
	l.AddIndexNode(3, b) // 1, 2, 3, 4, 5    // mid (shifted)
	l.AddIndexNode(5, c) // 1, 2, 3, 4, 5, 6 // end (one after)

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
	l.AddFirstNode(newNode(3))
	l.AddFirstNode(newNode(2))
	l.AddFirstNode(newNode(1))

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
	l.AddLastNode(newNode(1))
	l.AddLastNode(newNode(2))
	l.AddLastNode(newNode(3))

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
	l := OfOrdered(1, 2, 3)
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

func TestList_FindNode(t *testing.T) {
	l := NewOrdered[int]()
	a := newNode(1)
	b := newNode(2)
	c := newNode(1)
	l.AddNode(a)
	l.AddNode(b)
	l.AddNode(c)
	if l.FindNode(1) != a {
		t.Errorf("expected found node to be the first inserted node")
	}
	if l.FindNode(2) != b {
		t.Errorf("expected found node to be the middle inserted node")
	}
	if l.FindNode(3) != nil {
		t.Errorf("expected found node to be nil")
	}
}

func TestList_FindLastNode(t *testing.T) {
	l := NewOrdered[int]()
	a := newNode(1)
	b := newNode(2)
	c := newNode(1)
	l.AddNode(a)
	l.AddNode(b)
	l.AddNode(c)
	if l.FindLastNode(1) != c {
		t.Errorf("expected found node to be the last inserted node")
	}
	if l.FindLastNode(2) != b {
		t.Errorf("expected found node to be the middle inserted node")
	}
	if l.FindLastNode(3) != nil {
		t.Errorf("expected found node to be nil")
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

func TestList_GetFirstNode(t *testing.T) {
	l := NewOrdered[int]()
	a := newNode(1)
	b := newNode(2)
	l.AddNode(a)
	l.AddNode(b)
	if l.GetFirstNode() != a {
		t.Errorf("expected first node to be the first inserted node")
	}
}

func TestList_GetLast(t *testing.T) {
	l := OfOrdered(1, 2)
	if l.GetLast() != 2 {
		t.Errorf("expected last value to be 2")
	}
}

func TestList_GetLastNode(t *testing.T) {
	l := NewOrdered[int]()
	a := newNode(1)
	b := newNode(2)
	l.AddNode(a)
	l.AddNode(b)
	if l.GetLastNode() != b {
		t.Errorf("expected last node to be the last inserted node")
	}
}

func TestList_GetNode(t *testing.T) {
	l := NewOrdered[int]()
	a := newNode(1)
	b := newNode(2)
	c := newNode(3)
	l.AddNode(a)
	l.AddNode(b)
	l.AddNode(c)
	if l.GetNode(0) != a {
		t.Errorf("expected found node to be the first inserted node")
	}
	if l.GetNode(1) != b {
		t.Errorf("expected found node to be the second inserted node")
	}
	if l.GetNode(2) != c {
		t.Errorf("expected found node to be the third inserted node")
	}
}

func TestList_IndexOf(t *testing.T) {
	l := OfOrdered(1, 2, 1)
	index := l.IndexOf(1)
	if index != 0 {
		t.Errorf("expected first index of 1 to be 0, got %v", index)
	}
	index = l.IndexOf(2)
	if index != 1 {
		t.Errorf("expected first index of 2 to be 1, got %v", index)
	}
	index = l.IndexOf(3)
	if index != -1 {
		t.Errorf("expected first index of 3 to be -1, got %v", index)
	}
}

func TestList_LastIndexOf(t *testing.T) {
	l := OfOrdered(1, 2, 1)
	index := l.LastIndexOf(1)
	if index != 2 {
		t.Errorf("expected first index of 1 to be 2, got %v", index)
	}
	index = l.LastIndexOf(2)
	if index != 1 {
		t.Errorf("expected first index of 2 to be 1, got %v", index)
	}
	index = l.LastIndexOf(3)
	if index != -1 {
		t.Errorf("expected first index of 3 to be -1, got %v", index)
	}
}

func TestList_IndexOfNode(t *testing.T) {
	l := NewOrdered[int]()
	a := newNode(1)
	b := newNode(2)
	c := newNode(1)
	l.AddNode(a)
	l.AddNode(b)
	l.AddNode(c)
	d := newNode(3)

	index := l.IndexOfNode(a)
	if index != 0 {
		t.Errorf("expected index of the first inserted node to be 0, got %v", index)
	}
	index = l.IndexOfNode(b)
	if index != 1 {
		t.Errorf("expected index of the middle inserted node to be 1, got %v", index)
	}
	index = l.IndexOfNode(c)
	if index != 2 {
		t.Errorf("expected index of the last inserted node to be 2, got %v", index)
	}
	index = l.IndexOfNode(d)
	if index != -1 {
		t.Errorf("expected index of the uninserted node to be -1, got %v", index)
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

func TestList_InsertAfterNode(t *testing.T) {
	l := NewOrdered[int]()
	a := newNode(1)
	b := newNode(3)
	l.AddNode(a) // 1
	l.AddNode(b) // 1, 3

	l.InsertAfterNode(a, 2) // 1, 2, 3
	l.InsertAfterNode(b, 4) // 1, 2, 3, 4

	expect := []int{1, 2, 3, 4}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_InsertNodeAfterNode(t *testing.T) {
	l := NewOrdered[int]()
	a := newNode(1)
	b := newNode(2)
	c := newNode(3)
	d := newNode(4)
	l.AddNode(a) // 1
	l.AddNode(c) // 1, 3

	l.InsertNodeAfterNode(a, b) // 1, 2, 3
	l.InsertNodeAfterNode(c, d) // 1, 2, 3, 4

	expect := []int{1, 2, 3, 4}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_InsertBefore(t *testing.T) {
	l := OfOrdered(2, 4)
	l.InsertBefore(2, 1) // 1, 2, 4
	fmt.Println(l.Values())
	l.InsertBefore(4, 3) // 1, 2, 3, 4
	fmt.Println(l.Values())

	expect := []int{1, 2, 3, 4}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_InsertBeforeNode(t *testing.T) {
	l := NewOrdered[int]()
	a := newNode(2)
	b := newNode(4)
	l.AddNode(a) // 1
	l.AddNode(b) // 1, 3

	l.InsertBeforeNode(a, 1) // 1, 2, 4
	l.InsertBeforeNode(b, 3) // 1, 2, 3, 4

	expect := []int{1, 2, 3, 4}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_InsertNodeBeforeNode(t *testing.T) {
	l := NewOrdered[int]()
	a := newNode(1)
	b := newNode(2)
	c := newNode(3)
	d := newNode(4)
	l.AddNode(b) // 2
	l.AddNode(d) // 2, 4

	l.InsertNodeBeforeNode(b, a) // 1, 2, 4
	l.InsertNodeBeforeNode(d, c) // 1, 2, 3, 4

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

func TestList_RemoveFirst(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)
	l.RemoveFirst()
	l.RemoveFirst()
	l.RemoveFirst()

	expect := []int{4, 5}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_RemoveLast(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)
	l.RemoveLast()
	l.RemoveLast()
	l.RemoveLast()

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

func TestList_RemoveNode(t *testing.T) {
	l := NewOrdered[int]()
	a := newNode(1)
	b := newNode(2)
	c := newNode(3)
	d := newNode(4)
	e := newNode(5)
	l.AddNode(a)
	l.AddNode(b)
	l.AddNode(c)
	l.AddNode(d)
	l.AddNode(e)

	l.RemoveNode(a) // first
	l.RemoveNode(c) // middle
	l.RemoveNode(e) // last

	expect := []int{2, 4}
	got := l.Values()
	if !slices.Equal(got, expect) {
		t.Errorf("expected list %v but got %v", expect, got)
	}
}

func TestList_RemoveIndex(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)
	l.RemoveIndex(1) // 2; 1, 3, 4, 5
	l.RemoveIndex(1) // 2; 1, 4, 5    (shifted)
	l.RemoveIndex(2) // 2; 1, 4       (end, shifted)

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

func TestList_CloneBack(t *testing.T) {
	l := OfOrdered(1, 2, 3, 4, 5)
	c := l.CloneBack()

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
