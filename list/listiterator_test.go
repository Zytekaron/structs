package list

import (
	"golang.org/x/exp/slices"
	"testing"
)

func TestListIterator(t *testing.T) {
	initial := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// filter: keep evens
	evens := OfOrdered(initial...)
	it := evens.Iterator()
	for it.HasNext() {
		num := it.Next()
		if num%2 == 1 {
			it.Remove()
		}
	}
	evensSlice := evens.Values()
	expectEvens := []int{0, 2, 4, 6, 8}
	if !slices.Equal(evensSlice, expectEvens) {
		t.Errorf("expected even values %v but got %v", expectEvens, evensSlice)
	}

	// filter: keep odds
	odds := OfOrdered(initial...)
	it = odds.Iterator()
	for it.HasNext() {
		num := it.Next()
		if num%2 == 0 {
			it.Remove()
		}
	}
	oddsSlice := odds.Values()
	expectOdds := []int{1, 3, 5, 7, 9}
	if !slices.Equal(oddsSlice, expectOdds) {
		t.Errorf("expected odd values %v but got %v", expectOdds, oddsSlice)
	}
}
