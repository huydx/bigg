package stack

import (
	"sync"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	s := NewStackWithCAS()
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

func mknumslice(n int) []int {
	var s = make([]int, n)
	for i := range s {
		s[i] = i
	}
	return s
}

func BenchmarkStackWithCAS(b *testing.B) {
	producers := 100
	stack := NewStackWithCAS()
	b.Run("SPSC", func(b *testing.B) {
		var wg sync.WaitGroup
		wg.Add(2 * producers)
		b.ResetTimer()
		for p := 0; p < producers; p++ {
			go func(p int) {
				var total = b.N/producers + 1
				var numbers = mknumslice(total)
				for i := range numbers {
					stack.Push(i)
				}
				wg.Done()
			}(p)
		}
		go func(n int) {
			for p := 0; p < producers; p++ {
				go func(p int) {
					var total = b.N/producers + 1
					var numbers = mknumslice(total)
					for _ = range numbers {
						stack.Pop()
					}
					wg.Done()
				}(p)
			}
		}(b.N)
		wg.Wait()
	})
}

func BenchmarkStackWithMut(b *testing.B) {
	producers := 100
	stack := NewStackWithMut()
	b.Run("SPSC", func(b *testing.B) {
		var wg sync.WaitGroup
		wg.Add(2 * producers)
		b.ResetTimer()
		for p := 0; p < producers; p++ {
			go func(p int) {
				var total = b.N/producers + 1
				var numbers = mknumslice(total)
				for i := range numbers {
					stack.Push(i)
				}
				wg.Done()
			}(p)
		}
		go func(n int) {
			for p := 0; p < producers; p++ {
				go func(p int) {
					var total = b.N/producers + 1
					var numbers = mknumslice(total)
					for _ = range numbers {
						stack.Pop()
					}
					wg.Done()
				}(p)
			}
		}(b.N)
		wg.Wait()
	})
}
