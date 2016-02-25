package bigg

import (
	"bytes"
	"github.com/huydx/murmur3"
)

type bbitMinHash struct {
	bitnum  int
	hashnum int
}

func NewBBitMinHash(b int, h int) *bbitMinHash {
	d := new(bbitMinHash)
	d.bitnum = b
	d.hashnum = h
	return d
}

func (b *bbitMinHash) Hash(set []string, seed uint32) []byte {
	hash := murmur3.New32(seed)
	mask := (1 << b.bitnum) - 1

	for _, item := range set {
		hash.Write([]byte(item))
		s := hash.Sum32()
		if s < minHash {
			minHash = s
		}
	}

	return []byte(minHash & mask)
}

func (b *bbitMinHash) Jaccard(set1 []string, set2 []string) float32 {
	correct := 0
	rand.Seed(time.Now().UnixNano())

	//http://research.microsoft.com/pubs/120078/wfc0398-lips.pdf
	//using above paper to estimate jaccard

	return 0
}
