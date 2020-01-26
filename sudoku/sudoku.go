package main

import (
	"fmt"
	"strings"
)

type grid []int

func (g grid) from(s string) {
	for i, c := range strings.ReplaceAll(s, "\n", "") {
		g[i] = int(c) - 48
	}
}

// i: row
// j: col
func (g grid) at(i int, j int) int {
	return g[i*9+j]
}

func (g grid) set(i int, j int, val int) {
	g[i*9+j] = val
}

func (g grid) setCopy(i int, j int, val int) grid {
	dest := grid(make([]int, 81))
	copy(dest, g)
	dest.set(i, j, val)
	return dest
}

// set to val and check valid at row i and col j, if not, reset it
func (g grid) valid(i int, j int, val int) bool {
	fillrow := map[int]bool{}
	fillcol := map[int]bool{}

	gc := g.setCopy(i, j, val)

	// validate row and col
	for ii := 0; ii < 9; ii++ {
		if gc.at(i, ii) != 0 {
			if _, ok := fillrow[gc.at(i, ii)]; ok {
				return false
			} else {
				fillrow[gc.at(i, ii)] = true
			}
		}

		if gc.at(ii, j) != 0 {
			if _, ok := fillcol[gc.at(ii, j)]; ok {
				return false
			} else {
				fillcol[gc.at(ii, j)] = true
			}
		}
	}

	// validate square
	fillsquare := map[int]bool{}
	baseI := (i / 3) * 3
	baseJ := (j / 3) * 3
	for ii := baseI; ii < baseI+3; ii += 1 {
		for jj := baseJ; jj < baseJ+3; jj += 1 {
			if gc.at(ii, jj) != 0 {
				if _, ok := fillsquare[gc.at(ii, jj)]; ok {
					return false
				} else {
					fillsquare[gc.at(ii, jj)] = true
				}
			}
		}
	}

	return true
}

func (g grid) next() (int, int) {
	for i := range g {
		if g[i] == 0 {
			return i / 9, i % 9
		}
	}
	return -1, -1
}

func (g grid) print() {
	for i := 0; i < 81; i += 9 {
		for j := 0; j < 9; j++ {
			fmt.Print(g[i+j])
		}
		fmt.Println()
	}
}

func main() {
	input := `
150086070
708900000
090500081
900008106
001020090
075600200
506040020
300102500
040800009`

	var (
		g = grid(make([]int, 81))
	)
	g.from(input)
	solve(grid(g)).print()
}

func solve(in grid) grid {

	var (
		// stack
		q []grid
		// current elem
		cur grid
	)
	q = append(q, in)
	for {
		// pop from stack
		if len(q) > 0 {
			n := len(q) - 1
			cur = q[n]
			q = q[:n]
		}

		nextI, nextJ := cur.next()
		// reach end
		if nextJ == -1 {
			break
		}

		for i := 1; i <= 9; i++ {
			if ok := cur.valid(nextI, nextJ, i); ok {
				q = append(q, cur.setCopy(nextI, nextJ, i))
			}
		}
	}

	return cur
}
