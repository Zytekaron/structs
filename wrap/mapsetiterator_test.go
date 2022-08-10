package wrap

import "testing"

func TestMapSetIterator(t *testing.T) {
	const count = 64
	found := make([]bool, count)

	// create a map of indices
	data := map[int]struct{}{}
	for i := 0; i < count; i++ {
		data[i] = struct{}{}
	}

	// mark all indices from the wrapping iterator
	it := MapSet(data)
	for it.HasNext() {
		key := it.Next()
		found[key] = true
	}

	// ensure no values were skipped
	for i, ok := range found {
		if !ok {
			t.Errorf("missing value %d in iterator", i)
		}
	}
}
