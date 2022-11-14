package wrap

import (
	"github.com/zytekaron/structs"
	"testing"
)

func TestSlice(t *testing.T) {
	const count = 10
	s := make([]int, count)
	for i := 0; i < count; i++ {
		s[i] = i
	}

	var _ structs.List[int] = Slice(nil, s)
	var _ structs.List[int] = Values(nil, s...)
	// todo: no meaningful test right now,
	//  so just testing for panics :shrug:
}
