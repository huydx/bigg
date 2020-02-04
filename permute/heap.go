package main

import "fmt"

// heap's algorithm

func main() {
	heap([]int{1,2,3,5,6,7}, 6, 6)
}

func heap(a []int, size int, n int) {
	if size == 1 {
		print(a)
	}

	for i := 0; i < size; i++ {
		heap(a, size-1, n)
		if size%2 == 1 {
			swap(a, 0, size-1)
		} else {
			swap(a, i, size-1)
		}
	}
}

func swap(arr []int, i int, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t
}

func print(arr []int) {
	for i := range arr {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}
