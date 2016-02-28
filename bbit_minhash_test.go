package bigg

import (
	"fmt"
	"testing"
)

func TestPopCountHD(t *testing.T) {
	x := uint64(1)
	y := uint64(7)
	z := uint64(8)
	if popcountHD(x) != 1 || popcountHD(y) != 3 || popcountHD(z) != 1 {
		t.Errorf("pop count mismatch")
	}
}

func TestRefBBitMinHash(t *testing.T) {
	m := NewBBitMinHash(1, 128)
	set1 := []string{"ff", "asda"}
	set2 := []string{"ff"}
	fmt.Println(m.Jaccard(set1, set2))
}
