package wrap

import "testing"

func TestSlice(t *testing.T) {
	const count = 10
	s := make([]int, count)
	for i := 0; i < count; i++ {
		s[i] = i
	}

	Slice(nil, s)
	Values(nil, s...)
	// todo: no meaningful test right now,
	//  so just testing for panics :shrug:
}
