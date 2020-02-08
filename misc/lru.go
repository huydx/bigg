package main

type element struct {
	entry *entry
	next  *element
	prev  *element
}

type entry struct {
	k key
	v interface{}
}

type list struct {
	root *element
	last *element
	// not thread safe
	len int
}

func (l *list) add(e *element) {
	l.last.next = e
	e.prev = l.last
	l.last = e
}

func (l *list) pushFront(ee *element) *element {
	ee.next = l.root
	l.root = ee
	return ee
}

func (l *list) moveToFront(ee *element) {
	prev := ee.prev
	next := ee.next
	prev.next = next
	ee.next = l.root
	ee.prev = nil
	l.root = ee
}

func (l *list) length() int {
	i := 0
	for e := l.root; e != l.last; e = e.next {
		i++
	}
	return i+1
}

func (l *lru) evict() {
	l.ll.last
}

type key interface{}

type lru struct {
	ll         *list
	cache      map[key]*element
	maxEntries int // could be int64?
}

func (l *lru) add(k interface{}, v interface{}) {
	if l.cache == nil {
		l.cache = make(map[key]*element)
		l.ll = &list{}
	}
	if ee, ok := l.cache[k]; ok {
		l.ll.moveToFront(ee)
		ee.entry.v = v
		return
	}
	ele := l.ll.pushFront(&element{entry: &entry{k, v}})
	l.cache[k] = ele
	if l.ll.length() >= l.maxEntries {
		l.evict()
	}
}
