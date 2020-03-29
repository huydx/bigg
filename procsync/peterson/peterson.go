package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	producer = 0
	consumer = 1
)

// ring is simple ring buffer with max size
type ring struct {
	arr  []interface{}
	head int // pos of head
	tail int // next pos to insert
	size int
	max  int
}

func (r *ring) enqueue(i interface{}) error {
	size := r.size + 1
	if size > r.max {
		return fmt.Errorf("queue full")
	}
	r.arr[r.tail] = i
	r.tail = (r.tail + 1) % r.max
	r.size += 1
	return nil
}

func (r *ring) dequeue() interface{} {
	if r.size == 0 {
		return nil
	}
	h := r.arr[r.head]
	r.head = (r.head + 1) % r.max
	r.size -= 1
	return h
}

type SPSC struct {
	flag []bool
	turn int
	ring *ring
}

// implement peterson algorithm by lock(i) with i == producer or consumer
// just consider producer and consumer are separate processes
func (pc *SPSC) produce(i interface{}) {
	j := 1 - producer
	pc.flag[producer] = true
	turn := j
	for ; pc.flag[j] && turn == j; {
		runtime.Gosched()
	}
	_ = pc.ring.enqueue(i)
	pc.flag[producer] = false
}

func (pc *SPSC) consume() interface{} {
	j := 1 - consumer
	pc.flag[consumer] = true
	turn := j
	for ; pc.flag[j] && turn == j; {
		runtime.Gosched()
	}
	d := pc.ring.dequeue()
	pc.flag[consumer] = false
	return d
}

func main() {
	spsc := &SPSC{
		flag: []bool{false, false},
		ring: &ring{max: 10000, arr: make([]interface{}, 10000)},
	}

	N := 1000

	// to test spsc, we just produce N items
	// and wait until consumer success to consume N items using wait group
	wg := sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			for {
				if d := spsc.consume(); d == nil {
					time.Sleep(time.Second)
				} else {
					wg.Done()
					return
				}
			}
		}()
	}

	for i := 0; i < N; i++ {
		go spsc.produce(1)
	}

	wg.Wait()
}
