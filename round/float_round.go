// Copyright (c) 2016 LINE Corporation. All rights reserved.
// LINE Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.

package main

import (
	"math"
	"fmt"
)

const (
	mask     = 0x7FF // exponent mask
	shift    = 64 - 11 - 1 // mantissa bit num
	bias     = 1023
)

// test code for https://go-review.googlesource.com/c/43652/4/src/math/floor.go#b52
func Round(x float64) float64 {
	const (
		signMask = 1 << 63
		fracMask = (1 << shift) - 1
		halfMask = 1 << (shift - 1)
		one      = bias << shift
	)
	bits := math.Float64bits(x)
	e := uint(bits>>shift) & mask

	switch {
	case e < bias:
		// Round abs(x)<1 including denormals.
		bits &= signMask // +-0
		if e == bias-1 {
			bits |= one // +-1
		}
	case e < bias+shift:
		// Round any abs(x)>=1 containing a fractional component [0,1).
		e -= bias
		bits += halfMask >> e
		bits &^= fracMask >> e
	}

	return math.Float64frombits(bits)
}


func main() {
	fmt.Println(Round(122.22))
	fmt.Println(Round(122.62))
	fmt.Println(Round(122.65))
	fmt.Println(Round(-1.2))
	fmt.Println(Round(-1.5))
	fmt.Println(Round(-1.6))
	fmt.Println(Round(-1.64))
}
