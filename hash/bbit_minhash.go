package bigg

import (
	"fmt"
	"github.com/huydx/murmur3"
	"math/rand"
	"time"
)

type bbitMinHash struct {
	bitnum  uint64
	hashnum uint64
}

func NewBBitMinHash(b uint64, h uint64) *bbitMinHash {
	d := new(bbitMinHash)
	d.bitnum = b
	d.hashnum = h
	return d
}

func (b *bbitMinHash) MinHash(set []string, seed uint64) uint64 {
	hash := murmur3.New64(seed)
	minHash := ^uint64(0)

	for _, item := range set {
		hash.Write([]byte(item))
		s := hash.Sum64()
		if s < minHash {
			minHash = s
		}
	}

	return minHash
}

// Straight and simple C to Go translation from https://en.wikipedia.org/wiki/Hamming_weight
func popcountHD(x uint64) uint64 {
	x -= (x >> 1) & 0x55555555
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	return uint64(x & 0x0000003f)
}

func (b *bbitMinHash) Signature(set []string) uint64 {
	rand.Seed(time.Now().UnixNano())
	sig := uint64(0)
	mask := uint64(0x1)

	for i := uint64(0); i < b.hashnum; i++ {
		rand_seed := uint64(rand.Intn(99999))
		mhash := b.MinHash(set, rand_seed)

		sig += (mhash & mask) << i
	}

	return sig
}

func printAsBitArray(in uint64) {
	for in > 0 {
		if in&1 == 1 {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
		in = in >> 1
	}
}

func (b *bbitMinHash) Jaccard(set1 []string, set2 []string) float64 {
	s1 := b.Signature(set1)
	s2 := b.Signature(set2)

	printAsBitArray(s1)
	fmt.Println()
	printAsBitArray(s2)
	fmt.Println()

	common := s1 ^ s2

	return 2.0 * ((float64(b.hashnum)-float64(popcountHD(common)))/float64(b.hashnum) - 0.5)
}
