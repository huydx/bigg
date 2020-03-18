package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue(4)
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	if i, ok := q.dequeue(); !ok || i != 1 {
		t.Errorf("need 1, got %v %d", ok, i)
	}
	if i, ok := q.dequeue(); !ok || i != 2 {
		t.Errorf("need 2, got %v %d", ok, i)
	}
	if i, ok := q.dequeue(); !ok || i != 3 {
		t.Errorf("need 3, got %v %d", ok, i)
	}
	if i, ok := q.dequeue(); !ok || i != 4 {
		t.Errorf("need 3, got %v %d", ok, i)
	}
	if _, ok := q.dequeue(); ok {
		t.Errorf("need false, got %v", ok)
	}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	if i, ok := q.dequeue(); !ok || i != 1 {
		t.Errorf("need 1, got %v %d", ok, i)
	}
	if i, ok := q.dequeue(); !ok || i != 2 {
		t.Errorf("need 2, got %v %d", ok, i)
	}
	if i, ok := q.dequeue(); !ok || i != 3 {
		t.Errorf("need 3, got %v %d", ok, i)
	}
	if i, ok := q.dequeue(); !ok || i != 4 {
		t.Errorf("need 3, got %v %d", ok, i)
	}


	q2 := NewQueue(1)
	q2.enqueue(1)
	if ok := q2.enqueue(2); ok {
		t.Errorf("need false, got %v", ok)
	}
	if i, ok := q2.dequeue(); !ok || i != 1 {
		t.Errorf("need 1, got %v %d", ok, i)
	}
	if ok := q2.enqueue(2); !ok {
		t.Errorf("need true, got %v", ok)
	}
	if i, ok := q2.dequeue(); !ok || i != 2 {
		t.Errorf("need 1, got %v %d", ok, i)
	}
}
