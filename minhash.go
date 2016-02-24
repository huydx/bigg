package bigg

import (
	"github.com/huydx/murmur3"
	"math/rand"
	"time"
)

const (
	MAXUINT32 = ^uint32(0)
	MAXUINT   = ^uint(0)
	TRIAL     = 1919
)

func MinHash(set []string, seed uint32) uint64 {
	minHash := MAXUINT32

	hash := murmur3.New32(seed)

	for _, item := range set {
		hash.Write([]byte(item))
		s := hash.Sum32()
		if s < minHash {
			minHash = s
		}
	}

	return uint64(minHash)
}

func BBitMinHash(set []string) uint64 {
	return 0
}

func Jaccard(set1 []string, set2 []string) float32 {
	correct := 0
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < TRIAL; i++ {
		rand_seed := uint32(rand.Intn(99999))
		min_hash1 := MinHash(set1, rand_seed)
		min_hash2 := MinHash(set2, rand_seed)

		if min_hash1 == min_hash2 {
			correct += 1
		}
	}

	return float32(correct) / float32(TRIAL)
}
