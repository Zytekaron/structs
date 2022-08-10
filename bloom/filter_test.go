package bloom

import (
	"encoding/binary"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	const capacity = 10_000   // size of the internal bitset
	const hashes = 5          // number of hash functions
	const trueTests = 1_000   // numbers to test for true positives
	const falseTests = 5_000  // numbers to test for false positives
	const falseRate = 1 / 20. // allowed rate of false positives (%)
	bf := New(capacity, hashes)
	for i := 0; i < trueTests; i++ {
		bf.Add(testIntToBytes(i))
	}

	for i := 0; i < trueTests; i++ { // first 10% of tests
		if !bf.Test(testIntToBytes(i)) {
			t.Errorf("expected int %d to be present in bloom filter", i)
		}
	}

	falsePositives := 0
	for i := trueTests; i < trueTests+falseTests; i++ { // remaining 90% of tests
		if bf.Test(testIntToBytes(i)) {
			falsePositives++
		}
	}

	falsePercentage := float64(falsePositives) / falseTests
	if falsePercentage > falseRate {
		t.Errorf("too many false positives: expected less than %.2f%%, got %.2f%%", falseRate*100, falsePercentage*100)
	}
}

func testIntToBytes(i int) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(i))
	return bytes
}
