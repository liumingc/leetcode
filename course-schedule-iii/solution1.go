package main

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	ih := &IntHeap{}
	heap.Init(ih)

	res := 0
	accDura := 0
	for i := 0; i < len(courses); i++ {
		dura, lastDay := courses[i][0], courses[i][1]
		nxtDura := accDura + dura
		if nxtDura <= lastDay {
			accDura = nxtDura
			res++
			heap.Push(ih, dura)
			continue
		}
		if ih.Len() > 0 {
			maxDura := (*ih)[0]
			if maxDura > dura {
				accDura = accDura + dura - maxDura
				// heap.Pop(ih)
				// heap.Push(ih, dura)
				(*ih)[0] = dura
				heap.Fix(ih, 0)
			}
		}
	}
	return res
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
