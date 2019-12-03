// Copyright (c) 2016 LINE Corporation. All rights reserved.
// LINE Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.

package stack

import (
	"unsafe"
	"sync/atomic"
)

type node struct {
	elem interface{}
	next *node
}

type Stack struct {
	top unsafe.Pointer // *node
}

func NewStack() *Stack {
	t := unsafe.Pointer(&node{})
	return &Stack{
		top: t,
	}
}

func (s *Stack) Pop() interface{} {
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

func (s *Stack) Push(elem interface{}) {
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
