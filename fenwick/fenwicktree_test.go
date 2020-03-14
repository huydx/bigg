package fenwick

import "testing"

func TestNewFenwickTree(t *testing.T) {
	f := []int{2,4,5,5,6,6,6,7,7,8,9}
	fw := NewFenwickTree(10)
	for i := 0; i < 11; i++ {
		fw.adjust(f[i],1)
	}

	if fw.rsq(1,1) != 0 {
		t.Errorf("expect 0")
	}
	if fw.rsq(1,2) != 1 {
		t.Errorf("expect 1")
	}
	if fw.rsq(1,10) != 11 {
		t.Errorf("expect 11")
	}
	if fw.rsq(3,6) != 6 {
		t.Errorf("expect 6")
	}
}
