package udfs

import "testing"

func TestNewUnionfind(t *testing.T) {
	u := NewUnionfind(5)
	u.unionSet(0, 1)
	if u.numDisjointSets() != 4 {
		t.Errorf("need 4")
	}
	u.unionSet(2, 3)
	if u.numDisjointSets() != 3 {
		t.Errorf("need 3")
	}
	u.unionSet(4, 3)
	if u.numDisjointSets() != 2 {
		t.Errorf("need 2")
	}

	if u.isSameSet(0, 3) {
		t.Errorf("need false")
	}

	if !u.isSameSet(4, 3) {
		t.Errorf("need true")
	}
}
