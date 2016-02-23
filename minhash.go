package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
)

func HashBytes(b []byte) []byte {
	h1, h2 := murmur3.Sum128(b)
	return []byte{
		byte(h1 >> 56), byte(h1 >> 48), byte(h1 >> 40), byte(h1 >> 32),
		byte(h1 >> 24), byte(h1 >> 16), byte(h1 >> 8), byte(h1),
		byte(h2 >> 56), byte(h2 >> 48), byte(h2 >> 40), byte(h2 >> 32),
		byte(h2 >> 24), byte(h2 >> 16), byte(h2 >> 8), byte(h2),
	}
}

func HashString(s string) []byte {
	return HashBytes([]byte(s))
}

func CalcMinHash(set []string) string {
	//minHash := -1
	for _, item := range set {
		s := string(HashString(item))
		fmt.Printf("%s", s)
	}
	return ""
}

func main() {
	CalcMinHash([]string{"ff", "yy", "zz"})
}
