package bigg

import (
	"github.com/huydx/murmur3"
	"math/rand"
	"time"
)

type minhash struct {
	trial     int
	randrange int
}

const (
	MAXUINT32 = ^uint32(0)
	MAXUINT   = ^uint(0)
)

func NewMinHash(t int) *minhash {
	d := new(minhash)
	d.trial = t
	d.randrange = 99999
	return d
}

func (m *minhash) Hash(set []string, seed uint32) uint64 {
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

func (m *minhash) Jaccard(set1 []string, set2 []string) float32 {
	correct := 0
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < m.trial; i++ {
		rand_seed := uint32(rand.Intn(m.randrange))
		min_hash1 := m.Hash(set1, rand_seed)
		min_hash2 := m.Hash(set2, rand_seed)

		if min_hash1 == min_hash2 {
			correct += 1
		}
	}

	return float32(correct) / float32(m.trial)
}
