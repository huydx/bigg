

package lsm

import (
	"fmt"
	"math/rand"
	"sync/atomic"
)

// Skiplist
// O(logN) search
// O(logN) insert

func init() {
	rand.Seed(1)
}

const (
	kMaxHeight = 32
	kBranch    = 4
)

/**
HEAD                                                        TAIL
   |----------------------------->e-----------------------------|
   |                              |                             |
   |----------------->c---------->e---------->g-----------------|
   |                  |           |           |                 |
   |<---->a<--->b<--->c<--->d<--->e<--->f<--->g<--->h<--->i<--->|
 */

type Key struct {
	k string
	v interface{}
}

// node represent lowest layer?
type SkipListNode struct {
	val  *Key
	next []*SkipListNode // note: point of this ds
}

type Skiplist struct {
	height int32
	head   *SkipListNode // link to lowest lane
}

func New() *Skiplist {
	head := &SkipListNode{
		val:  &Key{k: "", v: nil},
		next: make([]*SkipListNode, kMaxHeight),
	}

	return &Skiplist{
		height: 0,
		head:   head,
	}
}

func (s *Skiplist) Insert(k *Key) {
	level := int32(0)
	for ; rand.Int31n(2) == 1; level++ {
		if level > s.height {
			atomic.StoreInt32(&s.height, level)
			break
		}
	}
	fmt.Println(level)
	node := &SkipListNode{
		val:  k,
		next: make([]*SkipListNode, level+1),
	}
	current := s.head
	for i := s.height; i >= 0; i-- { //note: index bug
	Search:
		for ; current.next[i] != nil; current = current.next[i] {
			if current.next[i].val.k > k.k {
				break Search
			}
		}
		if i > level {
			continue
		}
		node.next[i] = current.next[i]
		current.next[i] = node
	}
}

func (s *Skiplist) Search(k string) *Key {
	current := s.head
	for i := s.height; i >= 0; i-- {
		for ; current.next[i] != nil; current = current.next[i] {
			if current.next[i].val.k > k {
				break
			}
			if current.next[i].val.k == k {
				return current.next[i].val
			}
		}
	}
	return nil
}
