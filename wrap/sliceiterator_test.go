package wrap

import "testing"

func TestSliceIterator(t *testing.T) {
	const count = 10
	s := make([]int, count)
	for i := 0; i < count; i++ {
		s[i] = i
	}

	sw := Slice(s)
	it := sw.Iterator()

	expect := 0
	for it.HasNext() {
		value := it.Next()
		if value != expect {
			t.Errorf("expected %d but got %d", expect, value)
		}
		expect++
	}
	if expect != count {
		t.Error("expected 5 elements to be read from the iterator, got only", expect)
	}
}
