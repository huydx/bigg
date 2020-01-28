package main

import "fmt"

type arr []int

// implement next permutation as knuth's TAO's L algorithm
func (a arr) permute() bool {
	if len(a) <= 1 {
		return false
	}
	j := len(a) - 2
	for ; j > 0 && a[j] > a[j+1]; j -= 1 {
	}
	if j == -1 {
		return false
	}

	l := len(a) - 1
	a.swap(j, l)

	lo := j + 1
	hi := len(a) - 1
	for {
		if lo >= hi {
			break
		}
		a.swap(lo, hi)
		lo += 1
		hi -= 1
	}
	return true
}

func (a arr) swap(i int, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

func main() {
	a := arr([]int{1, 2, 3, 4})
	a.permute()
	fmt.Println(a)
	a.permute()
	fmt.Println(a)
	a.permute()
	fmt.Println(a)
}
