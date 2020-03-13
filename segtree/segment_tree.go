package main

import "fmt"

type SegmentTree struct {
	st []int
	A  []int
	n  int
}

func NewSegmentTree(init []int) *SegmentTree {
	n := len(init)
	A := init
	st := make([]int, 4*len(A))
	for i := 0; i < 4*n; i++ {
		st[i] = 0
	}

	s := &SegmentTree{
		A:  init,
		n:  len(init),
		st: st,
	}
	s.build(1, 0, n-1)
	return s
}

func (s *SegmentTree) left(p int) int {
	return p << 1
}

func (s *SegmentTree) right(p int) int {
	return (p << 1) + 1
}

func (s *SegmentTree) build(p int, L int, R int) {
	if L > R {
		panic("invalid")
	}
	if L == R {
		s.st[p] = L
	} else {
		s.build(s.left(p), L, (L+R)/2)
		s.build(s.right(p), (L+R)/2+1, R)
		p1 := s.st[s.left(p)]
		p2 := s.st[s.right(p)]
		if s.A[p1] <= s.A[p2] {
			s.st[p] = p1
		} else {
			s.st[p] = p2
		}
	}
}

func (s *SegmentTree) rmq0(p int, L int, R int, i int, j int) int {
	if i > R || j < L {
		return -1
	}
	if L >= i && R <= j {
		return s.st[p]
	}

	p1 := s.rmq0(s.left(p), L, (L+R)/2, i, j)
	p2 := s.rmq0(s.right(p), (L+R)/2+1, R, i, j)

	if p1 == -1 {
		return p2
	}
	if p2 == -1 {
		return p1
	}

	if s.A[p1] <= s.A[p2] {
		return p1
	} else {
		return p2
	}
}

func (s *SegmentTree) rmq(i int, j int) int {
	return s.rmq0(1, 0, s.n-1, i, j)
}

func main() {
	a := []int{18, 17, 13, 19, 15, 11, 20}
	s := NewSegmentTree(a)
	fmt.Println(s.rmq(1, 3))
	fmt.Println(s.rmq(4, 6))
	fmt.Println(s.rmq(3, 4))
	fmt.Println(s.rmq(0, 0))
	fmt.Println(s.rmq(0, 1))
}
