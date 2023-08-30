package main

import (
	"container/heap"
	"fmt"
)

func trap0(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}

	// 需要两个结构保存信息
	// 按高度递减的列表(priority heap)，每个列表的元素有高度，以及出现高度的位置坐标
	tbl := make(map[int]*sameHeight, n)
	pq := &htHeap{}
	heap.Init(pq)
	for i := 0; i < n; i++ {
		ht := height[i]
		sh := tbl[ht]
		if sh == nil {
			heap.Push(pq, ht)
			sh = &sameHeight{
				arr: make([]int, 0, n),
			}
			tbl[ht] = sh
		}
		sh.arr = append(sh.arr, i)
	}

	vol := 0
	bit := make([]bool, n)
	for pq.Len() > 1 {
		ht := pq.Pop().(int)
		sh := tbl[ht]
		fmt.Println("pop ", ht, sh.arr)
		for len(sh.arr) >= 2 {
			left, right := sh.arr[0], sh.arr[1]
			fmt.Printf("left=%d, right=%d\n", left, right)
			sh.arr = sh.arr[2:len(sh.arr)]
			for i := left + 1; i < right; i++ {
				if bit[i] {
					continue
				}
				x := height[i] - ht
				if x > 0 {
					vol += x
				}
				bit[i] = true
			}
		}

	}

	return vol
}

type sameHeight struct {
	arr []int // index with same height
}

type htHeap []int

func (h htHeap) Len() int {
	return len(h)
}

func (h htHeap) Less(i, j int) bool {
	return h[i] < h[j] // with higher height, pops first
}

func (h htHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *htHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *htHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
