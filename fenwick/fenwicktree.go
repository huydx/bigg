package fenwick

type FenwickTree struct {
	ft []int
}

func NewFenwickTree(n int) *FenwickTree {
	ft := make([]int, n+1)
	return &FenwickTree{ft: ft}
}

func (f *FenwickTree) rsq0(b int) int {
	sum := 0
	for ; b > 0; b -= LSOne(b) {
		sum += f.ft[b]
	}
	return sum
}

func (f *FenwickTree) rsq(a int, b int) int {
	if a == 1 {
		return f.rsq0(b)
	} else {
		return f.rsq0(b) - f.rsq0(a-1)
	}
}

func (f *FenwickTree) adjust(k int, v int) {
	for ; k < len(f.ft); k += LSOne(k) {
		f.ft[k] = f.ft[k] + v
	}
}

func LSOne(k int) int {
	return k & (-k)
}
