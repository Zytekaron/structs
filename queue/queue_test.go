package queue

import "testing"

func TestNew(t *testing.T) {
	const count = 10
	q := NewOrdered[int]()
	for i := 0; i < count; i++ {
		q.Enqueue(i)
	}

	expect := 0
	for !q.IsEmpty() {
		value := q.Dequeue()
		if value != expect {
			t.Errorf("expected %d but got %d", expect, value)
		}
		expect++
	}
	if expect != count {
		t.Errorf("expected %d elements but got %d", count, expect)
	}
}
