# structs

**version:** v0.1.1

### This library is not production-ready. Use it only in testing environments.

This project is an effort to provide a variety of data structures
and accompanying interfaces to Go developers using 1.18 or later,
and was originally inspired by the Java Collections API interfaces.

The API for the current types is not stable. Significant changes may be
made at any point. Some types may not fully implement interfaces which
they are meant to, and documentation remains mostly incomplete.

If there are any real or potential bugs you notice within any of these
data structures, incomplete test cases, or other problems, feel free
to create an issue. I welcome your feedback and support.

## Installation

```
go get github.com/zytekaron/structs
```

## Data Structures

- [`bitset`](./bitset) - A resizable bit set.
- [`bloom`](./bloom) - A bloom filter backed by [`bitset`](./bitset).
- [`heap`](./heap) - A binary heap.
- [`list`](./list) - A doubly linked list.
- [`queue`](./queue)
    - A regular double-ended queue backed by [`list`](./list).
    - A priority queue backed by [`heap`](./heap).

- [`wrap`](./wrap) - To use Go types as Collections (see examples below).

## Usage

```go
package main

import (
	"fmt"
	"github.com/zytekaron/structs/list"
	"github.com/zytekaron/structs/set"
	"github.com/zytekaron/structs/wrap"
)

func main() {
	l := list.NewOrdered[int]()
	l.Add(2)
	l.AddLast(3)
	l.AddFirst(1)
	l.Each(func(i int) { fmt.Println(i) }) // 1, 2, 3

	s := set.NewMapSet[int]()
	s.AddAll(wrap.OrderedValues(10, 20, 30))             // 10, 20, 30
	s.AddIterator(wrap.ValueIterator(30, 40, 50))        // 10, 20, 30, 40, 50
	s.RetainAll(wrap.OrderedSlice([]int{0, 20, 30, 60})) // 20, 30
	fmt.Println(s.Values())                              // [20, 30]
}
```

## Possible Additions

- optional reallocation (downsizing) when, for example `size / capacity < 0.25`

## License

**structs** is licensed under the [MIT License](./LICENSE).
