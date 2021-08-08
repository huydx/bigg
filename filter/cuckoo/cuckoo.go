package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash/fnv"
	"log"
)

const (
	N = 1000
	maxLoop = 100
)

var (
	T1 = make([]uint64, N)
	T2 = make([]uint64, N)
)

func hash(x uint64) (uint64, uint64) {
	hasher := fnv.New64()
	b := make([]byte, 8) // construct 64 bit string
	binary.LittleEndian.PutUint64(b, uint64(x))
	_, err := hasher.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	t1 := hasher.Sum64()

	// reuse t1 to not make second hash function, is that ok?
	hasher.Reset()
	binary.LittleEndian.PutUint64(b, t1)
	_, err = hasher.Write(b)
	if err != nil {
		log.Fatal(err)
	}

	t2 := hasher.Sum64()

	return t1, t2
}

func insert(x uint64, cnt int) error {
	for cnt < maxLoop {
		h1, h2 := hash(x)
		idx1, idx2 := h1%N, h2%N
		if T1[idx1] == 0 {
			T1[idx1] = x
			return nil
		}
		x, T1[idx1] = T1[idx1], x
		if T2[idx2] == 0 {
			T2[idx2] = x
			return nil
		}
		x, T2[idx2] = T2[idx2], x
		cnt++
	}
	return errors.New("need rehash")
}

func query(x uint64) bool {
	h1, h2 := hash(x)
	idx1, idx2 := h1%N, h2%N
	if T1[idx1] == x || T2[idx2] == x {
		return true
	} else {
		return false
	}
}

// how to do it?
func rehash() {

}



func main() {
	_ = insert(10, 0)
	_ = insert(20, 0)
	_ = insert(30, 0)
	_ = insert(40, 0)
	_ = insert(50, 0)
	fmt.Println(query(10))
	fmt.Println(query(20))
	fmt.Println(query(40))
	fmt.Println(query(60))
}
