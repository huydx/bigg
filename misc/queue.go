package main

import "fmt"

type queue struct {
	a []interface{}
}

func (q *queue) enqueue(i interface{}) {
	q.a = append(q.a, i)
}

func (q *queue) dequeue() interface{} {
	if len(q.a) == 0 {
		return nil
	}
	i := q.a[0]
	q.a = q.a[1:len(q.a)]
	return i
}

func main() {
	q := queue{
		a: make([]interface{}, 0),
	}
	fmt.Println(q.dequeue())
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
}
