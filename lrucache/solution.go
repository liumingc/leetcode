package main

/*
type LRUCache struct {
	tbl        map[int]*node
	head, tail *node
	cap        int
}

type node struct {
	pre, nex *node
	val      int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		tbl: make(map[int]*node, capacity),
		cap: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	x, ok := this.tbl[key]
	if ok {
		return x.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	x, ok := this.tbl[key]
	sz := len(this.tbl)
	if ok {
		// update the node
		x.val = value
		if sz == 1 || x == this.tail {
			return
		}
		// remove the node from original position
		pre := x.pre
		nex := x.nex
		pre.nex = nex
		nex.pre = pre

		// and append to tail
		x.pre = this.tail
		x.nex = nil
		this.tail = x
		return
	}

	x = &node{
		nil, nil, value,
	}
	this.tbl[key] = x
	// if first element
	if this.tail == nil {
		this.head = x
		this.tail = x
		return
	}
	// otherwise, append to tail
	tail := this.tail
	x.pre = tail
	tail.nex = x

	if sz == this.cap {
		// drop head
		this.head = this.head.nex
		this.head.pre = nil
	}
}
*/

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
