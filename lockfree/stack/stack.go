package stack

import (
	"sync"
	"unsafe"
	"sync/atomic"
)

type node struct {
	elem interface{}
	next *node
}

type StackWithMut struct {
	root *node
	mut  sync.Mutex
}

func NewStackWithMut() StackWithMut {
	return StackWithMut{root: nil, mut: sync.Mutex{}}
}

func (s StackWithMut) Pop() interface{} {
	s.mut.Lock()
	defer s.mut.Unlock()
	r := s.root
	if r != nil {
		s.root = r.next
		return r.elem
	} else {
		return nil
	}
}

func (s StackWithMut) Push(e interface{}) {
	s.mut.Lock()
	sync.Mutex{}
	defer s.mut.Unlock()
	if s.root != nil {
		n := &node{elem: e}
		r := s.root
		n.next = r
		s.root = n
	} else {
		s.root = &node{elem: e}
	}
}

type StackWithCAS struct {
	top unsafe.Pointer // *node
}

func NewStackWithCAS() *StackWithCAS {
	t := unsafe.Pointer(&node{})
	return &StackWithCAS{
		top: t,
	}
}

func (s *StackWithCAS) Pop() interface{} {
	var next, old *node
	for {
		old = (*node)(atomic.LoadPointer(&s.top))
		if old == nil {
			return nil
		}
		next = old.next
		if cas(&s.top, old, next) {
			break
		}
	}
	return old.elem
}

func (s *StackWithCAS) Push(elem interface{}) {
	n := &node{elem: elem, next: nil}
	var oldTop *node
	for {
		n.next = oldTop
		oldTop = (*node)(atomic.LoadPointer(&s.top))
		if cas(&s.top, oldTop, n) {
			break
		}
	}
}

func cas(addr *unsafe.Pointer, old, new *node) bool {
	return atomic.CompareAndSwapPointer(addr, unsafe.Pointer(old), unsafe.Pointer(new))
}
