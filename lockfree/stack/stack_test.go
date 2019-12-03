package stack

import "testing"

func TestStack_Pop(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	if s.Pop() != 5 {
		t.Error("need 1")
	}
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	if s.Pop() != nil {
		t.Error("need nil")
	}
}
