package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func run4(input []byte) {
	s := strings.Split(string(input), "\n")
	N, _ := strconv.ParseInt(s[0], 10, 64)
	arr := make([]int, 0, N)
	for _, c := range s[1:] {
		a, _ := strconv.ParseInt(c, 10, 64)
		arr = append(arr, int(a))
	}

	sum1 := int64(0)
	for i := 1; i <= int(N); i++ {
		sum1 += int64(i)
	}

	sum2 := int64(0)
	for _, i := range arr {
		sum2 += int64(i)
	}

	abs := int(math.Abs(float64(sum2 - sum1)))
	sum := sum2 + sum1
	fmt.Println(abs)
	fmt.Println(sum)
}
