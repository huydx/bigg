package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type window struct {
	w []int
}

func (w *window) insert(n int) bool {
	if n < w.w[0] {
		return false
	} else {
		// remove min at 0 idx
		w.w[0] = n
		sort.Slice(w.w, func(i, j int) bool {
			return w.w[i] < w.w[j]
		})
		return true
	}
}

func run3(input []byte) {
	in := string(input)
	arr := make([]int, 0)
	for _, s := range strings.Split(in, " ") {
		a, _ := strconv.ParseInt(s, 10, 64)
		arr = append(arr, int(a))
	}

	w := make([]int, 3)
	for i := range w {
		w[i] = -1
	}

	wd := window{
		w: w,
	}

	for _, a := range arr {
		wd.insert(a)
	}

	fmt.Println(wd.w[0])
}
