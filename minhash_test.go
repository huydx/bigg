package bigg

import (
	"testing"
)

var data = []struct {
	set1    []string
	set2    []string
	jaccard float32
}{
	{[]string{"ff", "yy", "zz"}, []string{"ff"}, 0.33},
	{[]string{""}, []string{""}, 1.0},
	{[]string{"xxx", "yyy"}, []string{"xxx", "yyy"}, 1.0},
	{[]string{"x", "y"}, []string{"x"}, 0.5},
	{[]string{"x", "y", "z", "a", "b", "c"}, []string{"m"}, 0.0},
}

func approximate(x1 float32, x2 float32) bool {
	threshold := float32(0.2)
	if x1+threshold > x2 && x1-threshold < x2 {
		return true
	} else {
		return false
	}
}

func TestRefMinHash(t *testing.T) {
	for _, elem := range data {
		m := NewMinHash(128)
		jaccard := m.Jaccard(elem.set1, elem.set2)
		if !approximate(jaccard, elem.jaccard) {
			t.Errorf("jaccard value mismatch")
		}
	}
}
