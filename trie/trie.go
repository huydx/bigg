

package trie

//import "fmt"

type TrieNode struct {
	isWord   bool
	value    string
	origin   string
	children map[string]*TrieNode
}

var count = 0

func NewTrie() *TrieNode {
	return &TrieNode{
		isWord:   false,
		value:    "",
		children: map[string]*TrieNode{},
	}
}

func (r *TrieNode) Origin() string {
	return r.origin
}

func (r *TrieNode) Insert(word string, origin string) {
	var currentNode *TrieNode
	currentNode = r
	for _, c := range word {
		var ch = string(c)
		if cn := currentNode.children[ch]; cn != nil {
			currentNode = cn
		} else {
			n := new(TrieNode)
			n.value = currentNode.value + ch
			count++
			n.children = make(map[string]*TrieNode)
			currentNode.children[ch] = n
			currentNode = n
		}
	}
	currentNode.origin = origin
	currentNode.isWord = true // end node
}

func (r *TrieNode) FindPrefix(prefix string, max int) []*TrieNode {
	currentNode := r
	ret := make([]*TrieNode, 0)
	for _, _w := range prefix {
		w := string(_w)
		if c := currentNode.children[w]; c != nil {
			currentNode = c
		} else {
			return ret
		}
	}

	//current node is found node, dfs now
	queue := NewQueue(1000)
	queue.Push(currentNode)
	var found int
	for {
		if queue.count == 0 {
			break
		}
		cn := queue.Pop()
		if cn.isWord {
			found++
			ret = append(ret, cn)
		}
		if found >= max {
			return ret
		}
		for _, v := range cn.children {
			queue.Push(v)
		}
	}
	return ret
}

// FIFO queue from https://gist.github.com/moraes/2141121
type FIFOQueue struct {
	nodes []*TrieNode
	size  int
	head  int
	tail  int
	count int
}

// Push adds a node to the queue.
func (q *FIFOQueue) Push(n *TrieNode) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*TrieNode, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

func NewQueue(size int) *FIFOQueue {
	return &FIFOQueue{
		nodes: make([]*TrieNode, size),
		size:  size,
	}
}

// Pop removes and returns a node from the queue in first to last order.
func (q *FIFOQueue) Pop() *TrieNode {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func (r *TrieNode) Format() string {
	ret := ""
	var queue = NewQueue(100)
	queue.Push(r)
	for {
		if queue.count == 0 {
			break
		}
		cn := queue.Pop()
		if cn.isWord {
			ret += cn.value + "\n"
		}
		for _, v := range cn.children {
			queue.Push(v)
		}
	}
	return ret
}

func (r *TrieNode) Count() int {
	return count
}
