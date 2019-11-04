package main

import "fmt"

const M = 4

type Entry struct {
	key  int
	val  interface{}
	next *Node
}

type Node struct {
	m        int // number of children
	children []*Entry
}

func newNode(i int) *Node {
	return &Node{
		m:        i,
		children: make([]*Entry, 0),
	}
}

type Btree struct {
	root   *Node
	height int
	n      int // num of key-value pairs???
}

func NewBtree() *Btree {
	return &Btree{
		root: newNode(0),
	}
}

func (b *Btree) Get(k int) interface{} {
	return b.search(b.root, k, b.height)
}

func (b *Btree) Put(k int, v interface{}) {
	u := b.insert(b.root, k, v, b.height)
	b.n = b.n + 1
	if u == nil {
		return
	}

	t := newNode(2)
	t.children[0] = &Entry{key: b.root.children[0].key, val: nil, next: b.root}
	t.children[1] = &Entry{key: u.children[0].key, val: nil, next: u}
	b.root = t
	b.height = b.height + 1
}

// ht: height
func (b *Btree) search(node *Node, key int, ht int) interface{} {
	children := node.children

	if ht == 0 {
		// external node
		for j := 0; j < node.m; j++ {
			if key == children[j].key {
				return children[j].val
			}
		}
	} else {
		// internal node
		for j := 0; j < node.m; j++ {
			if j+1 == node.m || key < children[j+1].key {
				return b.search(children[j].next, key, ht-1)
			}
		}
	}

	return nil
}

func (b *Btree) split(node *Node) *Node {
	n := newNode(M / 2)
	node.m = M / 2
	for j := 0; j < M/2; j++ {
		n.children[j] = node.children[M/2+j]
	}
	return n
}

func (b *Btree) insert(node *Node, key int, val interface{}, ht int) *Node {
	entry := &Entry{
		key:  key,
		val:  val,
		next: nil,
	}
	var j int

	if ht == 0 {
		// internal node
		for j = 0; j < node.m; j++ {
			if key < node.children[j].key {
				break
			}
		}
	} else {
		// external node
		for j = 0; j < node.m; j++ {
			if j+1 == node.m || key < node.children[j+1].key {
				j = j + 1
				u := b.insert(node.children[j].next, key, val, ht-1)
				if u == nil {
					return nil
				}
				entry.key = u.children[0].key
				entry.next = u
				break
			}
		}
	}

	for i := node.m; i > j; i-- {
		node.children[i] = node.children[i-1]
	}

	node.children[j] = entry
	node.m = node.m + 1

	if node.m < M {
		return nil
	} else {
		return b.split(node)
	}
}

func main() {
	b := NewBtree()
	b.Put(1, "2")
	b.Put(2, "3")
	b.Put(2, "4")
	b.Put(3, "5")
	fmt.Println(b.Get(1))
}
