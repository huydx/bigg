package main

import "fmt"

type Heap struct {
	heap []int
	pos int
}

func NewHeap() *Heap {
	h := &Heap{
		heap: make([]int, 1),
		pos: 0,
	}
	return h
}


func (h *Heap) insert(item int) {
	h.heap = append(h.heap, item)
	var i = len(h.heap) - 1
	h.bubleUp(i)
}

func (h *Heap) deleteMin() int {
	m := h.min()
	s := len(h.heap)-1
	h.heap[1] = h.heap[s]
	h.heap = h.heap[:s]

	// min heapify root
	h.bubleDown(1)

	return m
}

func (h *Heap) bubleDown(i int) {

}

func (h *Heap) bubleUp(i int) {
	var p = h.parentIndex(i)
	for {
		if p != i && h.heap[i] < h.heap[p] {
			h.swap(i, p)
			i = p
			p = h.parentIndex(i)
	 	} else {
	 		break
		}
	}
}

func (h *Heap) min() int {
	return h.heap[1]
}

func (h *Heap) parentIndex(i int) int {
	if i == 1 { return 1 }
	return i / 2
}

func (h *Heap) swap(i int, j int) {
	tmp := h.heap[j]
	h.heap[j] = h.heap[i]
	h.heap[i] = tmp
}

func (h *Heap) isEmpty() bool {
	return len(h.heap) == 0
}

func (h *Heap) leftChild(i int) int {
	return h.heap[2*i]
}

func (h *Heap) rightChild(i int) int {
	return h.heap[2*i+1]
}

func main() {
	h := NewHeap()
	h.insert(10)
	h.insert(1)
	h.insert(22)
	h.insert(23)
	h.insert(5)
	h.insert(4)
	h.insert(-1)
	h.insert(3)
	h.insert(-2)
	fmt.Println(h)
}
