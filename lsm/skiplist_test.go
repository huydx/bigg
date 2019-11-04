

package lsm

import (
	"testing"
)

func TestSkiplist(t *testing.T) {
	skl := New()
	skl.Insert(&Key{"a", "bar"})
	skl.Insert(&Key{"b", "bar2"})
	skl.Insert(&Key{"c", "bar3"})
	skl.Insert(&Key{"d", "bar4"})
	skl.Insert(&Key{"e", "bar5"})
	skl.Insert(&Key{"f", "bar5"})
	skl.Insert(&Key{"g", "bar5"})
	skl.Insert(&Key{"h", "bar5"})

	skl.draw()

	if skl.Search("g").v != "bar5" {
		t.Errorf("need bar")
	}

	if skl.Search("i") != nil {
		t.Errorf("need nil")
	}
}
