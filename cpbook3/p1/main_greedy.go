package main

import (
	"fmt"
	"math"
)

/**
An illustration: UVa Online Judge [47] Problem Number 10911 (Forming Quiz Teams). Abridged Problem Description:
Let (x, y) be the coordinates of a student’s house on a 2D plane.
There are 2N students and we want to pair them into N groups.
Let di be the distance between the houses of 2 students in group i.
Form N groups such that cost = 􏰀Ni=1 di is minimized.
Output the minimum cost. Constraints: 1 ≤ N ≤ 8 and 0 ≤ x, y ≤ 1000.
Sample input:
N = 2; Coordinates of the 2N = 4 houses are {1, 1}, {8, 6}, {6, 8}, and {1, 3}.
Sample output: cost = 4.83.
*/

type pair struct {
	x int
	y int
}

func main() {
	input := []pair{{x: 1, y: 1}, {x: 8, y: 6}, {x: 6, y: 8}, {x: 1, y: 3}}
	fmt.Println(solve(input))
}

func swap(p []pair, i int, j int) {
	var t pair
	t = p[i]
	p[i] = p[j]
	p[j] = t
}

func shuffle(p []pair, start int, end int, out chan [] pair) {
	if start == end {
		p2 := make([]pair, len(p))
		copy(p2, p)
		out <- p2
	} else {
		for i := start; i < end; i++ {
			swap(p, start, i)
			shuffle(p, start, end, out)
			swap(p, start, i)
		}
	}
}

func dist(p pair, q pair) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func sumdist(p []pair) float64 {
	sum := 0.0
	for i := 0; i < N; i += 2 {
		sum += dist(p[i], p[i+1])
	}
	return sum
}

func solve(input []pair) float64 {
	out := make(chan []pair, 1)
	shuffle(input, 0, len(input)-1, out)
	min := 0.0
	for {
		select {
		case t := <-out:
			s := sumdist(t)
			if min < s {
				min = s
			}
		}
	}
	return min
}
