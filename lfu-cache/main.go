package main

func main() {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Get(1)
	lfu.Put(3, 3)
	lfu.Get(2)
	lfu.Get(3)
	lfu.Put(4, 4)
	lfu.Get(1)
	lfu.Get(3)
	lfu.Get(4)
}
