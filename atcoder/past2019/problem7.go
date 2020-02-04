package main

var (
	mtx [][]int
)


func sum(arr []int) int {
	s := 0
	for i := range arr {
		for j := i; j < len(arr); j++ {
			s += mtx[i][j]
		}
	}
	return s
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
