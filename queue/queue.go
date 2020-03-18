package queue

// simple queue implement using ring buffer
type Queue struct {
	arr  []int
	head int
	tail int
	size int
	max  int
}

func NewQueue(size int) *Queue {
	return &Queue{
		arr:  make([]int, size),
		head: 0,
		tail: 0,
		size: 0,
		max:  size,
	}
}

func (q *Queue) enqueue(i int) (ok bool) {
	if q.size == q.max {
		return false
	}
	q.arr[q.tail] = i
	q.tail = (q.tail + 1) % q.max
	q.size += 1
	return true
}

func (q *Queue) dequeue() (i int, ok bool) {
	if q.size == 0 {
		return 0, false
	}
	r := q.arr[q.head]
	q.head = (q.head + 1) % q.max
	q.size -= 1
	return r, true
}
