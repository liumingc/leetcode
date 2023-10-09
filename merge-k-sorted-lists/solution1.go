package main

import "container/heap"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	nL := len(lists)
	if nL == 0 {
		return nil
	}
	if nL == 1 {
		return lists[0]
	}

	lh := &LH{}
	heap.Init(lh)
	for i := 0; i < nL; i++ {
		if lists[i] == nil {
			continue
		}
		heap.Push(lh, lists[i])
	}

	dummy := &ListNode{}
	pre := dummy

	for lh.Len() > 0 {
		sel := heap.Pop(lh).(*ListNode)
		pre.Next = sel
		pre = sel
		nxt := sel.Next

		if lh.Len() == 0 {
			break
		}

		if nxt != nil {
			heap.Push(lh, nxt)
		}
		// sel.Next = nil
	}
	return dummy.Next
}

type LH []*ListNode // list heap

func (lh LH) Len() int           { return len(lh) }
func (lh LH) Less(i, j int) bool { return lh[i].Val < lh[j].Val }
func (lh LH) Swap(i, j int)      { lh[i], lh[j] = lh[j], lh[i] }

func (lh *LH) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*lh = append(*lh, x.(*ListNode))
}

func (lh *LH) Pop() any {
	old := *lh
	n := len(old)
	x := old[n-1]
	*lh = old[0 : n-1]
	return x
}
