package main

import (
	"runtime"
	"time"
)

// lamport's bakery algorithm for software lock
// the idea is number machine represent by `tickets`
// each thread given a ticket number which monotonically incremented
// 2 thread could have same ticket, so we use thread id as priority evaluator for the case
type bakery struct {
	n        int
	entering []bool
	tickets  []int
}

func (b *bakery) lock(pid int) {
	b.entering[pid] = true
	b.tickets[pid] = max(b.tickets) + 1
	b.entering[pid] = false

	for j := 0; j < b.n; j++ {
		for ; b.entering[j]; {
			runtime.Gosched()
		}
		for ; b.tickets[j] != 0 && (b.tickets[pid] > b.tickets[j] || b.tickets[pid] == b.tickets[j] && pid > j); {
			runtime.Gosched()
		}
	}
}

func (b *bakery) unlock(pid int) {
	b.tickets[pid] = 0
}

func max(arr []int) int {
	var max = -1
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func compare(i1, i2, j1, j2 int) bool {
	if i1 < j1 {
		return true
	} else if i1 > j1 {
		return false
	}
	return i2 < j2
}

func main() {
	m := make(map[int]int)
	n := 10000

	bak := &bakery{
		n:        n,
		entering: make([]bool, n+1),
		tickets:  make([]int, n+1),
	}

	for i := 0; i < n; i++ {
		go func(i int) {
			bak.lock(i)
			m[i] = i
			bak.unlock(i)
		}(i)
	}

	for {
		if len(m) == n {
			break
		}
		time.Sleep(time.Second)
	}
}
