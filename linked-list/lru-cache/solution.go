package main

type LRUCache struct {
	capacity int
	nodes    map[string]*Node
	head     *Node
	tail     *Node
}

func NewLRUCache(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{Prev: head}
	head.Next = tail

	return &LRUCache{
		capacity: capacity,
		nodes:    map[string]*Node{},
		head:     head,
		tail:     tail,
	}
}

// Node is type of record in doubly linked list of records
// Represents cache record with reference to next and previous record
type Node struct {
	Key  string
	Val  int
	Prev *Node
	Next *Node
}

// addNode adds node at head of list
func (c *LRUCache) addNode(n *Node) {
	n.Next = c.head.Next
	n.Prev = c.head
	c.head.Next.Prev = n
	c.head.Next = n
}

// removeNode remove node from list
func (c *LRUCache) removeNode(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
}

// moveToHead moves list element at head position
func (c *LRUCache) moveToHead(n *Node) {
	c.removeNode(n)
	c.addNode(n)
}

// popTail removes last element(from tail) in list
func (c *LRUCache) popTail() *Node {
	res := c.tail.Prev
	c.removeNode(res)

	return res
}

func (c *LRUCache) Put(key string, value int) {
	n, ok := c.nodes[key]
	if ok {
		n.Val = value
		c.moveToHead(n)
		return
	}

	n = &Node{Key: key, Val: value}
	c.nodes[key] = n
	c.addNode(n)
	if len(c.nodes) > c.capacity {
		delLast := c.popTail()
		delete(c.nodes, delLast.Key)
	}
}

func (c *LRUCache) Get(key string) (int, bool) {
	if n, ok := c.nodes[key]; ok {
		c.moveToHead(n)
		return n.Val, true
	}

	return -1, false
}
