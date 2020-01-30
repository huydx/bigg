package main

import (
	"fmt"
	"math"
)

func run1(input []byte) {
	l := len(input)
	if l != 3 {
		fmt.Println("error")
		return
	}

	idx := 0
	res := 0
	for i := l - 1; i >= 0; i-- {
		j, err := toint(input[i])
		if err != nil {
			fmt.Println("error")
			return
		} else {
			res += j * int(math.Pow(10, float64(idx)))
		}
		idx++
	}
	fmt.Println(res * 2)
}

func toint(c byte) (int, error) {
	i := int(c) - 48
	if i < 0 || i > 9 {
		return 0, fmt.Errorf("not char")
	}
	return i, nil
}
