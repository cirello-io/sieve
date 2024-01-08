package sieve

import "fmt"

type Node[T comparable] struct {
	Value T

	visited    bool
	prev, next *Node[T]
}

type Cache[T comparable] struct {
	capacity int

	cache            map[T]*Node[T]
	head, tail, hand *Node[T]
	size             int
}

func New[T comparable](capacity int) *Cache[T] {
	return &Cache[T]{
		capacity: capacity,
		cache:    make(map[T]*Node[T]),
	}
}

func (c *Cache[T]) addToHead(node *Node[T]) {
	node.next = c.head
	node.prev = nil
	if c.head != nil {
		c.head.prev = node
	}
	c.head = node
	if c.tail == nil {
		c.tail = node
	}
}

func (c *Cache[T]) removeNode(node *Node[T]) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		c.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		c.tail = node.prev
	}
}

func (c *Cache[T]) evict() {
	obj := c.tail
	if c.hand != nil {
		obj = c.hand
	}
	for obj != nil && obj.visited {
		obj.visited = false
		obj = c.tail
		if obj.prev != nil {
			obj = obj.prev
		}
	}
	c.hand = nil
	if obj.prev != nil {
		c.hand = obj.prev
	}
	delete(c.cache, obj.Value)
	c.removeNode(obj)
	c.size--
}

func (c *Cache[T]) Access(x T) {
	if _, ok := c.cache[x]; ok {
		node := c.cache[x]
		node.visited = true
	} else {
		if c.size == c.capacity {
			c.evict()
		}
		newNode := &Node[T]{
			Value: x,
		}
		c.addToHead(newNode)
		c.cache[x] = newNode
		c.size++
		newNode.visited = false
	}
}

func (self *Cache[T]) Show() {
	current := self.head
	for current != nil {
		fmt.Printf("%v (Visited: %v)", current.Value, current.visited)
		if current.next != nil {
			fmt.Printf(" -> ")
		} else {
			fmt.Printf("\n")
		}
		current = current.next
	}
}
