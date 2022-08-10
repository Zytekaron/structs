package bloom

import (
	"github.com/vcaesar/murmur"
	"github.com/zytekaron/structs/bitset"
)

// Filter is an implementation of a bloom filter.
type Filter struct {
	bits     *bitset.BitSet
	capacity int
	hashes   int
}

// New creates a new Filter.
func New(capacity, hashes int) *Filter {
	return &Filter{
		bits:     bitset.New(capacity),
		capacity: capacity,
		hashes:   hashes,
	}
}

// NewWithBitSet creates a new Filter with an existing BitSet.
func NewWithBitSet(hashes int, bits *bitset.BitSet) *Filter {
	return &Filter{
		bits:     bits,
		capacity: bits.Capacity(),
		hashes:   hashes,
	}
}

// Add adds a value to the bloom filter.
func (f *Filter) Add(bytes []byte) {
	for i := 0; i < f.hashes; i++ {
		index := int(murmur.Murmur3(bytes, uint32(i))) % f.capacity
		f.bits.Set(index)
	}
}

// AddString adds a string to the bloom filter.
func (f *Filter) AddString(str string) {
	f.Add([]byte(str))
}

// Test tests whether a value is present in the bloom filter.
//
// When the result is false, the value is not present in the bloom filter.
// When the result is true, the value may have been added to the
// bloom filter, but it is not guaranteed to be present.
func (f *Filter) Test(bytes []byte) bool {
	for i := 0; i < f.hashes; i++ {
		index := int(murmur.Murmur3(bytes, uint32(i))) % f.capacity
		if !f.bits.Get(index) {
			return false
		}
	}
	return true
}

// TestString tests whether a string is present in the bloom filter.
//
// When the result is false, the value is not present in the bloom filter.
// When the result is true, the value may have been added to the
// bloom filter, but it is not guaranteed to be present.
func (f *Filter) TestString(str string) bool {
	return f.Test([]byte(str))
}

// TestAdd tests whether a value is in the bloom filter, and adds it in the process.
//
// Equivalent to a call to Test and then Add, but more efficient than calling both.
func (f *Filter) TestAdd(bytes []byte) bool {
	present := true
	for i := 0; i < f.hashes; i++ {
		index := int(murmur.Murmur3(bytes, uint32(i))) % f.capacity
		if !f.bits.Get(index) {
			f.bits.Set(index)
			present = false
		}
	}
	return present
}

// TestAddString tests whether a string is in the bloom filter, and adds it in the process.
//
// Equivalent to a call to Test and then Add, but more efficient than calling both.
func (f *Filter) TestAddString(str string) bool {
	return f.TestAdd([]byte(str))
}

// Clear clears the bloom filter, resetting every bit in the internal bit set.
func (f *Filter) Clear() {
	f.bits.Clear()
}
