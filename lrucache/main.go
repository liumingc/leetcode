package main

func main() {
	// ["LRUCache","put","put","get","put","get","put","get","get","get"]
	// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	/*
		["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
		[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
	*/
	lru := Constructor(3)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	lru.Get(4)
	lru.Get(3)
	lru.Get(2)
	lru.Get(1)
	lru.Put(5, 5)
	lru.Get(1)
	lru.Get(2)
	lru.Get(3)
	lru.Get(4)
	lru.Get(5)
}
