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
	fmt.Println(mindist(input))
}

func mindist(pairs []pair) float64 {
	if len(pairs) == 0 {
		return 0
	}
	var min = math.MaxFloat64
	p1 := pairs[0]
	for i, p := range pairs[1:] {
		d1 := dist(p1, p)
		pairs2 := make([]pair, 0, len(pairs)-2)
		for j := 1; j < len(pairs); j++ {
			if j != i+1 {
				pairs2 = append(pairs2, pairs[j])
			}
		}
		d2 := mindist(pairs2)


		if d1+d2 < min {
			min = d1 + d2
		}
	}
	return min
}

func dist(p1 pair, p2 pair) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}
