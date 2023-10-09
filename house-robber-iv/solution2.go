package main

import (
	"container/heap"
	"fmt"
)

// bad result, not passed
func minCapability2(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if k == 1 {
		res := nums[0]
		for _, x := range nums[1:] {
			if x > res {
				res = x
			}
		}
		return res
	}

	sel0 := prioq{}
	sel1 := prioq{}
	heap.Init(&sel0)
	heap.Init(&sel1)
	bound := 2 * k
	for i := 0; i < bound; i += 2 {
		heap.Push(&sel0, &house{i, nums[i]})
	}

	if n < bound {
		return sel0[0].money
	}

	for i := 1; i < bound+1; i += 2 {
		heap.Push(&sel1, &house{i, nums[i]})
	}

	repl := func(sel0, sel1 *prioq, i int) prioq {
		maxHouse := (*sel0)[0]
		if maxHouse.money > nums[i] {
			maxHouse.index = i
			maxHouse.money = nums[i]
			heap.Fix(sel0, 0)
		}
		maxHouse = (*sel1)[0]
		if maxHouse.money > nums[i] {
			if maxHouse.index == i-1 {
				maxHouse = (*sel1)[1] // sub max
				if maxHouse.money > nums[i] {
					maxHouse.index = i
					maxHouse.money = nums[i]
					heap.Fix(sel1, 1)
				}
			} else {
				maxHouse.index = i
				maxHouse.money = nums[i]
				heap.Fix(sel1, 0)
			}
		}
		max0 := (*sel0)[0].money
		max1 := (*sel1)[0].money
		if max0 < max1 {
			return *sel0
		}
		// what about max0 == max1
		return *dupPrioq(sel1)
	}

	fmt.Println(nums, ", ", k)
	printPrioq("sel0", &sel0)
	printPrioq("sel1", &sel1)

	for i := bound; i < n; i++ {
		if i&1 == 0 {
			sel0 = repl(&sel0, &sel1, i)
			printPrioq(fmt.Sprintf("sel0 %d", i), &sel0)
		} else {
			sel1 = repl(&sel1, &sel0, i)
			printPrioq(fmt.Sprintf("sel1 %d", i), &sel1)
		}
	}
	max0 := sel0[0].money
	max1 := sel1[0].money
	return min(max0, max1)
}

type house struct {
	index int
	money int
}

type prioq []*house

func (pq prioq) Len() int { return len(pq) }

func (pq prioq) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].money > pq[j].money
}

func (pq prioq) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *prioq) Push(x any) {
	h := x.(*house)
	*pq = append(*pq, h)
}

func (pq *prioq) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func dupPrioq(pq *prioq) *prioq {
	res := &prioq{}
	heap.Init(res)
	q := *pq
	for i := 0; i < pq.Len(); i++ {
		heap.Push(res, q[i])
	}
	return res
}

func printPrioq(prefix string, pq *prioq) {
	fmt.Println(prefix, " ==>")
	q := (*pq)
	for i := 0; i < pq.Len(); i++ {
		fmt.Printf("  [%d] -> %d\n", q[i].index, q[i].money)
	}
}
