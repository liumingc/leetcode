package main

import "fmt"

type LRUCache struct {
	tbl        map[int]kty
	head, tail kty
	cap        int
	sto        []node // store the LRU list
	fi         kty    // free index
}

type kty int16

type node struct {
	pre, nex kty
	key      int
	val      int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		tbl:  make(map[int]kty, capacity),
		cap:  capacity,
		sto:  make([]node, capacity),
		head: -1,
		tail: -1,
	}
}

func (this *LRUCache) Get(key int) int {
	defer func() {
		fmt.Println("get ", key)
		this.Show()
	}()
	x, ok := this.tbl[key]
	if !ok {
		return -1
	}
	if x != this.tail {
		// remove from current position
		this.remove(x)

		// and append to tail
		this.append(x)
	}
	return this.sto[x].val
}

func (this *LRUCache) remove(x kty) {
	pre := this.sto[x].pre
	nex := this.sto[x].nex
	if pre != -1 {
		this.sto[pre].nex = nex
	}
	if nex != -1 {
		this.sto[nex].pre = pre
	}
	if this.head == x {
		this.head = nex
	}
}

func (this *LRUCache) append(x kty) {
	this.sto[x].pre = this.tail
	this.sto[x].nex = -1
	this.sto[this.tail].nex = x
	this.tail = x
}

// drop head element
func (this *LRUCache) drop() {
	if this.cap == 1 {
		this.head = -1
		this.tail = -1
		this.fi = 0
		return
	}

	head := this.head
	// update head
	this.head = this.sto[this.head].nex
	this.sto[this.head].pre = -1

	this.fi--
	if head == this.fi {
		return
	}
	// swap last & old head
	lastN := this.sto[this.fi]
	this.tbl[lastN.key] = head
	pre := lastN.pre
	nex := lastN.nex
	if pre != -1 {
		this.sto[pre].nex = head
	}
	if nex != -1 {
		this.sto[nex].pre = head
	}
	this.sto[head] = lastN
	if this.tail == this.fi {
		this.tail = head
	}
}

func (this *LRUCache) Put(key int, value int) {
	defer func() {
		fmt.Println("put ", key, ",", value)
		this.Show()
	}()
	x, ok := this.tbl[key]
	sz := len(this.tbl)
	if ok {
		// update the node
		this.sto[x].val = value
		if sz == 1 || x == this.tail {
			return
		}
		// remove the node from original position
		this.remove(x)

		// and append to tail
		this.append(x)
		return
	}

	if sz == this.cap {
		// drop head
		delete(this.tbl, this.sto[this.head].key)
		this.drop()
	}

	fi := this.fi
	this.fi++
	this.tbl[key] = fi
	this.sto[fi].key = key
	this.sto[fi].val = value
	this.sto[fi].pre = this.tail
	this.sto[fi].nex = -1
	// if first element
	if this.tail == -1 {
		this.head = fi
		this.tail = fi
		return
	}
	// otherwise, append to tail
	tail := this.tail
	this.sto[tail].nex = fi
	this.tail = fi
}

func (this *LRUCache) Show() {
	fmt.Printf("cap=%d, tbl=%v\n", this.cap, this.tbl)
	for i := this.head; i != -1; {
		xN := this.sto[i]
		fmt.Printf("[%d] pre=%d, nex=%d, %d => %d\n",
			i, xN.pre, xN.nex, xN.key, xN.val)
		i = xN.nex
	}
	fmt.Println("=====")
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
