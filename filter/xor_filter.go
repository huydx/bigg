package main

import (
	"fmt"
	"github.com/FastFilter/xorfilter"
)

// key x, produce k-bit fingerprint
// construct map F from all elements to k-bit integers
// construction: acyclic 3 partite random hypergraph???
// pick hash by gen random seed
// membership: cal h0,h1,h2 -> cal fingerprints from entries -> compare fingerprints
// XOR filters have one main disadvantage relative to Bloom filters,
// and that's that all the items to store in the XOR filter must be known
// in advance before the filter is constructed.
// This contrasts with Bloom filters, where items can be added incrementally over a long period of time
// https://stackoverflow.com/questions/67527507/what-is-an-xor-filter
// http://web.stanford.edu/class/archive/cs/cs166/cs166.1216/lectures/13/Slides13.pdf#page=49

// why static set is ok for sstable?
// construct once, immutable

func main() {
	keys := []uint64{1,2,3,4,5,6,7}
	filter, err := xorfilter.PopulateBinaryFuse8(keys)
	if err != nil {
		panic(err)
	}
	fmt.Println(filter.Contains(1))
	fmt.Println(filter.Contains(10))
}
