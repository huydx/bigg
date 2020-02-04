package main

import (
	"fmt"
	"strconv"
	"strings"
)

func run5(input []byte) {
	ss := strings.Split(string(input), "\n")
	nq := strings.Split(ss[0], " ")
	n, _ := strconv.ParseInt(nq[0], 10, 64)
	N := int(n)

	mtx := make([][]bool, N+1)
	for i := range mtx {
		mtx[i] = make([]bool, N+1)
	}

	following := func(i int) []int {
		r := make([]int, 0)
		for j := 1; j < N+1; j++ {
			if mtx[i][j] {
				r = append(r, j)
			}
		}
		return r
	}

	follower := func(i int) []int {
		r := make([]int, 0)
		for j := 1; j < N+1; j++ {
			if mtx[j][i] {
				r = append(r, j)
			}
		}
		return r
	}

	dofollow := func(i, j int) {
		mtx[i][j] = true
	}

	for _, s := range ss[1:] {
		cmds := strings.Split(s, " ")
		switch cmds[0] {
		case "1":
			i, _ := strconv.ParseInt(cmds[1], 10, 64)
			j, _ := strconv.ParseInt(cmds[2], 10, 64)
			dofollow(int(i), int(j))
		case "2":
			a, _ := strconv.ParseInt(cmds[1], 10, 64)
			for _, i := range follower(int(a)) {
				dofollow(int(a), int(i))
			}
		case "3":
			a, _ := strconv.ParseInt(cmds[1], 10, 64)
			for _, i := range following(int(a)) {
				for _, x := range follower(int(i)) {
					dofollow(int(a), x)
				}
			}
		default:
			panic("invalid command")
		}
	}

	for i := 1; i <= N; i++ {
		fmt.Println(mtx[i][1:])
	}
}
