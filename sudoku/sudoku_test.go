package main

import "testing"

func TestValidRowCol(t *testing.T) {
	in := `
159486372
728965413
692534781
937258146
481723695
875691234
516349827
394172568
046853219`

	g := grid(make([]int, 81))
	g.from(in)
	if ok := g.valid(8, 0, 7); ok {
		t.Fatalf("not valid due to 7")
	}
}

func TestValidSquare(t *testing.T) {
	in := `
159486372
728965413
692534780
937258146
461723895
875619234
516347928
384192567
243871659`

	g := grid(make([]int, 81))
	g.from(in)
	if ok := g.valid(2, 8, 1); ok {
		t.Fatalf("not valid due to 1")
	}
}
