package main

import (
	"container/heap"
	"fmt"
)

type LFUCache struct {
	cache map[int]*node
	cap   int
	q     *pq
	tick  int
}

type node struct {
	key     int
	val     int
	freq    int
	lastAcc int
	i       int
}

func Constructor(capacity int) LFUCache {
	res := LFUCache{
		cache: make(map[int]*node),
		cap:   capacity,
		q:     &pq{},
	}
	heap.Init(res.q)
	return res
}

func (this *LFUCache) Get(key int) int {
	x, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.tick++
	x.lastAcc = this.tick
	x.freq++
	heap.Fix(this.q, x.i)
	fmt.Printf("get [%d => %d]\n", key, x.val)
	return x.val
}

func (this *LFUCache) Put(key int, value int) {
	x, ok := this.cache[key]
	this.tick++
	if !ok {
		if len(this.cache) >= this.cap {
			item := heap.Pop(this.q).(*node)
			fmt.Printf("elim key=%d\n", item.key)
			delete(this.cache, item.key)
		}

		x = &node{
			key:     key,
			val:     value,
			freq:    1,
			lastAcc: this.tick,
		}
		heap.Push(this.q, x)
		fmt.Printf("put [%d => %d], new\n", key, value)
	} else {
		x.freq++
		x.lastAcc = this.tick
		x.val = value
		heap.Fix(this.q, x.i)
		fmt.Printf("put [%d => %d]\n", key, value)
	}
	this.cache[key] = x
}

// An IntHeap is a min-heap of ints.
type pq []*node

func (h pq) Len() int { return len(h) }
func (h pq) Less(i, j int) bool {
	return h[i].freq < h[j].freq ||
		(h[i].freq == h[j].freq && h[i].lastAcc < h[j].lastAcc)
}

func (h pq) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].i = i
	h[j].i = j
}

func (h *pq) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	item := x.(*node)
	n := h.Len()
	item.i = n
	*h = append(*h, item)
}

func (h *pq) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	x.i = -1 //
	*h = old[0 : n-1]
	return x
}
