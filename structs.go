package structs

import (
	"golang.org/x/exp/constraints"
)

// PredicateFunc is a function which returns a true or false
// value for a given input, for example a function used to
// indicate whether a particular value is a prime number.
type PredicateFunc[V any] func(v V) (result bool)

// LessFunc is a function which returns whether the first value
// is strictly smaller than (not equal to) the second value.
type LessFunc[V any] func(a, b V) (less bool)

// EqualFunc is a function which returns whether the first value
// is equal to the second value, using custom behavior.
//
// EqualOrdered can be used for any type which implements constraints.Ordered.
type EqualFunc[V any] func(a, b V) (equal bool)

// CompareFunc is a function which returns an integer indicating
// the difference between the first and second value.
//  n == 0: a == b
//  n < 0: a < b
//  n > 0: a > b
// CompareOrdered or ReverseCompareOrdered can be used
// for any type which implements constraints.Ordered.
type CompareFunc[V any] func(a, b V) (result int)

// CapacityFunc is a function used to return a new capacity for a slice
// when it needs to be increased. It must return an integer equal to or
// greater than need, otherwise it may cause a panic from the caller.
type CapacityFunc func(old, need int) int

// LessOrdered is a LessFunc which compares two values of any type which
// implements constraints.Ordered, returning their natural ordering.
//  a < b
func LessOrdered[V constraints.Ordered](a, b V) bool {
	return a < b
}

// GreaterOrdered is a LessFunc which compares two values of any type which
// implements constraints.Ordered, returning the inverse of their natural ordering.
//  b < a
func GreaterOrdered[V constraints.Ordered](a, b V) bool {
	return b < a
}

// CompareOrdered is a CompareFunc which compares two values of any type
// which implements constraints.Ordered, returning their natural ordering.
//  a == b: 0
//  a < b: -1
//  a > b: +1
func CompareOrdered[V constraints.Ordered](a, b V) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// ReverseCompareOrdered is a CompareFunc which compares two values of any type
// which implements constraints.Ordered, returning the inverse of their natural ordering.
//  a == b: 0
//  a < b: +1
//  a > b: -1
func ReverseCompareOrdered[V constraints.Ordered](a, b V) int {
	if a > b {
		return -1
	}
	if a < b {
		return 1
	}
	return 0
}

// Reverse reverses a comparison function, effectively reversing the
// sorting order of the values based on an existing compare function.
func Reverse[V any](cmp CompareFunc[V]) CompareFunc[V] {
	return func(a, b V) int {
		return -cmp(a, b)
	}
}

// EqualOrdered is an EqualFunc which operates on a type
// which satisfies the constraints.Ordered interface.
func EqualOrdered[V constraints.Ordered](a, b V) bool {
	return a == b
}

func Filter[V any](collection Collection[V], keep PredicateFunc[V]) {
	it := collection.Iterator()
	for it.HasNext() {
		if !keep(it.Next()) {
			it.Remove()
		}
	}
}

// Realloc reallocates a slice with the new size.
// It can be used to increase or decrease the slice size.
func Realloc[T any](size int, s []T) []T {
	res := make([]T, size)
	if size > len(s) {
		copy(res, s)
	} else {
		copy(res, s[:size])
	}
	return res
}

// DoubleCapacity returns a capacity double that of the old capacity
// until it exceeds need, or 1 if the old capacity was 0.
func DoubleCapacity(before, need int) (after int) {
	if before == 0 {
		return 1
	}
	for before < need {
		before <<= 1
	}
	return before
}

//// HalveCapacity returns a capacity half that of the old capacity.
//func HalveCapacity(before, need int) (after int) {
//	return before >> 1
//}
