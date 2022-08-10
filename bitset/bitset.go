package bitset

import (
	"math/bits"
	"strconv"
	"strings"
)

// the number of bits per element used in the bitset's data slice.
const blockBits = 8

// BitSet is an implementation of a resizable bit set.
type BitSet struct {
	data []uint8
	cap  int
}

// New creates a BitSet with the given bit capacity.
//
// All methods will panic if the specified bit index is out of range.
func New(capacity int) *BitSet {
	size := (capacity + blockBits - 1) / blockBits
	return &BitSet{
		data: make([]uint8, size),
		cap:  capacity,
	}
}

// Get gets the value of a bit.
func (b *BitSet) Get(bit int) bool {
	if bit < 0 || bit >= b.cap {
		panic("index out of range")
	}

	mask := uint8(1) << (bit % blockBits)
	return (b.data[bit/blockBits] & mask) == mask
}

// Set sets the value of a bit to 1.
func (b *BitSet) Set(bit int) {
	if bit < 0 || bit >= b.cap {
		panic("index out of range")
	}

	b.data[bit/blockBits] |= 1 << (bit % blockBits)
}

// Unset sets the value of a bit to 0.
func (b *BitSet) Unset(bit int) {
	if bit < 0 || bit >= b.cap {
		panic("index out of range")
	}

	b.data[bit/blockBits] &= ^(1 << (bit % blockBits))
}

// Toggle toggles the value of a bit.
func (b *BitSet) Toggle(bit int) {
	if bit < 0 || bit >= b.cap {
		panic("index out of range")
	}

	b.data[bit/blockBits] ^= 1 << (bit % blockBits)
}

// Not toggles all the bits in the BitSet
func (b *BitSet) Not() {
	for i := range b.data {
		b.data[i] = ^b.data[i]
	}
	b.clearExtraBits()
}

// Clear clears all the bits in the BitSet
func (b *BitSet) Clear() {
	for i := range b.data {
		b.data[i] = 0
	}
}

// IsEmpty returns whether every bit's value is 0.
func (b *BitSet) IsEmpty() bool {
	for _, v := range b.data {
		if v > 0 {
			return true
		}
	}
	return false
}

// Or performs the bitwise OR operation with another BitSet.
//
// If the other BitSet is larger than the current one, this
// operation will panic. If it is smaller, only the existing
// bits from the other BitSet will be used for the operation.
func (b *BitSet) Or(other *BitSet) {
	if b.cap < other.cap {
		panic("other length is larger than own length")
	}

	for i := range other.data {
		b.data[i] |= other.data[i]
	}
}

// And performs the bitwise AND operation with another BitSet.
//
// If the other BitSet is larger than the current one, this
// operation will panic. If it is smaller, only the existing
// bits from the other BitSet will be used for the operation.
func (b *BitSet) And(other *BitSet) {
	if b.cap < other.cap {
		panic("other size is larger than own size")
	}

	for i := range b.data {
		b.data[i] &= other.data[i]
	}
}

// Xor performs the bitwise XOR operation with another BitSet.
//
// If the other BitSet is larger than the current one, this
// operation will panic. If it is smaller, only the existing
// bits from the other BitSet will be used for the operation.
func (b *BitSet) Xor(other *BitSet) {
	if b.cap < other.cap {
		panic("other size is larger than own size")
	}

	for i := range other.data {
		b.data[i] ^= other.data[i]
	}
}

// Grow grows the BitSet to a greater capacity.
//
// Panics if the capacity is not greater than the current capacity.
func (b *BitSet) Grow(capacity int) {
	if capacity <= b.cap {
		panic("new capacity is not larger")
	}
	oldSize := len(b.data)

	b.cap = capacity
	newSize := (capacity + blockBits - 1) / blockBits

	slc := make([]uint8, newSize-oldSize)
	b.data = append(b.data, slc...)
}

// Shrink shrinks the BitSet to a smaller capacity.
//
// Panics if the capacity is not smaller than the current capacity.
func (b *BitSet) Shrink(capacity int) {
	if capacity < 1 {
		panic("new capacity is too small")
	}
	if capacity >= b.cap {
		panic("new capacity is not smaller")
	}

	b.cap = capacity
	newSize := (capacity + blockBits - 1) / blockBits

	newData := make([]uint8, newSize)
	copy(newData, b.data)
	b.data = newData
	b.clearExtraBits()
}

// Clone clones the BitSet, returning a new instance with the same bits set.
func (b *BitSet) Clone() *BitSet {
	set := New(b.cap)
	for i, value := range b.data {
		set.data[i] = value
	}
	return set
}

// Capacity returns the current capacity.
func (b *BitSet) Capacity() int {
	return b.cap
}

// String converts the internal bits to a binary string representation.
func (b *BitSet) String() string {
	var buf strings.Builder
	if b.cap%blockBits == 0 {
		for i := 0; i < len(b.data); i++ {
			n := bits.Reverse8(b.data[i])
			str := strconv.FormatUint(uint64(n), 2)
			buf.WriteString(strings.Repeat("0", blockBits-len(str)))
			buf.WriteString(str)
		}
	} else {
		for i := 0; i < len(b.data)-1; i++ {
			n := bits.Reverse8(b.data[i])
			str := strconv.FormatUint(uint64(n), 2)
			buf.WriteString(strings.Repeat("0", blockBits-len(str)))
			buf.WriteString(str)
		}

		n := bits.Reverse8(b.data[len(b.data)-1])
		str := strconv.FormatUint(uint64(n), 2)
		// development note: if bits beyond the capacity
		// are set, this will be negative and will panic.
		// call b.clearExtraBits() to fix this issue.
		buf.WriteString(strings.Repeat("0", b.cap%blockBits-len(str)))
		buf.WriteString(str)
	}

	return buf.String()
}

// clear bits between the capacity and size*blockBits when
// they may have been set. this ensures that inactive bits
// which may be grown into later will always remain zeroed.
// also fixes an issue within String() if OOB bits are set.
func (b *BitSet) clearExtraBits() {
	if b.cap%blockBits > 0 {
		mask := ^uint8(0) >> (b.cap - blockBits)
		b.data[len(b.data)-1] ^= mask
	}
}
