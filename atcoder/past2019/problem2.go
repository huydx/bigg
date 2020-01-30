package main

import (
	"fmt"
	"strconv"
	"strings"
)

func run2(input []byte) {
	in := string(input)

	var (
		arr = make([]int, 0)
	)

	for i, n := range strings.Split(in, "\n") {
		if i == 0 {
			// pass
		} else {
			a, _ := strconv.ParseInt(n, 10, 64)
			arr = append(arr, int(a))
		}
	}

	solve2(arr)
}

func solve2(arr []int) {
	if len(arr) == 1 {
		return
	}
	prev := arr[0]
	for _, a := range arr[1:] {
		if a == prev {
			fmt.Println("stay")
		} else if a < prev {
			fmt.Printf("down %d\n", prev-a)
		} else if a > prev {
			fmt.Printf("up %d\n", a-prev)
		}
		prev = a
	}
}
